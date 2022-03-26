package predicate

import (
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type Predicate func(obj client.Object) bool

func hasAnnotation(obj client.Object, annotation string) bool {
    annotations := obj.GetAnnotations()
    if annotations != nil {
        _, ok := annotations[annotation]
        return ok
    }
    return false
}

func hasAnnotationWithValue(
    obj client.Object,
    annotation string,
    value string,
) bool {
    if hasAnnotation(obj, annotation) {
        if obj.GetAnnotations()[annotation] == value {
            return true
        }
    }
    return false
}

// HasAnnotation creates a predicate function that will return true
// if the annotation exists in resource metadata
func HasAnnotation(annotation string) Predicate {
    return func(obj client.Object) bool {
        return hasAnnotation(obj, annotation)
    }
}

// HasAnnotationWithValue creates a predicate function that will return true
// if the annotation exists in resource metadata and the value of the annotation
// matches the provided value
func HasAnnotationWithValue(annotation string, value string) Predicate {
    return func(obj client.Object) bool {
        return hasAnnotationWithValue(obj, annotation, value)
    }
}

// HasFinalizer creates a predicate function that will return true if the
// object metadata contains the provided finalizer
func HasFinalizer(finalizer string) Predicate {
    return func(obj client.Object) bool {
        return controllerutil.ContainsFinalizer(obj, finalizer)
    }
}