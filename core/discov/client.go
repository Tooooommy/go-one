package discov

import (
	"context"
	"github.com/Tooooommy/go-one/core/syncx"
	"github.com/go-kit/kit/sd/etcdv3"
)

type (
	Client struct {
		cfg *Config
	}
)

var (
	manager = syncx.NewManager()
)

func NewClient(cfg *Config) *Client {
	client := &Client{cfg: cfg}
	return client
}

func (c *Client) getClient() (etcdv3.Client, error) {
	val, ok := manager.Get(c.cfg.Key)
	if !ok {
		cli, err := etcdv3.NewClient(
			context.Background(),
			c.cfg.Hosts,
			c.cfg.ClientOptions(),
		)
		if err != nil {
			return nil, err
		}
		manager.Set(c.cfg.Key, cli)
		return cli, nil
	}
	return val.(etcdv3.Client), nil
}
