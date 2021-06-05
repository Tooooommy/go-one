package ginx

import (
	"github.com/Tooooommy/go-one/server"
	"github.com/gin-gonic/gin"
)

// Server
type Server struct {
	eng *gin.Engine
	cfg *Conf
}

// ServerOption
type ServerOption func(s *Server)

// NewServer: 实例化Server
func NewServer(options ...ServerOption) *Server {
	s := &Server{
		eng: gin.New(),
		cfg: &Conf{Conf: server.DefaultConfig()},
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// WithEngine: 设置GinEngine
func WithEngine(eng *gin.Engine) ServerOption {
	return func(s *Server) {
		s.eng = eng
	}
}

// WithConf: 设置Config
func WithConf(cfg *Conf) ServerOption {
	return func(s *Server) {
		s.cfg = cfg
	}
}

// Engine: 获取gin.Engine
func (s *Server) Engine() *gin.Engine {
	return s.eng
}

// Conf: 获取config.HttpConfig配置
func (s *Server) Conf() *Conf {
	return s.cfg
}

// Start: 启动服务
func (s *Server) Start() error {
	if s.cfg.HaveCert() {
		return s.eng.RunTLS(s.cfg.Address(), s.cfg.CertFile, s.cfg.KeyFile)
	} else {
		return s.eng.Run(s.cfg.Address())
	}
}
