package create

import (
    "github.com/johnhoman/controller-tools/crud"
    corev1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "sigs.k8s.io/controller-runtime/pkg/client"

    "github.com/johnhoman/controller-tools"
    "github.com/johnhoman/controller-tools/prefab"
)

// ServiceAccount creates a new service account with a random name. Default values
// can be applied with options functions e.g.
// ServiceAccount(mgr, func(obj client.Object) { obj.SetName("my-service-account") })
func ServiceAccount(
    mgr *controllertools.Manager,
    options ...func(object client.Object),
) (*corev1.ServiceAccount, error) {
    obj := &corev1.ServiceAccount{
        TypeMeta: metav1.TypeMeta{
            Kind: "ServiceAccount",
            APIVersion: "v1",
        },
    }
    for _, opt := range options {
        opt(obj)
    }
    opts := []func(client.Object){}
    if obj.GetName() == "" {
        opts = append(opts, prefab.RandomName())
        prefab.RandomName()(obj)
    }
    if obj.GetNamespace() == "" {
        opts = append(opts, prefab.InNamespace(mgr.GetNamespace()))
    }
    for _, opt := range opts {
        opt(obj)
    }

    err := crud.Create(mgr, obj)
    if err != nil {
        return nil, err
    }
    return obj, nil
}
