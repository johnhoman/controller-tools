package eventually

import (
	"context"
	"github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PredicateFunc func(client.Object) bool

type Eventually interface {
	Create(ctx context.Context, obj client.Object) gomega.AsyncAssertion
	Update(ctx context.Context, obj client.Object) gomega.AsyncAssertion
	Get(ctx context.Context, key types.NamespacedName, obj client.Object) gomega.AsyncAssertion
	GetWhen(ctx context.Context, key types.NamespacedName, obj client.Object, predicateFunc PredicateFunc) gomega.AsyncAssertion
}
