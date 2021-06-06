package rpcx

import (
	"context"
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/core/grpcx"
	"time"
)

type (
	Client interface {
		Invoke(ctx context.Context, invoker grpcx.Invoker, request interface{}) (interface{}, error)
	}

	client struct {
		cfg *ClientConf
	}
)

// NewClient
func NewClient(cfg *ClientConf) Client {
	client := &client{cfg: cfg}
	return client
}

// Invoke
func (c *client) Invoke(ctx context.Context, invoker grpcx.Invoker, request interface{}) (interface{}, error) {
	instancer, err := discov.NewClient(&c.cfg.Etcd).NewInstancer(c.cfg.Address)
	if err != nil {
		return nil, err
	}
	return invoker.Invoke(ctx, instancer, c.cfg.Retries, time.Duration(c.cfg.Timeout), request)
}
