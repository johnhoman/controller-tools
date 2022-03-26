package predicate_test

import (
    corev1 "k8s.io/api/core/v1"
    "testing"

    "github.com/johnhoman/controller-tools/predicate"
)


func TestHasAnnotation(t *testing.T) {
    obj := &corev1.ServiceAccount{}
    obj.SetAnnotations(map[string]string{
        "eks.amazonaws.com/role-arn": "arn:aws:iam::01234567890:role/asdf",
    })
    result := predicate.HasAnnotation("eks.amazonaws.com/role-arn")(obj)
    if !result {
        t.Errorf("predicate failed. Have %v, want %v", result, true)
    }

    result = predicate.HasAnnotation("eks.amazonaws.com")(obj)
    if result {
        t.Errorf("predicate failed. Have %v, want %v", result, false)
    }
}

func TestHasAnnotationWithValue(t *testing.T) {
    obj := &corev1.ServiceAccount{}
    obj.SetAnnotations(map[string]string{
        "eks.amazonaws.com/role-arn": "arn:aws:iam::01234567890:role/asdf",
    })
    pred := func(value string) bool {
        return predicate.HasAnnotationWithValue("eks.amazonaws.com/role-arn", value)(obj)
    }
    result := pred("arn:aws:iam::01234567890:role/asdf")
    if !result {
        t.Errorf("predicate failed. Have %v, want %v", result, "arn:aws:iam::01234567890:role/asdf")
    }

    result = pred("role/asdf")
    if result {
        t.Errorf("predicate failed. Have %v, want %v", result, false)
    }
}

func TestHasFinalizer(t *testing.T) {
    obj := &corev1.ServiceAccount{}
    obj.SetFinalizers([]string{"external-resource"})

    res := predicate.HasFinalizer("external-resource")(obj)
    if !res {
        t.Errorf("want %v, have %v", true, res)
    }
}