package ginx

import (
	"fmt"
	"github.com/Tooooommy/go-one/server"
	"github.com/gin-gonic/gin"
)

// Server
type Server struct {
	eng *gin.Engine
	cfg Config
}

// ServerOption
type ServerOption func(s *Server)

// NewServer: 实例化Server
func NewServer(options ...ServerOption) *Server {
	cfg := Config{
		Config: server.DefaultConfig(),
	}
	eng := gin.Default()
	s := &Server{
		eng: eng,
		cfg: cfg,
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// WithGinEngine: 设置GinEngine
func WithGinEngine(eng *gin.Engine) ServerOption {
	return func(s *Server) {
		s.eng = eng
	}
}

// WithConfig: 设置Config
func WithConfig(cfg Config) ServerOption {
	return func(s *Server) {
		s.cfg = cfg
	}
}

// Engine: 获取gin.Engine
func (s *Server) Engine() *gin.Engine {
	return s.eng
}

// Config: 获取config.HttpConfig配置
func (s *Server) Config() Config {
	return s.cfg
}

// Start: 启动服务
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	if s.cfg.HaveCert() {
		return s.eng.RunTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
	} else {
		return s.eng.Run(addr)
	}
}
