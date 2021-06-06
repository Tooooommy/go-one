package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type (
	// Server
	Server interface {
		Register(ServiceFactory)
		Serve(options ...grpc.ServerOption) error
	}

	// server
	server struct {
		cfg     *ServerConf
		factory ServiceFactory
	}

	ServiceFactory func(*grpc.Server)
)

// NewServer
func NewServer(cfg *ServerConf) Server {
	return &server{cfg: cfg}
}

// Register
func (s *server) Register(factory ServiceFactory) {
	s.factory = factory
}

// Serve
func (s *server) Serve(options ...grpc.ServerOption) error {
	lis, err := net.Listen("tcp", s.cfg.Address())
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
		options = append(options, grpc.Creds(tls))
	}

	//
	server := grpc.NewServer(options...)
	defer func() {
		cli.Deregister()
		server.GracefulStop()
	}()

	// 注册服务
	s.factory(server)
	return server.Serve(lis)
}
