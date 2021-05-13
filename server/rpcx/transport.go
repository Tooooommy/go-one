package rpcx

import (
	"github.com/go-kit/kit/endpoint"
	grpctranspot "github.com/go-kit/kit/transport/grpc"
)

type (
	EncodeFunc grpctranspot.EncodeRequestFunc
	DecodeFunc grpctranspot.DecodeResponseFunc

	Transport struct {
		options     []grpctranspot.ServerOption
		middlewares []endpoint.Middleware
		encode      EncodeFunc
		decode      DecodeFunc
	}
)

// NewTransport
func NewTransport() *Transport {
	return &Transport{}
}

// With
func (s *Transport) With(options ...grpctranspot.ServerOption) *Transport {
	s.options = append(s.options, options...)
	return s
}

// Use
func (s *Transport) Use(middlewares ...endpoint.Middleware) *Transport {
	s.middlewares = append(s.middlewares, middlewares...)
	return s
}

// Encode
func (s *Transport) Encode(enc EncodeFunc) *Transport {
	s.encode = enc
	return s
}

// Decode
func (s *Transport) Decode(dec DecodeFunc) *Transport {
	s.decode = dec
	return s
}

// Handle
func (s *Transport) Handle(e endpoint.Endpoint) grpctranspot.Handler {
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
