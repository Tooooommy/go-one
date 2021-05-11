package rpcx

import (
	"fmt"
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
func NewServer(cfg Config, options ...ServerOption) *Server {
	if cfg.Host == "" {
		cfg.Host = "0.0.0.0"
	}
	if cfg.Port <= 0 {
		cfg.Port = 9080
	}
	rpc := NewGrpcServer(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
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
	err := s.EnableCert()
	if err != nil {
		return err
	}
	err = s.rpc.EnableEtcd(s.cfg.Discov)
	if err != nil {
		return err
	}
	defer s.rpc.Stop()
	return s.rpc.Start()
}

// EnableCert
func (s *Server) EnableCert() error {
	if s.cfg.HaveCert() { // 验证
		tls, err := credentials.NewServerTLSFromFile(s.cfg.CertFile, s.cfg.KeyFile)
		if err != nil {
			return err
		}
		s.rpc.WithOption(grpc.Creds(tls))
	}
	return nil
}
