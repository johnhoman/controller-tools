package manager

import (
	"context"
	"net/http"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/johnhoman/controller-tools/prefab"
)

type Manager struct {
	manager   ctrl.Manager
	ctx       context.Context
	namespace string
}

func (m *Manager) SetFields(i interface{}) error { return m.manager.SetFields(i) }
func (m *Manager) GetConfig() *rest.Config       { return m.manager.GetConfig() }
func (m *Manager) GetScheme() *runtime.Scheme    { return m.manager.GetScheme() }
func (m *Manager) GetClient() client.Client      { return m.manager.GetClient() }
func (m *Manager) GetFieldIndexer() client.FieldIndexer {
	return m.manager.GetFieldIndexer()
}

func (m *Manager) GetCache() cache.Cache { return m.manager.GetCache() }
func (m *Manager) GetEventRecorderFor(name string) record.EventRecorder {
	return m.manager.GetEventRecorderFor(name)
}
func (m *Manager) GetRESTMapper() meta.RESTMapper { return m.manager.GetRESTMapper() }
func (m *Manager) GetAPIReader() client.Reader    { return m.manager.GetAPIReader() }
func (m *Manager) Start(ctx context.Context) error {
	if m.ctx == nil {
		m.ctx = ctx
	}
	return m.manager.Start(m.ctx)
}

func (m *Manager) Add(runnable manager.Runnable) error { return m.manager.Add(runnable) }

func (m *Manager) Elected() <-chan struct{} { return m.manager.Elected() }
func (m *Manager) AddMetricsExtraHandler(path string, handler http.Handler) error {
	return m.manager.AddMetricsExtraHandler(path, handler)
}

func (m *Manager) AddHealthzCheck(name string, check healthz.Checker) error {
	return m.manager.AddHealthzCheck(name, check)
}

func (m *Manager) AddReadyzCheck(name string, check healthz.Checker) error {
	return m.manager.AddReadyzCheck(name, check)
}

func (m *Manager) GetWebhookServer() *webhook.Server { return m.manager.GetWebhookServer() }
func (m *Manager) GetLogger() logr.Logger            { return m.manager.GetLogger() }
func (m *Manager) GetControllerOptions() v1alpha1.ControllerConfigurationSpec {
	return m.manager.GetControllerOptions()
}

func (m *Manager) GetContext() context.Context {
	ctx := m.ctx
	if m.ctx == nil {
		ctx = context.Background()
	}
	return ctx
}

func (m *Manager) GetNamespace() string {
	return m.namespace
}

func (m *Manager) NamespacedClient() client.Client {
	if m.namespace == "" {
		// no namespace is set
		ns, c, err := m.RandomNamespace()
		if err != nil {
			panic(err)
		}
		m.namespace = ns.GetName()
		return c
	}
	cli, err := client.New(
		m.GetConfig(),
		client.Options{Scheme: m.GetScheme()},
	)
	if err != nil {
		panic(err)
	}
	return client.NewNamespacedClient(cli, m.namespace)
}

func (m *Manager) RandomNamespace() (*corev1.Namespace, client.Client, error) {
	ns := &corev1.Namespace{}
	prefab.RandomName()(ns)
	ctx := m.ctx
	if ctx == nil {
		ctx = context.TODO()
	}
	if err := m.GetClient().Create(ctx, ns); err != nil {
		return nil, nil, err
	}
	cli, err := client.New(
		m.GetConfig(),
		client.Options{Scheme: m.GetScheme()},
	)
	if err != nil {
		return nil, nil, err
	}
	return ns, client.NewNamespacedClient(cli, ns.GetName()), nil
}

var _ ctrl.Manager = &Manager{}

func NewManager(cfg *rest.Config, scheme *runtime.Scheme, opts ...func(options ctrl.Options)) *Manager {
	options := ctrl.Options{
		HealthProbeBindAddress: "0",
		MetricsBindAddress:     "0",
		Scheme:                 scheme,
		LeaderElection:         false,
	}
	for _, opt := range opts {
		opt(options)
	}
	mgr, err := ctrl.NewManager(cfg, options)
	if err != nil {
		panic(err)
	}
	return &Manager{manager: mgr}
}
