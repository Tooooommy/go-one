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
	if c.cfg.Discov.HaveEtcd() {
		etcd, err := etcdx.NewEtcd(c.cfg.Discov)
		if err != nil {
			return nil, err
		}
		instancer, err := etcd.NewInstancer(key)
		if err != nil {
			return nil, err
		}
		inv.Instancer(instancer)
	} else {
		inv.Address(c.cfg.Discov.Hosts...)
	}
	return inv, nil
}
