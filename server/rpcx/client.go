package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"sync"
)

type Client struct {
	etcd *discov.Etcd
	invs sync.Map
}

type ClientOption func(*Client)

// NewClient
func NewClient(cfg Config) (*Client, error) {
	cli, err := discov.NewEtcd(cfg.Discov)
	if err != nil {
		return nil, err
	}
	return &Client{etcd: cli}, nil
}

// 加上读写时
func (c *Client) Invoker(prefix string) (*discov.Invoker, error) {
	if val, ok := c.invs.Load(prefix); ok {
		return val.(*discov.Invoker), nil
	} else {
		ins, err := c.etcd.NewInvoker(prefix)
		if err != nil {
			return nil, err
		}
		c.invs.Store(prefix, ins)
		return ins, err
	}
}
