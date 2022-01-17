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
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/johnhoman/controller-tools/eventually"
)

type Client interface {
	Get(key types.NamespacedName, obj client.Object) gomega.Assertion
	Create(client.Object) gomega.Assertion
	Update(client.Object) gomega.Assertion
	Delete(client.Object) gomega.Assertion
}

type AwaitClient interface {
	Get(key types.NamespacedName, obj client.Object) gomega.AsyncAssertion
	Create(client.Object) gomega.AsyncAssertion
	Update(client.Object) gomega.AsyncAssertion
	GetWhen(key types.NamespacedName, obj client.Object, predicateFunc eventually.PredicateFunc) gomega.AsyncAssertion
}


type IntegrationTest interface {
	ctrl.Manager
	StartManager()
	StopManager()
	StopManagerFunc() interface{}
	GetContext() context.Context
	Eventually() AwaitClient
	Expect() Client
}

type Builder interface {
	WithWebhookInstallOptions(options envtest.WebhookInstallOptions) Builder
	WithScheme(*runtime.Scheme) Builder
	Complete(cfg *rest.Config) IntegrationTest
}

