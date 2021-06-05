package rpcx

import (
	"fmt"
	"github.com/Tooooommy/go-one/core/discov"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type (
	// Server
	Server struct {
		cfg     *ServerConf
		service []ServiceFactory
		options []grpc.ServerOption
	}

	ServiceFactory func(*ServerConf, *grpc.Server)
)

// NewServer
func NewServer(cfg *ServerConf, options ...grpc.ServerOption) *Server {
	svr := &Server{
		cfg:     cfg,
		options: options,
	}
	return svr
}

// ServerConf
func (s *Server) Config() *ServerConf {
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

	// Etcd服务发现
	cli := discov.NewRegistry(&s.cfg.Etcd)
	err = cli.Register()
	if err != nil {
		return err
	}

	// TLS拦截器
	if s.cfg.HaveCert() {
		tls, err := credentials.NewServerTLSFromFile(s.cfg.CertFile, s.cfg.KeyFile)
		if err != nil {
			return err
		}
		s.options = append(s.options, grpc.Creds(tls))
	}

	//
	server := grpc.NewServer(s.options...)
	defer func() {
		cli.Deregister()
		server.GracefulStop()
	}()

	// 注册服务
	for _, service := range s.service {
		service(s.cfg, server)
	}

	return server.Serve(lis)
}
