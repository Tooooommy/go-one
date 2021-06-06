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
	val, ok := manager.Get(c.cfg.Name)
	if !ok {
		ctx := context.Background()
		opt := etcdv3.ClientOptions{
			Username: c.cfg.Username,
			Password: c.cfg.Password,
		}
		cli, err := etcdv3.NewClient(ctx, c.cfg.Hosts, opt)
		if err != nil {
			return nil, err
		}
		manager.Set(c.cfg.Name, cli)
		return cli, nil
	}
	return val.(etcdv3.Client), nil
}
