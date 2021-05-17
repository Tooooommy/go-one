package rpcx

import (
	"github.com/go-kit/kit/endpoint"
	grpctranspot "github.com/go-kit/kit/transport/grpc"
)

type (
	EncodeFunc grpctranspot.EncodeRequestFunc
	DecodeFunc grpctranspot.DecodeResponseFunc

	Service struct {
		options     []grpctranspot.ServerOption
		middlewares []endpoint.Middleware
		encode      EncodeFunc
		decode      DecodeFunc
	}
)

// NewService
func NewService() *Service {
	return &Service{}
}

// With
func (s *Service) With(options ...grpctranspot.ServerOption) *Service {
	s.options = append(s.options, options...)
	return s
}

// Use
func (s *Service) Use(middlewares ...endpoint.Middleware) *Service {
	s.middlewares = append(s.middlewares, middlewares...)
	return s
}

// Encode
func (s *Service) Encode(enc EncodeFunc) *Service {
	s.encode = enc
	return s
}

// Decode
func (s *Service) Decode(dec DecodeFunc) *Service {
	s.decode = dec
	return s
}

// Handle
func (s *Service) Handle(e endpoint.Endpoint) grpctranspot.Handler {
	for _, middleware := range s.middlewares {
		e = middleware(e)
	}
	return grpctranspot.NewServer(
		e,
		grpctranspot.DecodeRequestFunc(s.decode),
		grpctranspot.EncodeResponseFunc(s.encode),
		s.options...,
	)
}
