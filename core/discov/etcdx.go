package discov

import (
	"context"
	"errors"
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	grpctranspot "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"io"
	"time"
)

var (
	ErrNoEndpoints = errors.New("no endpoints available")
)

type (
	EncodeRequest  grpctranspot.EncodeRequestFunc
	DecodeResponse grpctranspot.DecodeResponseFunc
	ConnectFactory func(string) (*grpc.ClientConn, error)

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

	Invoker struct {
		instancer *etcdv3.Instancer
		prefix    string
		max       int
		timeout   int
		conn      ConnectFactory
		factory   sd.Factory
		endpoint  endpoint.Endpoint
		encode    EncodeRequest
		decode    DecodeResponse
		method    string
		service   string
		request   interface{}
	}

	EndpointFactory func(conn *grpc.ClientConn) endpoint.Endpoint
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

// Register
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

// Retry
func (i *Invoker) Retry(max int, timeout int) *Invoker {
	i.max = max
	i.timeout = timeout
	return i
}

// Factory
func (i *Invoker) Factory(factory sd.Factory) *Invoker {
	i.factory = factory
	return i
}

// FactoryFor
func (i *Invoker) FactoryFor(options ...grpctranspot.ClientOption) *Invoker {
	if i.conn == nil {
		i.conn = func(s string) (*grpc.ClientConn, error) {
			return grpc.Dial(s, grpc.WithInsecure())
		}
	}
	i.factory = func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := i.conn(instance)
		if err != nil {
			return nil, nil, err
		}
		return grpctranspot.NewClient(
			conn,
			i.service,
			i.method,
			grpctranspot.EncodeRequestFunc(i.encode),
			grpctranspot.DecodeResponseFunc(i.decode),
			i.request,
			options...,
		).Endpoint(), conn, nil
	}
	return i
}

// Connection
func (i *Invoker) Connection(conn ConnectFactory) *Invoker {
	i.conn = conn
	return i
}

// Request
func (i *Invoker) Request(request interface{}) {
	i.request = request
}

// Encode
func (i *Invoker) Encode(enc EncodeRequest) *Invoker {
	i.encode = enc
	return i
}

// Decode
func (i *Invoker) Decode(dec DecodeResponse) *Invoker {
	i.decode = dec
	return i
}

// Method
func (i *Invoker) Method(method string) {
	i.method = method
}

// Service
func (i *Invoker) Service(service string) {
	i.service = service
}

// Endpoint
func (i *Invoker) Endpoint() *Invoker {
	endpointer := sd.NewEndpointer(i.instancer, i.factory, logx.KitL())
	balancer := lb.NewRoundRobin(endpointer)
	if i.max > 0 && i.timeout > 0 {
		i.endpoint = lb.Retry(i.max, time.Duration(i.timeout)*time.Millisecond, balancer)
	} else {
		i.endpoint, _ = balancer.Endpoint()
	}

	return i
}

// Invoke
func (i *Invoker) Invoke(context context.Context) (interface{}, error) {
	if i.endpoint == nil {
		i.Endpoint()
	}

	if i.endpoint == nil {
		return nil, ErrNoEndpoints
	}

	return i.endpoint(context, i.request)
}
