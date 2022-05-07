package crud

import (
	"testing"

	"github.com/stretchr/testify/require"

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
	require.Nil(suite.T(), Create(suite.Manager(), pod))
	require.NotEqual(suite.T(), "", pod.GetResourceVersion())
}

func TestSuite(t *testing.T) { suite.Run(t, &CrudSuite{}) }
