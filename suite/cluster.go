package suite

import (
    "context"

    "k8s.io/client-go/rest"
    "sigs.k8s.io/controller-runtime/pkg/client"

    "github.com/johnhoman/controller-tools"
)

type Kubernetes struct {
    Suite
    cfg *rest.Config
    context context.Context
    cancelFunc context.CancelFunc
    namespace string
    mgr *controllertools.Manager
}

func (k *Kubernetes) SetupTest() {}
func (k *Kubernetes) TearDownTest() {}
func (k *Kubernetes) NamespaceClient() client.Client { panic("not implemented!") }