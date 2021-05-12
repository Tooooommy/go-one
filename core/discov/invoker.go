package discov

import (
	"context"
	"errors"
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
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

	Invoker struct {
		instancer sd.Instancer
		factory   sd.Factory
		conn      ConnectFactory
		endpoint  endpoint.Endpoint
		encode    EncodeRequest
		decode    DecodeResponse
		max       int
		timeout   int
		address   []string
		method    string
		service   string
		request   interface{}
	}
)

// NewInvoker
func NewInvoker() *Invoker {
	return &Invoker{}
}

// Instancer
func (i *Invoker) Instancer(instancer sd.Instancer) *Invoker {
	i.instancer = instancer
	return i
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

// Address
func (i *Invoker) Address(addr ...string) *Invoker {
	i.address = addr
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
func (i *Invoker) Endpoint() endpoint.Endpoint {
	if i.instancer == nil {
		i.instancer = sd.FixedInstancer(i.address)
	}

	endpointer := sd.NewEndpointer(i.instancer, i.factory, logx.KitL())
	balancer := lb.NewRoundRobin(endpointer)
	if i.max > 0 && i.timeout > 0 {
		i.endpoint = lb.Retry(i.max, time.Duration(i.timeout)*time.Millisecond, balancer)
	} else {
		i.endpoint, _ = balancer.Endpoint()
	}

	return i.endpoint
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
