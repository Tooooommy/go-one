package httpx

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type (
	DecodeRequestFunc   func(*gin.Context, interface{}) error
	EncodeResponseFunc  func(*gin.Context, interface{}) error
	ServerRequestFunc   func(*gin.Context)
	ServerResponseFunc  func(*gin.Context, interface{})
	ServerFinalizerFunc func(*gin.Context)

	Service interface {
		Serve(request interface{}) gin.HandlerFunc
	}

	HandleErr func(*gin.Context, error)

	service struct {
		e         endpoint.Endpoint
		decode    DecodeRequestFunc
		encode    EncodeResponseFunc
		before    []ServerRequestFunc
		after     []ServerResponseFunc
		finalizer []ServerFinalizerFunc
		handleErr HandleErr
	}

	ServiceOption func(*service)
)

// NewService
func NewService(e endpoint.Endpoint, options ...ServiceOption) Service {
	s := &service{
		e: e,
		handleErr: func(c *gin.Context, err error) {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		},
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// ServiceDecode
func ServiceDecode(decode ...DecodeRequestFunc) ServiceOption {
	return func(s *service) {
		s.decode = ChainDecoder(s.decode, decode...)
	}
}

// ServiceEncode
func ServiceEncode(encode EncodeResponseFunc) ServiceOption {
	return func(s *service) {
		s.encode = encode
	}
}

// ServiceBefore
func ServiceBefore(before ...ServerRequestFunc) ServiceOption {
	return func(s *service) {
		s.before = append(s.before, before...)
	}
}

// ServiceAfter
func ServiceAfter(after ...ServerResponseFunc) ServiceOption {
	return func(s *service) {
		s.after = append(s.after, after...)
	}
}

// ServiceFinalizer
func ServiceFinalizer(finalizer ...ServerFinalizerFunc) ServiceOption {
	return func(s *service) {
		s.finalizer = append(s.finalizer, finalizer...)
	}
}

// ServiceHandleErr
func ServiceHandleErr(handleErr HandleErr) ServiceOption {
	return func(s *service) {
		s.handleErr = handleErr
	}
}

// Serve
func (s *service) Serve(request interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(s.finalizer) > 0 {
			defer func() {
				for _, f := range s.finalizer {
					f(c)
				}
			}()
		}

		for _, f := range s.before {
			f(c)
		}

		if s.decode == nil {
			err := s.decode(c, request)
			if err != nil {
				s.handleErr(c, err)
				return
			}
		}

		response, err := s.e(c, request)
		if err != nil {
			s.handleErr(c, err)
			return
		}

		for _, f := range s.after {
			f(c, response)
		}

		if s.encode == nil {
			err = s.encode(c, response)
			if err != nil {
				s.handleErr(c, err)
				return
			}
		}
	}
}
