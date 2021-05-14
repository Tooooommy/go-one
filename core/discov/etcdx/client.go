package etcdx

import (
	"context"
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

type (
	client etcdv3.Client
	// Client
	Client struct {
		cfg Config
		cli client
	}
)

// NewClient
func NewClient(cfg Config) (*Client, error) {
	options := cfg.ClientOptions()
	cli, err := etcdv3.NewClient(context.Background(), cfg.Hosts, options)
	return &Client{
		cfg: cfg,
		cli: cli,
	}, err
}

// register
func (c *Client) Register(s discov.Service) error {
	return c.cli.Register(etcdv3.Service{
		Key:   s.Key,
		Value: s.Val,
		TTL: etcdv3.NewTTLOption(
			time.Duration(s.Heartbeat),
			time.Duration(s.TTL),
		),
	})
}

// Deregister
func (c *Client) Deregister(s discov.Service) error {
	return c.cli.Deregister(etcdv3.Service{
		Key:   s.Key,
		Value: s.Val,
		TTL: etcdv3.NewTTLOption(
			time.Duration(s.Heartbeat),
			time.Duration(s.TTL),
		),
	})
}

func (c *Client) NewInstancer(key string) (sd.Instancer, error) {
	return etcdv3.NewInstancer(c.cli, key, zapx.KitL())
}
