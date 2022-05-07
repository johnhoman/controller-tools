package predicate

import (
    "fmt"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "strings"
    "testing"

    "github.com/stretchr/testify/require"
    corev1 "k8s.io/api/core/v1"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

func TestReadWhen(t *testing.T) {
    tests := []struct{
        client.Object
        jsonPath string
        condition func(v interface{}) bool
        expected bool
    } {
        {
            &corev1.Pod{Status: corev1.PodStatus{
                PodIP: "192.168.2.1",
            }},
            "status.podIP",
            func(v interface{}) bool { return strings.HasPrefix(v.(string), "192.168")},
            true,
        },
        {
            &corev1.Pod{Status: corev1.PodStatus{
                PodIP: "10.20.30.40",
            }},
            "status.podIP",
            func(v interface{}) bool { return strings.HasPrefix(v.(string), "192.168")},
            false,
        },
        {
            &corev1.Pod{
                ObjectMeta: metav1.ObjectMeta{
                    Finalizers: []string{"keep-alive", "upstream-delete"},
                },
            },
            "metadata.finalizers",
            func(v interface{}) bool {
                finalizers := v.([]interface{})
                for _, item := range finalizers {
                    if item.(string) == "keep-alive" { return true }
                }
                return false
            },
            true,
        },
    }
    for k, subtest := range tests {
        t.Run(fmt.Sprintf("%d", k), func(t *testing.T) {
            fn := ReadyWhen(subtest.jsonPath, subtest.condition)
            require.Equal(t, subtest.expected, fn(subtest.Object))
        })
    }
}