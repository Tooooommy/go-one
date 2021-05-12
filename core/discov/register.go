package discov

import (
	"google.golang.org/grpc"
	"net"
)

type (
	// Register
	Register struct {
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
func NewRegister() *Register {
	return &Register{}
}

// Option
func (r *Register) Option(option ...grpc.ServerOption) *Register {
	r.option = append(r.option, option...)
	return r
}

// StreamInterceptor
func (r *Register) StreamInterceptor(interceptors ...grpc.StreamServerInterceptor) *Register {
	r.streamInterceptor = append(r.streamInterceptor, interceptors...)
	return r
}

// UnaryInterceptor
func (r *Register) UnaryInterceptor(interceptor ...grpc.UnaryServerInterceptor) *Register {
	r.unaryInterceptor = append(r.unaryInterceptor, interceptor...)
	return r
}

// Service
func (r *Register) Service(service ...Service) *Register {
	r.service = append(r.service, service...)
	return r
}

// Register
func (r *Register) Register(Register ...RegisterFactory) *Register {
	r.register = append(r.register, Register...)
	return r
}

func (r *Register) Discovery(discovery Discovery) *Register {
	r.discovery = discovery
	return r
}

// Serve
func (r *Register) Serve(address string) error {
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
