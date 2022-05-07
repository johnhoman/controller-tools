package serviceaccount

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/require"
)

func TestImagePullSecret(t *testing.T) {
	expected := "registry-credentials"
	sa := &corev1.ServiceAccount{}
	mutator := ImagePullSecret(expected)
	require.Len(t, sa.ImagePullSecrets, 0)
	mutator.Apply(sa)
	require.Len(t, sa.ImagePullSecrets, 1)
	require.Equal(t, sa.ImagePullSecrets[0].Name, expected)
}
