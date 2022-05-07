package defaulting

import (
	"github.com/google/uuid"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

type Name string

func (n Name) Apply(obj client.Object) {
	if obj.GetName() != "" {
		return
	}
	kind := obj.GetObjectKind().GroupVersionKind().Kind
	if kind == "" {
		kind = "default"
	}
	kind = strings.ToLower(kind)
	obj.SetName(kind + "-" + string(n))
}

func RandomName() Name {
	return Name(uuid.New().String()[:8])
}

type Namespace string

func (n Namespace) Apply(obj client.Object) {
	if obj.GetNamespace() == "" {
		obj.SetNamespace(string(n))
	}
}

func InNamespace(namespace string) Namespace {
	return Namespace(namespace)
}

type annotation struct {
	key   string
	value string
}

func (a *annotation) Apply(obj client.Object) {
	initAnnotations(obj)
	m := setDefault(obj.GetAnnotations(), a.key, a.value)
	obj.SetAnnotations(m)
}

func WithAnnotation(key, value string) *annotation {
	return &annotation{key, value}
}

type label struct {
	key   string
	value string
}

func (a *label) Apply(obj client.Object) {
	initLabels(obj)
	m := setDefault(obj.GetLabels(), a.key, a.value)
	obj.SetLabels(m)
}

func WithLabels(key, value string) *label {
	return &label{key, value}
}

func setDefault(m map[string]string, key, value string) map[string]string {
	if _, ok := m[key]; !ok {
		m[key] = value
	}
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
