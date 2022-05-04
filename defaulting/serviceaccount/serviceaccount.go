package serviceaccount

import (
    "fmt"
    "github.com/johnhoman/controller-tools/defaulting"
    corev1 "k8s.io/api/core/v1"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

func ImagePullSecret(name string) defaulting.Func {
    return func(obj client.Object) {
        serviceAccount, ok := obj.(*corev1.ServiceAccount)
        if !ok {
            panic(fmt.Sprintf("cannot use type %T as *ServiceAccount", obj))
        }
        if serviceAccount.ImagePullSecrets == nil {
            serviceAccount.ImagePullSecrets = make([]corev1.LocalObjectReference, 0, 1)
        }
        serviceAccount.ImagePullSecrets = append(
            serviceAccount.ImagePullSecrets,
            corev1.LocalObjectReference{Name: name},
        )
    }
}
