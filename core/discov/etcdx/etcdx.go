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
	Client etcdv3.Client
	// Etcd
	Etcd struct {
		cfg discov.Config
		cli Client
	}
)

// NewEtcd
func NewEtcd(cfg discov.Config) (*Etcd, error) {
	e := &Etcd{cfg: cfg}
	_, err := e.newClient()
	return e, err
}

// newClient
func (m *Etcd) newClient() (etcdv3.Client, error) {
	if m.cli == nil {
		options := m.cfg.GetEtcdClientOptions()
		cli, err := etcdv3.NewClient(context.Background(), m.cfg.Hosts, options)
		if err != nil {
			return nil, err
		}
		m.cli = cli
	}
	return m.cli, nil
}

// register
func (m *Etcd) Register(s discov.Service) error {
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
func (m *Etcd) Deregister(s discov.Service) error {
	return m.cli.Deregister(etcdv3.Service{
		Key:   s.Key,
		Value: s.Val,
		TTL: etcdv3.NewTTLOption(
			time.Duration(s.Heartbeat),
			time.Duration(s.TTL),
		),
	})
}

func (m *Etcd) NewInstancer(key string) (sd.Instancer, error) {
	return etcdv3.NewInstancer(m.cli, key, logx.KitL())
}
