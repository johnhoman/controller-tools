package create

import (
	"github.com/johnhoman/controller-tools/defaulting"
	"github.com/johnhoman/controller-tools/internal/testing/crud"
	"github.com/johnhoman/controller-tools/testing"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceAccount creates a new service account with a random name. Default values
// can be applied with options functions e.g.
func ServiceAccount(mgr testing.Manager, options ...defaulting.Default) (*corev1.ServiceAccount, error) {
	obj := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
	}
	options = append(options, defaulting.RandomName())
	options = append(options, defaulting.InNamespace(mgr.GetNamespace()))
	for _, opt := range options {
		opt.Apply(obj)
	}
	if err := crud.Create(mgr, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
