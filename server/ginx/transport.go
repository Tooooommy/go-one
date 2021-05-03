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

// Handler
func Handler(e endpoint.Endpoint, request interface{}, dec DecodeFunc, enc EncodeFunc, options ...httptransport.ServerOption) gin.HandlerFunc {
	return func(c *gin.Context) {
		httptransport.NewServer(
			e,
			dec(c, request),
			enc(c),
			options...,
		).ServeHTTP(c.Writer, c.Request)
	}
}

// NoHandler
func NoHandler(e endpoint.Endpoint, options ...httptransport.ServerOption) gin.HandlerFunc {
	return Handler(e, nil, NoDecoder, NoEncoder, options...)
}

// JSONHandler
func JSONHandler(e JSONEndpoint, response interface{}, options ...httptransport.ServerOption) gin.HandlerFunc {
	return Handler(JSONToEndpoint(e), response, JSONDecoder, JSONEncoder, options...)
}

// FileHandler
func FileHandler(e endpoint.Endpoint, response interface{}, options ...httptransport.ServerOption) gin.HandlerFunc {
	return Handler(e, response, FileDecoder, JSONEncoder, options...)
}
