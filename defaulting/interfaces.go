package defaulting

import "sigs.k8s.io/controller-runtime/pkg/client"

type Default interface {
	Apply(client.Object)
}
