package rpcx

import (
	"fmt"
	"github.com/Tooooommy/go-one/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server
type Server struct {
	cfg Config
	rpc *GrpcServer
}

// ServerOption
type ServerOption func(s *Server)

// NewServer
func NewServer(options ...ServerOption) *Server {
	cfg := Config{
		Config: server.DefaultConfig(),
	}
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	rpc := NewGrpcServer(address)
	svr := &Server{
		cfg: cfg,
		rpc: rpc,
	}
	for _, option := range options {
		option(svr)
	}
	return svr
}

// WithConfig
func WithConfig(cfg Config) ServerOption {
	return func(s *Server) {
		s.cfg = cfg
	}
}

// WithServerRpc
func WithServerRpc(rpc *GrpcServer) ServerOption {
	return func(s *Server) {
		s.rpc = rpc
	}
}

// Start
func (s *Server) Start() error {
	err := s.enableCert()
	if err != nil {
		return err
	}
	s.rpc.EnableEtcd(s.cfg.Discov)
	defer s.rpc.Stop()
	return s.rpc.Start()
}

// enableCert
func (s *Server) enableCert() error {
	if s.cfg.HaveCert() { // 验证
		tls, err := credentials.NewServerTLSFromFile(s.cfg.CertFile, s.cfg.KeyFile)
		if err != nil {
			return err
		}
		s.rpc.WithOption(grpc.Creds(tls))
	}
	return nil
}
