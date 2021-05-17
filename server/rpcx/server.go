package rpcx

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type (
	// Server
	Server struct {
		cfg     Config
		options []grpc.ServerOption
		service []ServiceFactory
	}

	ServiceFactory func(*grpc.Server)
)

// NewServer
func NewServer(cfg Config, options ...grpc.ServerOption) *Server {
	svr := &Server{
		cfg:     cfg,
		options: options,
	}
	return svr
}

// Config
func (s *Server) Config() Config {
	return s.cfg
}

// StreamInterceptor
func (s *Server) StreamInterceptor(interceptors ...grpc.StreamServerInterceptor) {
	s.options = append(s.options, grpc.ChainStreamInterceptor(interceptors...))
}

// UnaryInterceptor
func (s *Server) UnaryInterceptor(interceptor ...grpc.UnaryServerInterceptor) {
	s.options = append(s.options, grpc.ChainUnaryInterceptor(interceptor...))
}

// Register
func (s *Server) Register(service ...ServiceFactory) {
	s.service = append(s.service, service...)
}

// Start
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	if s.cfg.HaveCert() { // 验证
		tls, err := credentials.NewServerTLSFromFile(s.cfg.CertFile, s.cfg.KeyFile)
		if err != nil {
			return err
		}
		s.options = append(s.options, grpc.Creds(tls))
	}

	server := grpc.NewServer(s.options...)
	defer server.GracefulStop()

	return server.Serve(lis)
}
