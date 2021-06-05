package rpcx

import (
	"context"
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/core/grpcx"
	"time"
)

type (
	Client struct {
		cfg *ClientConf
	}

	ClientOption func(c *ClientConf)
)

// NewClient
func NewClient(cfg *ClientConf, options ...ClientOption) *Client {
	for _, opt := range options {
		opt(cfg)
	}
	client := &Client{cfg: cfg}
	return client
}

// ClientRetries
func ClientRetries(retries int) ClientOption {
	return func(c *ClientConf) {
		c.Retries = retries
	}
}

// ClientTimeout
func ClientTimeout(timeout int64) ClientOption {
	return func(c *ClientConf) {
		c.Timeout = timeout
	}
}

// ClientToken
func ClientToken(token string) ClientOption {
	return func(c *ClientConf) {
		c.Token = token
	}
}

// Invoke
func (c *Client) Invoke(ctx context.Context, invoker grpcx.Invoker, request interface{}) (interface{}, error) {
	instancer, err := discov.NewClient(&c.cfg.Etcd).NewInstancer(c.cfg.Address)
	if err != nil {
		return nil, err
	}
	return invoker.Invoke(ctx, instancer, c.cfg.Retries, time.Duration(c.cfg.Timeout), request)
}
