package predicate

import (
    "strings"

    "k8s.io/apimachinery/pkg/util/json"
    "sigs.k8s.io/controller-runtime/pkg/client"
)

// ReadyWhen returns the provided wrapped predicate condition, passing the object specified
// by jsonPath to condition
func ReadyWhen(jsonPath string, condition func(v interface{}) bool) func(client.Object) bool {
    return func(obj client.Object) bool {
        parts := strings.Split(jsonPath, ".")
        raw, err := json.Marshal(obj)
        if err != nil {
            panic(err)
        }
        m := make(map[string]interface{})
        if err := json.Unmarshal(raw, &m); err != nil {
            panic(err)
        }
        for _, part := range parts[:len(parts)-1] {
            m = m[part].(map[string]interface{})
        }
        v := m[parts[len(parts)-1]]
        return condition(v)
    }
}