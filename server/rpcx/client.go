package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/core/discov/etcdx"
)

type Client struct {
	cfg Config
}

// NewClient
func NewClient(cfg Config) *Client {
	return &Client{cfg: cfg}
}

// Invoker
func (c *Client) Invoker(key string) (*discov.Invoker, error) {
	inv := discov.NewInvoker()
	if c.cfg.Etcd.HaveEtcd() {
		etcd, err := etcdx.NewClient(c.cfg.Etcd)
		if err != nil {
			return nil, err
		}
		instancer, err := etcd.NewInstancer(key)
		if err != nil {
			return nil, err
		}
		inv.Instancer(instancer)
	}
	return inv, nil
}
