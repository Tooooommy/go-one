package etcdx

import (
	"context"
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

type (
	ClientV3 etcdv3.Client
	// Client
	Client struct {
		cfg discov.Config
		cli ClientV3
	}
)

// NewEtcd
func NewEtcd(cfg discov.Config) (*Client, error) {
	options := cfg.GetEtcdClientOptions()
	cli, err := etcdv3.NewClient(context.Background(), cfg.Hosts, options)
	return &Client{cfg: cfg, cli: cli}, err
}

// register
func (m *Client) Register(s discov.Service) error {
	return m.cli.Register(etcdv3.Service{
		Key:   s.Key,
		Value: s.Val,
		TTL: etcdv3.NewTTLOption(
			time.Duration(s.Heartbeat),
			time.Duration(s.TTL),
		),
	})
}

// Deregister
func (m *Client) Deregister(s discov.Service) error {
	return m.cli.Deregister(etcdv3.Service{
		Key:   s.Key,
		Value: s.Val,
		TTL: etcdv3.NewTTLOption(
			time.Duration(s.Heartbeat),
			time.Duration(s.TTL),
		),
	})
}

func (m *Client) NewInstancer(key string) (sd.Instancer, error) {
	return etcdv3.NewInstancer(m.cli, key, logx.KitL())
}
