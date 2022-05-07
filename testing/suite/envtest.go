package suite

import (
	"context"
	"github.com/johnhoman/controller-tools/testing/manager"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes/scheme"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

type EnvTest struct {
	Suite
	cfg        *rest.Config
	testEnv    *envtest.Environment
	context    context.Context
	cancelFunc context.CancelFunc
	namespace  string
	mgr        *manager.Manager
}

func (suite *EnvTest) Manager() *manager.Manager {
	return suite.mgr
}

func (suite *EnvTest) SetupTest() {
	suite.testEnv = &envtest.Environment{}

	var err error
	suite.cfg, err = suite.testEnv.Start()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), suite.cfg)

	suite.context, suite.cancelFunc = context.WithCancel(context.Background())
	suite.mgr = manager.NewManager(suite.cfg, scheme.Scheme)

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

func (suite *EnvTest) TearDownTest() {
	suite.cancelFunc()
	require.Nil(suite.T(), suite.testEnv.Stop())
}
