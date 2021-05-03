package ginx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

var ErrReturnIsNil = errors.New("return response is nil")

type DecodeFunc func(c *gin.Context, request interface{}) httptransport.DecodeRequestFunc
type EncodeFunc func(c *gin.Context) httptransport.EncodeResponseFunc

// Transport
type Transport struct {
	options     []httptransport.ServerOption
	middlewares []endpoint.Middleware
}

// NewTransport
func NewTransport() *Transport {
	return &Transport{
		options: []httptransport.ServerOption{},
	}
}

// WithOptions
func (s *Transport) With(options ...httptransport.ServerOption) *Transport {
	s.options = options
	return s
}

func (s *Transport) Use(middlewares ...endpoint.Middleware) *Transport {
	s.middlewares = middlewares
	return s
}

// Handler
func (s *Transport) Handler(e endpoint.Endpoint, request interface{}, dec DecodeFunc, enc EncodeFunc) gin.HandlerFunc {
	for _, middleware := range s.middlewares {
		e = middleware(e)
	}
	return func(c *gin.Context) {
		httptransport.NewServer(
			e,
			dec(c, request),
			enc(c),
			s.options...,
		).ServeHTTP(c.Writer, c.Request)
	}
}

// NoHandler
func (s *Transport) NoHandler(e endpoint.Endpoint) gin.HandlerFunc {
	return s.Handler(e, nil, NoDecoder, NoEncoder)
}

// JSONHandler
func (s *Transport) JSONHandler(e JSONEndpoint, response interface{}) gin.HandlerFunc {
	return s.Handler(JSONToEndpoint(e), response, JSONDecoder, JSONEncoder)
}

// FileHandler
func (s *Transport) FileHandler(e FileEndpoint, response interface{}) gin.HandlerFunc {
	return s.Handler(FileToEndpoint(e), response, FileDecoder, JSONEncoder)
}
