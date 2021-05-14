package rpcx

import (
	"fmt"
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/core/discov/etcdx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server
type Server struct {
	cfg Config
	reg *discov.Registry
}

// ServerOption
type ServerOption func(s *Server)

// NewServer
func NewServer(cfg Config, options ...ServerOption) *Server {
	reg := discov.NewRegister()
	svr := &Server{
		cfg: cfg,
		reg: reg,
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

// WithRegister
func WithRegister(reg *discov.Registry) ServerOption {
	return func(s *Server) {
		s.reg = reg
	}
}

// Start
func (s *Server) Start() error {
	if s.cfg.HaveCert() { // 验证
		tls, err := credentials.NewServerTLSFromFile(s.cfg.CertFile, s.cfg.KeyFile)
		if err != nil {
			return err
		}
		s.reg.Option(grpc.Creds(tls))
	}

	if s.cfg.Etcd.HaveEtcd() {
		cli, err := etcdx.NewClient(s.cfg.Etcd)
		if err != nil {
			return err
		}
		s.reg.Discovery(cli)
	}
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	return s.reg.Serve(addr)
}
