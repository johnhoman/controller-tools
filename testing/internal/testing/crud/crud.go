package crud

import (
	"context"
	"fmt"
	"github.com/johnhoman/controller-tools/defaulting"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

const defaultTimeout = time.Second

type Manager interface {
	GetClient() client.Client
	GetContext() context.Context
	GetNamespace() string
	GetScheme() *runtime.Scheme
}

func Create(mgr Manager, obj client.Object, options ...interface{}) error {
	timeout := defaultTimeout
	opts := make([]defaulting.Default, 0)
	for _, option := range options {
		switch opt := option.(type) {
		case time.Duration:
			timeout = opt
		case defaulting.Default:
			opts = append(opts, opt)
		}
	}
	for _, opt := range opts {
		opt.Apply(obj)
	}
	cli := mgr.GetClient()
	err := cli.Create(mgr.GetContext(), obj)
	if err != nil {
		return err
	}

	doGet := func() error {
		if err := cli.Get(mgr.GetContext(), client.ObjectKeyFromObject(obj), obj); err != nil {
			return err
		}
		return nil
	}

	now := time.Now()
	nextSleep := time.Millisecond
	for {
		if err := doGet(); err != nil {
			fmt.Println(err.Error())
			nextSleep *= 2
			if now.Add(nextSleep).Before(now.Add(timeout)) {
				time.Sleep(nextSleep)
				continue
			}
			return err
		}
		break
	}
	return nil
}
