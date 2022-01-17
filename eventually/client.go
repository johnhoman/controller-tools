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
	"errors"
	"fmt"

	"github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	k8 client.Client
}

func (c *Client) Create(ctx context.Context, obj client.Object) gomega.AsyncAssertion {
	return gomega.Eventually(func() error {
		key := types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}
		err := c.k8.Get(ctx, key, obj)
		if client.IgnoreNotFound(err) != nil {
			return err
		}
		if apierrors.IsNotFound(err) {
			err = c.k8.Create(ctx, obj)
			if err != nil {
				return err
			}
		}
		return err
	})
}

func (c *Client) Get(ctx context.Context, key types.NamespacedName, obj client.Object) gomega.AsyncAssertion {
	return gomega.Eventually(func() error {
		return c.k8.Get(ctx, key, obj)
	})
}

func (c *Client) GetWhen(ctx context.Context, key types.NamespacedName, obj client.Object, predicate PredicateFunc) gomega.AsyncAssertion {
	return gomega.Eventually(func() error {
		err := c.k8.Get(ctx, key, obj)
		if err != nil {
			return err
		}
		if !predicate(obj) {
			return errors.New(fmt.Sprintf("predicate failed: %#v", obj))
		}
		return nil
	})
}

func (c *Client) Update(ctx context.Context, obj client.Object) gomega.AsyncAssertion {
	return gomega.Eventually(func() error {
		version := obj.GetResourceVersion()
		err := c.k8.Update(ctx, obj)
		if err != nil {
			return err
		}
		key := types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}
		err = c.k8.Get(ctx, key, obj)
		if err != nil {
			return err
		}
		if obj.GetResourceVersion() != version {
			return errors.New("waiting for update to propagate")
		}
		return nil
	})
}

func New(client client.Client) *Client {
	return &Client{k8: client}
}

var _ Eventually = &Client{}
