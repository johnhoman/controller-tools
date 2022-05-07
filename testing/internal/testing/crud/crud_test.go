package crud

import (
	"testing"

	"github.com/johnhoman/controller-tools/prefab"
	"github.com/johnhoman/controller-tools/testing/suite"
)

type CrudSuite struct {
	suite.EnvTest
}

func (suite *CrudSuite) TestCreate() {
	pod := prefab.NewPod(
		prefab.InNamespace(suite.Manager().GetNamespace()),
		prefab.Nginx(),
	)
	suite.Require().Nil(Create(suite.Manager(), pod))
	suite.Require().NotEqual("", pod.GetResourceVersion())
}

func TestSuite(t *testing.T) { suite.Run(t, &CrudSuite{}) }
