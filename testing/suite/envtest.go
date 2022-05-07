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

func (suite *EnvTest) Start() {
	go func() {
		defer func() {
			if e := recover(); e != nil {
				suite.Fail(e.(error).Error())
			}
		}()
		suite.Nil(suite.mgr.Start(suite.context))
	}()
	suite.NotNil(suite.mgr.NamespacedClient())
}

func (suite *EnvTest) Manager() *manager.Manager {
	return suite.mgr
}

func (suite *EnvTest) SetupSuite() {
	suite.testEnv = &envtest.Environment{}

	var err error
	suite.cfg, err = suite.testEnv.Start()
	require.Nil(suite.T(), err)
	require.NotNil(suite.T(), suite.cfg)
}

func (suite *EnvTest) SetupTest() {
	require.NotNil(suite.T(), suite.cfg)
	suite.context, suite.cancelFunc = context.WithCancel(context.Background())
	suite.mgr = manager.NewManager(suite.cfg, scheme.Scheme)
}

func (suite *EnvTest) TearDownTest() {
	suite.cancelFunc()
}

func (suite *EnvTest) TearDownSuite() {
	require.Nil(suite.T(), suite.testEnv.Stop())
}
