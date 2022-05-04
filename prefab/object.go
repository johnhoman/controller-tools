package prefab

import (
	"github.com/google/uuid"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

func uid() string {
	return uuid.New().String()[:8]
}

func SetName(name string) func(client.Object) {
	return func(obj client.Object) {
		obj.SetName(name)
	}
}

func RandomName() func(client.Object) {
	return func(obj client.Object) {
		kind := obj.GetObjectKind().GroupVersionKind().Kind
		if kind == "" {
			kind = "default"
		}
		kind = strings.ToLower(kind)
		obj.SetName(kind + "-" + uid())
	}
}

func InNamespace(namespace string) func(client.Object) {
	return func(obj client.Object) {
		obj.SetNamespace(namespace)
	}
}
