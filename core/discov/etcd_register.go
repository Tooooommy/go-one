package discov

import (
	"google.golang.org/grpc"
	"net"
)

type (
	register struct {
		server            *grpc.Server
		option            []grpc.ServerOption
		streamInterceptor []grpc.StreamServerInterceptor
		unaryInterceptor  []grpc.UnaryServerInterceptor
		register          []RegisterFactory
		service           []Service
		before            []BeforeFactory
		finalizer         []FinalizerFactory
	}

	Register interface {
		Option(...grpc.ServerOption) Register
		StreamInterceptor(...grpc.StreamServerInterceptor) Register
		UnaryInterceptor(...grpc.UnaryServerInterceptor) Register
		Service(...Service) Register
		Register(...RegisterFactory) Register
		Before(...BeforeFactory) Register
		Finalizer(...FinalizerFactory) Register
		Serve(string) error
	}

	// Service
	Service struct {
		Key       string `json:"key"`
		Val       string `json:"val"`
		Heartbeat int64  `json:"heartbeat"`
		TTL       int64  `json:"ttl"`
	}

	BeforeFactory    func(r *register)
	FinalizerFactory func(r *register)
	RegisterFactory  func(*grpc.Server)
)

func NewRegister() Register {
	return &register{}
}

func (r *register) Option(option ...grpc.ServerOption) Register {
	r.option = append(r.option, option...)
	return r
}

func (r *register) StreamInterceptor(interceptors ...grpc.StreamServerInterceptor) Register {
	r.streamInterceptor = append(r.streamInterceptor, interceptors...)
	return r
}

func (r *register) UnaryInterceptor(interceptor ...grpc.UnaryServerInterceptor) Register {
	r.unaryInterceptor = append(r.unaryInterceptor, interceptor...)
	return r
}
func (r *register) Service(service ...Service) Register {
	r.service = append(r.service, service...)
	return r
}

func (r *register) Register(register ...RegisterFactory) Register {
	r.register = append(r.register, register...)
	return r
}

func (r *register) Before(before ...BeforeFactory) Register {
	r.before = append(r.before, before...)
	return r
}

func (r *register) Finalizer(finalizer ...FinalizerFactory) Register {
	r.finalizer = append(r.finalizer, finalizer...)
	return r
}

func (r *register) Serve(address string) error {
	defer func() {
		r.server.GracefulStop()

		for _, finalizer := range r.finalizer {
			finalizer(r)
		}
	}()

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	r.option = append(r.option, grpc.ChainUnaryInterceptor(r.unaryInterceptor...))
	r.option = append(r.option, grpc.ChainStreamInterceptor(r.streamInterceptor...))
	r.server = grpc.NewServer(r.option...)
	for _, register := range r.register {
		register(r.server)
	}

	for _, before := range r.before {
		before(r)
	}
	return r.server.Serve(lis)
}

func RegisterEtcd(cli *Etcd) BeforeFactory {
	return func(r *register) {
		for _, service := range r.service {
			_ = cli.Register(service)
		}
	}
}

func DeregisterEtcd(cli *Etcd) FinalizerFactory {
	return func(r *register) {
		for _, service := range r.service {
			_ = cli.Deregister(service)
		}
	}
}
