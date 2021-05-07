package rpcx

import (
	"fmt"
	"github.com/Tooooommy/go-one/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	cfg Config
	rpc *GrpcServer
}

type ServerOption func(s *Server)

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

func WithConfig(cfg Config) ServerOption {
	return func(s *Server) {
		s.cfg = cfg
	}
}

func WithRpc(rpc *GrpcServer) ServerOption {
	return func(s *Server) {
		s.rpc = rpc
	}
}

func (s *Server) Start() error {
	// TODO: file 路径
	if s.cfg.CertFile != "" && s.cfg.KeyFile != "" {
		tls, err := credentials.NewServerTLSFromFile(s.cfg.CertFile, s.cfg.KeyFile)
		if err != nil {
			return err
		}
		s.rpc.WithOption(grpc.Creds(tls))
	}
	return s.rpc.Start()
}
