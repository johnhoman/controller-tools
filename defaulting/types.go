package defaulting

import "sigs.k8s.io/controller-runtime/pkg/client"

type Func func(client.Object)
