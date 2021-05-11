package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	address            string
	server             *grpc.Server
	etcd               *discov.Etcd
	services           []discov.Service
	options            []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor
	register           []RegisterFunc
}

type RegisterFunc func(*grpc.Server)

func NewGrpcServer(address string) *GrpcServer {
	return &GrpcServer{
		address: address,
	}
}

func (s *GrpcServer) WithOption(option ...grpc.ServerOption) {
	s.options = append(s.options, option...)
}

func (s *GrpcServer) UseStreamInterceptor(interceptor ...grpc.StreamServerInterceptor) {
	s.streamInterceptors = append(s.streamInterceptors, interceptor...)
}

func (s *GrpcServer) UseUnaryInterceptors(interceptor ...grpc.UnaryServerInterceptor) {
	s.unaryInterceptors = append(s.unaryInterceptors, interceptor...)
}

func (s *GrpcServer) EnableEtcd(cfg discov.Config) (err error) {
	if cfg.HaveEtcd() {
		s.etcd, err = discov.NewEtcd(cfg)
	}
	return
}

func (s *GrpcServer) UseEtcdService(service ...discov.Service) {
	s.services = append(s.services, service...)
}

func (s *GrpcServer) Register(register ...RegisterFunc) {
	s.register = append(s.register, register...)
}

func (s *GrpcServer) Start() error {
	if s.etcd != nil {
		for _, service := range s.services {
			err := s.etcd.Register(service)
			if err != nil {
				return err
			}
		}
	}

	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	s.options = append(s.options, grpc.ChainUnaryInterceptor(s.unaryInterceptors...))
	s.options = append(s.options, grpc.ChainStreamInterceptor(s.streamInterceptors...))
	s.server = grpc.NewServer(s.options...)
	for _, register := range s.register {
		register(s.server)
	}
	return s.server.Serve(lis)
}

func (s *GrpcServer) Stop() error {
	if s.server != nil {
		s.server.GracefulStop()
	}
	if s.etcd != nil {
		for _, service := range s.services {
			return s.etcd.Deregister(service)
		}
	}
	return nil
}
