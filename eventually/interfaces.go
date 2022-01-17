/*
Copyright 2022 John Homan

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
