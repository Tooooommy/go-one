package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router interface {
	With(...http.Handler) Router
	Handle(string, string, http.Handler)
	HandleFunc(string, string, http.HandlerFunc)
	SetNoFound(http.Handler)
	SetNoMethod(http.Handler)
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
