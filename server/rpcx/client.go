package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/server"
	"github.com/go-kit/kit/endpoint"
)

type Client struct {
	cfg Config
	rpc *GrpcClient
}

type ClientOption func(*Client)

// NewClient
func NewClient(options ...ServerOption) *Client {
	cfg := Config{
		Config: server.DefaultConfig(),
	}
	cli := NewGrpcClient()
	return &Client{
		cfg: cfg,
		rpc: cli,
	}
}

func WithClientRpc(rpc *GrpcClient) ClientOption {
	return func(client *Client) {
		client.rpc = rpc
	}
}

func (c *Client) Endpoints(prefix string, factory discov.EndpointFactory) endpoint.Endpoint {
	return c.rpc.Endpoints(prefix, factory)
}
