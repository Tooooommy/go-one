package ginx

import (
	"fmt"
	"github.com/Tooooommy/go-one/server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	eng *gin.Engine
	cfg *config.HttpConfig
}

// NewServer: 实例化Server
func NewServer(config *config.HttpConfig) *Server {
	return &Server{
		eng: gin.New(),
		cfg: config,
	}
}

// GinEngine: 获取gin.Engine
func (s *Server) GinEngine() *gin.Engine {
	return s.eng
}

// Config: 获取config.HttpConfig配置
func (s *Server) Config() *config.HttpConfig {
	return s.cfg
}

// Start: 启动服务
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	if len(s.cfg.CertFile) == 0 && len(s.cfg.KeyFile) == 0 {
		return s.eng.Run(addr)
	} else {
		return s.eng.RunTLS(addr, s.cfg.CertFile, s.cfg.KeyFile)
	}
}

// ServeHTTP: 实现HTTP Serve
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.eng.ServeHTTP(w, r)
}

func wrapHs(handlers ...http.Handler) gin.HandlersChain {
	var chain gin.HandlersChain
	for _, handler := range handlers {
		chain = append(chain, gin.WrapH(handler))
	}
	return chain
}

// With: 使用http.Handler做中间件
func (s *Server) With(handlers ...http.Handler) Router {
	s.eng.Use(wrapHs(handlers...)...)
	return s
}

// Handle: http.Handler注册路由
func (s *Server) Handle(httpMethod, relativePath string, h http.Handler) {
	s.eng.Handle(httpMethod, relativePath, gin.WrapH(h))
}

// HandleFunc: http.HandleFunc注册路由
func (s *Server) HandleFunc(httpMethod, relativePath string, f http.HandlerFunc) {
	s.eng.Handle(httpMethod, relativePath, gin.WrapF(f))
}

// SetNoFound: 设置404
func (s *Server) SetNoFound(h http.Handler) {
	s.eng.NoRoute(gin.WrapH(h))
}

// SetNoMethod: 设置405
func (s *Server) SetNoMethod(h http.Handler) {
	s.eng.NoMethod(gin.WrapH(h))
}
