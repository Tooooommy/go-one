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
		middlewares []endpoint.Middleware
		encode      EncodeFunc
		decode      DecodeFunc
	}
)

// NewTransport
func NewTransport() *Transport {
	return &Transport{}
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
			// options
		).ServeHTTP(c.Writer, c.Request)
	}
}
