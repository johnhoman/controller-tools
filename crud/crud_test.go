package crud

import (
	"context"
	"github.com/johnhoman/controller-tools"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	"github.com/johnhoman/controller-tools/prefab"
)

type CrudSuite struct {
	suite.Suite
	cfg        *rest.Config
	testEnv    *envtest.Environment
	context    context.Context
	cancelFunc context.CancelFunc
	namespace  string
	mgr        *controllertools.Manager
}

func (suite *CrudSuite) SetupTest() {
	suite.testEnv = &envtest.Environment{}

	var err error
	suite.cfg, err = suite.testEnv.Start()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), suite.cfg)

	suite.context, suite.cancelFunc = context.WithCancel(context.Background())
	suite.mgr = controllertools.NewManager(suite.cfg, scheme.Scheme)

	go func() {
		defer func() {
			if e := recover(); e != nil {
				require.Fail(suite.T(), e.(error).Error())
			}
		}()
		require.Nil(suite.T(), suite.mgr.Start(suite.context))
	}()
    require.NotNil(suite.T(), suite.mgr.NamespacedClient())
    require.Nil(suite.T(), err)
}

func (suite *CrudSuite) TearDownTest() {
    suite.cancelFunc()
	require.Nil(suite.T(), suite.testEnv.Stop())
}

func (suite *CrudSuite) TestCreate() {
	pod := prefab.NewPod(
        prefab.InNamespace(suite.mgr.GetNamespace()),
        prefab.Nginx(),
    )
	require.Nil(suite.T(), Create(suite.mgr, pod))
	require.NotEqual(suite.T(), "", pod.GetResourceVersion())
}

func TestSuite(t *testing.T) { suite.Run(t, &CrudSuite{}) }
