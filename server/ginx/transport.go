package ginx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

var ErrReturnIsNil = errors.New("return response is nil")

type (
	DecodeFunc func(c *gin.Context, response interface{}) httptransport.DecodeRequestFunc
	EncodeFunc func(c *gin.Context) httptransport.EncodeResponseFunc

	// Transport
	Transport struct {
		options     []httptransport.ServerOption
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
func (s *Transport) With(options ...httptransport.ServerOption) *Transport {
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
func (s *Transport) Handle(e endpoint.Endpoint, resp interface{}) gin.HandlerFunc {
	for _, middleware := range s.middlewares {
		e = middleware(e)
	}
	return func(c *gin.Context) {
		httptransport.NewServer(
			e,
			s.decode(c, resp),
			s.encode(c),
			s.options...,
		).ServeHTTP(c.Writer, c.Request)
	}
}

// NoHandle
func (s *Transport) NoHandle(e endpoint.Endpoint) gin.HandlerFunc {
	return s.Decode(NoDecoder).Encode(NoEncoder).Handle(e, nil)
}

// JSONHandle
func (s *Transport) JSONHandle(e endpoint.Endpoint, resp interface{}) gin.HandlerFunc {
	return s.Decode(JSONDecoder).Encode(JSONEncoder).Handle(e, resp)
}

// FileHandle
func (s *Transport) FileHandle(e endpoint.Endpoint, resp interface{}) gin.HandlerFunc {
	return s.Decode(FileDecoder).Encode(JSONEncoder).Handle(e, resp)
}
