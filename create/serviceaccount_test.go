package create

import (
    "github.com/johnhoman/controller-tools/defaulting/serviceaccount"
    corev1 "k8s.io/api/core/v1"
    "testing"

    "github.com/johnhoman/controller-tools/suite"
)

type ServiceAccountCreate struct {
    suite.EnvTest
}

func (suite *ServiceAccountCreate) IsValid(sa *corev1.ServiceAccount) {
    suite.NotNil(sa)
    suite.Contains(sa.GetName(), "serviceaccount-")
    suite.Equal(suite.Manager().GetNamespace(), sa.GetNamespace())
}

func (suite *ServiceAccountCreate) TestServiceAccountCreate() {
    sa, err := ServiceAccount(suite.Manager())
    suite.Nil(err)
    suite.IsValid(sa)
}

func (suite *ServiceAccountCreate) TestServiceAccountCreateWithDefaults() {
    expected := "registry-credentials"
    sa, err := ServiceAccount(
        suite.Manager(),
        serviceaccount.ImagePullSecret(expected),
    )
    suite.Nil(err)
    suite.IsValid(sa)
    suite.Len(sa.ImagePullSecrets, 1)
    suite.Equal(expected, sa.ImagePullSecrets[0].Name)
}

func TestSuite(t *testing.T) { suite.Run(t, &ServiceAccountCreate{}) }
