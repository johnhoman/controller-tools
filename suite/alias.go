package suite

import (
    "github.com/stretchr/testify/suite"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "strings"
)

var (
	Run = suite.Run
    _ = Run
)

type Suite struct {
    suite.Suite
}

type metadata struct {
    suite *Suite
    metav1.Object
}

func (m *metadata) HasName(expected string) bool {
    return m.suite.Equal(m.GetName(), expected)
}

func (m *metadata) NameHasPrefix(expected string) bool {
    if !strings.HasPrefix(m.GetName(), expected) {
        m.suite.Fail(
            "%s does not start with %s",
            m.GetName(),
            expected,
        )
        return false
    }
    return true
}

func (m *metadata) HasAnnotation(expectedKey, expectedValue string) bool {
    ok := m.suite.Contains(m.GetAnnotations(), expectedKey)
    if !ok {
        return ok
    }

    return m.suite.Equal(expectedValue, m.GetAnnotations()[expectedKey])
}

func (m *metadata) HasLabel(expectedKey, expectedValue string) bool {
    ok := m.suite.Contains(m.GetLabels(), expectedKey)
    if !ok {
        return ok
    }

    return m.suite.Equal(expectedValue, m.GetLabels()[expectedKey])
}


func (suite *Suite) Metadata(obj metav1.Object) *metadata {
    return &metadata{suite: suite, Object: obj}
}
