package discov

import (
	"context"
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"google.golang.org/grpc"
	"io"
	"time"
)

type (
	// Etcd
	Etcd struct {
		cfg Config
		cli etcdv3.Client
	}

	// Service
	Service struct {
		Key       string `json:"key"`
		Val       string `json:"val"`
		Heartbeat int64  `json:"heartbeat"`
		TTL       int64  `json:"ttl"`
	}
)

// NewEtcd
func NewEtcd(cfg Config) *Etcd {
	e := &Etcd{cfg: cfg}
	_, err := e.newClient()
	if err != nil {
		panic(err)
	}
	return e
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

// GetEtcdService
func (s Service) GetEtcdService() etcdv3.Service {
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
	return etcdv3.NewRegistrar(m.cli, s.GetEtcdService(), logx.KitL())
}

// Register
func (m *Etcd) Register(s Service) error {
	return m.cli.Register(s.GetEtcdService())
}

// Deregister
func (m *Etcd) Deregister(s Service) error {
	return m.cli.Deregister(s.GetEtcdService())
}

func (m *Etcd) NewInstancer(prefix string) (*etcdv3.Instancer, error) {
	return etcdv3.NewInstancer(m.cli, prefix, logx.KitL())
}

type EndpointFactory func(conn *grpc.ClientConn) endpoint.Endpoint

func makeFactory(factory EndpointFactory) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance)
		if err != nil {
			return nil, nil, err
		}
		return factory(conn), conn, nil
	}
}

func (m *Etcd) Endpoints(ins *etcdv3.Instancer, factory EndpointFactory, retryMax, retryTimeout int) endpoint.Endpoint {
	endpointer := sd.NewEndpointer(ins, makeFactory(factory), logx.KitL())
	balancer := lb.NewRandom(endpointer, time.Now().UnixNano())
	return lb.Retry(retryMax, time.Duration(retryTimeout)*time.Millisecond, balancer)
}
