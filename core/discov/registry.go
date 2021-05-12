package discov

import (
	"google.golang.org/grpc"
	"net"
)

type (
	// Registry
	Registry struct {
		server            *grpc.Server
		option            []grpc.ServerOption
		streamInterceptor []grpc.StreamServerInterceptor
		unaryInterceptor  []grpc.UnaryServerInterceptor
		register          []RegisterFactory
		service           []Service
		discovery         Discovery
	}

	// Discovery
	Discovery interface {
		Register(Service) error
		Deregister(Service) error
	}

	// Service
	Service struct {
		Key       string `json:"key"`
		Val       string `json:"val"`
		Heartbeat int64  `json:"heartbeat"`
		TTL       int64  `json:"ttl"`
	}

	RegisterFactory func(*grpc.Server)
)

// NewRegister
func NewRegister() *Registry {
	return &Registry{}
}

// Option
func (r *Registry) Option(option ...grpc.ServerOption) *Registry {
	r.option = append(r.option, option...)
	return r
}

// StreamInterceptor
func (r *Registry) StreamInterceptor(interceptors ...grpc.StreamServerInterceptor) *Registry {
	r.streamInterceptor = append(r.streamInterceptor, interceptors...)
	return r
}

// UnaryInterceptor
func (r *Registry) UnaryInterceptor(interceptor ...grpc.UnaryServerInterceptor) *Registry {
	r.unaryInterceptor = append(r.unaryInterceptor, interceptor...)
	return r
}

// Service
func (r *Registry) Service(service ...Service) *Registry {
	r.service = append(r.service, service...)
	return r
}

// Registry
func (r *Registry) Register(Register ...RegisterFactory) *Registry {
	r.register = append(r.register, Register...)
	return r
}

func (r *Registry) Discovery(discovery Discovery) *Registry {
	r.discovery = discovery
	return r
}

// Serve
func (r *Registry) Serve(address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	r.option = append(r.option, grpc.ChainUnaryInterceptor(r.unaryInterceptor...))
	r.option = append(r.option, grpc.ChainStreamInterceptor(r.streamInterceptor...))
	r.server = grpc.NewServer(r.option...)

	defer func() {
		r.server.GracefulStop()
		for _, service := range r.service {
			err = r.discovery.Deregister(service)
			if err != nil {
				continue
			}
		}
	}()

	for _, service := range r.service {
		err = r.discovery.Register(service)
		if err != nil {
			return err
		}
	}

	for _, Register := range r.register {
		Register(r.server)
	}

	return r.server.Serve(lis)
}
