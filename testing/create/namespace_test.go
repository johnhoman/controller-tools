package create

import (
	"github.com/johnhoman/controller-tools/defaulting"
	suite2 "github.com/johnhoman/controller-tools/testing/suite"
	v1 "k8s.io/api/core/v1"
	"testing"
)

type NamespaceCreate struct{ suite2.EnvTest }

func (suite *NamespaceCreate) TestNamespaceCreate() {
	ns, err := Namespace(suite.Manager())
	suite.Nil(err)
	suite.Contains(ns.GetName(), "namespace-")
	suite.Equal(v1.NamespaceActive, ns.Status.Phase)
}

func (suite *NamespaceCreate) TestNamespaceCreateWithOptions() {
	ns, err := Namespace(
		suite.Manager(),
		defaulting.WithAnnotation("a.b.c", "true"),
		defaulting.WithLabels("aaa.bbb.cc", "false"),
	)
	suite.Nil(err)
	suite.Equal(v1.NamespaceActive, ns.Status.Phase)
	suite.Metadata(ns).NameHasPrefix("namespace")
	suite.Metadata(ns).HasAnnotation("a.b.c", "true")
	suite.Metadata(ns).HasLabel("aaa.bbb.cc", "false")
}

func TestNamespaceCreate(t *testing.T) { suite2.Run(t, &NamespaceCreate{}) }
