package httpx

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type HTTPConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	CertFile     string `json:"cert_file"`
	KeyFile      string `json:"key_file"`
	MaxConns     int    `json:"max_conns"`
	MaxBytes     int64  `json:"max_bytes"`
	Timeout      int64  `json:"timeout"`
	CpuThreshold int64  `json:"cpu_treshold"`
}

type Server struct {
	chi.Router
	cfg *HTTPConfig
}

func NewServer(config *HTTPConfig) *Server {
	return &Server{
		Router: chi.NewRouter(),
		cfg:    config,
	}
}

func (s *Server) GetConfig() *HTTPConfig {
	return s.cfg
}

func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	if len(s.cfg.CertFile) == 0 && len(s.cfg.KeyFile) == 0 {
		return http.ListenAndServe(addr, s)
	} else {
		return http.ListenAndServeTLS(addr, s.cfg.CertFile, s.cfg.KeyFile, s)
	}
}
