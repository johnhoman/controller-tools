/*
Copyright 2022 John Homan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package manager

import (
	"context"

	"github.com/google/uuid"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/johnhoman/controller-tools/eventually"
)

type integrationTestManager struct {
	ctrl.Manager
	ctx        context.Context
	cancelFunc context.CancelFunc
	eventually eventually.Eventually
	client     client.Client
}

func (i *integrationTestManager) StartManager() {
	go func() {
		defer ginkgo.GinkgoRecover()
		err := i.Manager.Start(i.ctx)
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
	}()
}

func (i *integrationTestManager) StopManager() {
	i.cancelFunc()
}

func (i *integrationTestManager) GetContext() context.Context {
	return i.ctx
}

func (i *integrationTestManager) Eventually() AwaitClient {
	aw := awaitClient{ctx: i.ctx, eventually: i.eventually}
	return &aw
}

func (i *integrationTestManager) Expect() Client {
	k := k8client{ctx: i.ctx, k8: i.client}
	return &k
}

func (i *integrationTestManager) Uncached() client.Client {
	return i.client
}

var _ IntegrationTest = &integrationTestManager{}

type builder struct {
	scheme            *runtime.Scheme
	namespace         string
	webhookOpts       *envtest.WebhookInstallOptions
	isolatedNamespace bool
}

func (b *builder) WithWebhookInstallOptions(options envtest.WebhookInstallOptions) Builder {
	b.webhookOpts = &options
	return b
}

func (b *builder) Isolate() Builder {
	return b.WithIsolatedNamespace(true)
}

func (b *builder) WithIsolatedNamespace(x bool) Builder {
	b.isolatedNamespace = x
	return b
}

func (b *builder) WithScheme(scheme *runtime.Scheme) Builder {
	b.scheme = scheme
	return b
}

func (b *builder) Complete(cfg *rest.Config) IntegrationTest {
	opts := ctrl.Options{
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
		LeaderElection:         false,
		// Namespace: b.namespace,
	}
	if b.isolatedNamespace {
		opts.Namespace = b.namespace
	}
	opts.Scheme = b.scheme
	if opts.Scheme == nil {
		opts.Scheme = scheme.Scheme
	}
	if b.webhookOpts != nil {
		opts.Host = b.webhookOpts.LocalServingHost
		opts.Port = b.webhookOpts.LocalServingPort
		opts.CertDir = b.webhookOpts.LocalServingCertDir
	}
	k8, err := client.New(cfg, client.Options{Scheme: opts.Scheme})
	gomega.Expect(err).To(gomega.Succeed())

	mgr, err := ctrl.NewManager(cfg, opts)
	gomega.Expect(err).ToNot(gomega.HaveOccurred())

	ctx, cancel := context.WithCancel(context.Background())

	ns := &corev1.Namespace{}
	ns.SetName(b.namespace)
	gomega.Expect(k8.Create(ctx, ns)).To(gomega.Succeed())

	return &integrationTestManager{
		Manager:    mgr,
		eventually: eventually.New(client.NewNamespacedClient(k8, b.namespace)),
		client:     client.NewNamespacedClient(k8, b.namespace),
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

var _ Builder = &builder{}

func IntegrationTestBuilder() *builder {
	ns := "testspace-" + uuid.New().String()[:8]
	return &builder{
		namespace:         ns,
		webhookOpts:       nil,
		scheme:            nil,
		isolatedNamespace: true,
	}
}

type awaitClient struct {
	ctx        context.Context
	eventually eventually.Eventually
}

func (a *awaitClient) Get(key types.NamespacedName, obj client.Object) gomega.AsyncAssertion {
	return a.eventually.Get(a.ctx, key, obj)
}

func (a *awaitClient) Create(object client.Object) gomega.AsyncAssertion {
	return a.eventually.Create(a.ctx, object)
}

func (a *awaitClient) Update(object client.Object) gomega.AsyncAssertion {
	return a.eventually.Update(a.ctx, object)
}

func (a *awaitClient) GetWhen(key types.NamespacedName, obj client.Object, predicateFunc eventually.PredicateFunc) gomega.AsyncAssertion {
	return a.eventually.GetWhen(a.ctx, key, obj, predicateFunc)
}

var _ AwaitClient = &awaitClient{}

type k8client struct {
	ctx context.Context
	k8  client.Client
}

func (k *k8client) Get(key types.NamespacedName, obj client.Object) gomega.Assertion {
	return gomega.Expect(k.k8.Get(k.ctx, key, obj))
}

func (k *k8client) Create(object client.Object) gomega.Assertion {
	return gomega.Expect(k.k8.Create(k.ctx, object))
}

func (k *k8client) Update(object client.Object) gomega.Assertion {
	return gomega.Expect(k.k8.Update(k.ctx, object))
}

func (k *k8client) Delete(object client.Object) gomega.Assertion {
	return gomega.Expect(k.k8.Delete(k.ctx, object))
}

var _ Client = &k8client{}
