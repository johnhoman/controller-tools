package suite

import (
	"context"
	"github.com/johnhoman/controller-tools/testing/manager"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Kubernetes struct {
	Suite
	cfg        *rest.Config
	context    context.Context
	cancelFunc context.CancelFunc
	namespace  string
	mgr        *manager.Manager
}

func (k *Kubernetes) SetupTest()                     {}
func (k *Kubernetes) TearDownTest()                  {}
func (k *Kubernetes) NamespaceClient() client.Client { panic("not implemented!") }
