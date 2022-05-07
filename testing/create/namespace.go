package create

import (
    "github.com/johnhoman/controller-tools/testing"
    "github.com/johnhoman/controller-tools/testing/internal/testing/crud"
    corev1 "k8s.io/api/core/v1"

    "github.com/johnhoman/controller-tools/defaulting"
)

func randomNamespace() *corev1.Namespace {
	ns := &corev1.Namespace{}
	ns.SetName("namespace-" + uid())
	return ns
}

func Namespace(mgr testing.Manager, defaults ...defaulting.Default) (*corev1.Namespace, error) {
	ns := randomNamespace()
	for _, opt := range defaults {
		opt.Apply(ns)
	}
	if err := crud.Create(mgr, ns); err != nil {
		return nil, err
	}
	return ns, nil
}
