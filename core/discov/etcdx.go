package discov

import (
	"context"
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

type (
	Client etcdv3.Client
	// Etcd
	Etcd struct {
		cfg Config
		cli Client
	}
)

// NewEtcd
func NewEtcd(cfg Config) (*Etcd, error) {
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

// getEtcdService
func (s Service) getEtcdService() etcdv3.Service {
	return etcdv3.Service{
		Key:   s.Key,
		Value: s.Val,
		TTL: etcdv3.NewTTLOption(
			time.Duration(s.Heartbeat),
			time.Duration(s.TTL),
		),
	}
}

// NewRegistrar
func (m *Etcd) NewRegistrar(s Service) *etcdv3.Registrar {
	return etcdv3.NewRegistrar(m.cli, s.getEtcdService(), logx.KitL())
}

// register
func (m *Etcd) Register(s Service) error {
	return m.cli.Register(s.getEtcdService())
}

// Deregister
func (m *Etcd) Deregister(s Service) error {
	return m.cli.Deregister(s.getEtcdService())
}

// NewInvoker
func (m *Etcd) NewInvoker(prefix string) (*Invoker, error) {
	instancer, err := etcdv3.NewInstancer(m.cli, prefix, logx.KitL())
	if err != nil {
		return nil, err
	}
	inv := &Invoker{instancer: instancer, prefix: prefix}
	return inv, nil
}
