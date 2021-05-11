package rpcx

import (
	"context"
	"github.com/Tooooommy/go-one/core/discov"
)

type Client struct {
	cfg Config
	rpc *GrpcClient
}

type ClientOption func(*Client)

// NewClient
func NewClient(cfg Config, options ...ServerOption) (*Client, error) {
	cli, err := NewGrpcClient(cfg.Discov)
	if err != nil {
		return nil, err
	}
	client := &Client{
		cfg: cfg,
		rpc: cli,
	}
	return client, nil
}

func WithClientCof(cfg Config) ClientOption {
	return func(client *Client) {
		client.cfg = cfg
	}
}

func WithClientRpc(rpc *GrpcClient) ClientOption {
	return func(client *Client) {
		client.rpc = rpc
	}
}

func (c *Client) Invoke(ctx context.Context, in interface{}, prefix string, factory discov.EndpointFactory) (interface{}, error) {
	return c.rpc.Invoke(prefix, factory)(ctx, in)
}
