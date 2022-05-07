package serviceaccount

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/johnhoman/controller-tools/defaulting"
)

type imagePullSecret string

func (im imagePullSecret) Apply(obj client.Object) {
	serviceAccount, ok := obj.(*corev1.ServiceAccount)
	if !ok {
		panic(fmt.Sprintf("cannot use type %T as *ServiceAccount", obj))
	}
	if serviceAccount.ImagePullSecrets == nil {
		serviceAccount.ImagePullSecrets = make([]corev1.LocalObjectReference, 0, 1)
	}
	if !hasLocalObjectRef(serviceAccount.ImagePullSecrets, string(im)) {
		serviceAccount.ImagePullSecrets = append(
			serviceAccount.ImagePullSecrets,
			corev1.LocalObjectReference{Name: string(im)},
		)
	}
}

func ImagePullSecret(name string) defaulting.Default {
	return imagePullSecret(name)
}

func hasLocalObjectRef(refs []corev1.LocalObjectReference, name string) bool {
	for _, item := range refs {
		if item.Name == name {
			return true
		}
	}
	return false
}
