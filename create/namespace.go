package create

import (
    corev1 "k8s.io/api/core/v1"

    "github.com/johnhoman/controller-tools"
    "github.com/johnhoman/controller-tools/crud"
    "github.com/johnhoman/controller-tools/defaulting"
)

func randomNamespace() *corev1.Namespace {
    ns := &corev1.Namespace{}
    ns.SetName("namespace-" + uid())
    return ns
}

func Namespace(mgr *controllertools.Manager, defaults ...defaulting.Func) (*corev1.Namespace, error) {
    ns := randomNamespace()
    for _, opt := range defaults {
        opt(ns)
    }
    if err := crud.Create(mgr, ns); err != nil {
        return nil, err
    }
    return ns, nil
}