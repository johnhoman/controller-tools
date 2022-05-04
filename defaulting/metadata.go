package defaulting

import "sigs.k8s.io/controller-runtime/pkg/client"


func WithAnnotation(key, value string) Func {
    return func(obj client.Object) {
        initAnnotations(obj)
        m := setDefault(obj.GetAnnotations(), key, value)
        obj.SetAnnotations(m)
    }
}

func WithLabels(key, value string) Func {
    return func(obj client.Object) {
        initLabels(obj)
        m := setDefault(obj.GetLabels(), key, value)
        obj.SetLabels(m)
    }
}

func setDefault(m map[string]string, key, value string) map[string]string {
    m[key] = value
    return m
}

func initAnnotations(obj client.Object) {
    annotations := obj.GetAnnotations()
    if annotations == nil {
        annotations = make(map[string]string)
    }
    obj.SetAnnotations(annotations)
}

func initLabels(obj client.Object) {
    labels := obj.GetLabels()
    if labels == nil {
        labels = make(map[string]string)
    }
    obj.SetLabels(labels)
}