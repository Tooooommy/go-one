package ginx

import (
	"github.com/gin-gonic/gin"
)

type (
	Server interface {
		Register(eng *gin.Engine)
		Start() error
	}

	// server
	server struct {
		eng *gin.Engine
		cfg *ServerConf
	}
)

// NewServer: 实例化Server
func NewServer(cfg *ServerConf) *server {
	s := &server{eng: gin.Default(), cfg: cfg}
	return s
}

// WithEngine: 设置GinEngine
func (s *server) Register(eng *gin.Engine) {
	s.eng = eng
}

// Start: 启动服务
func (s *server) Start() error {
	if s.cfg.HaveCert() {
		return s.eng.RunTLS(s.cfg.Address(), s.cfg.CertFile, s.cfg.KeyFile)
	} else {
		return s.eng.Run(s.cfg.Address())
	}
}
