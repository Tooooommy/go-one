package grpcx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctranspot "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type (
	Service interface {
		Serve(context.Context, interface{}) (interface{}, error)
	}

	service struct {
		e         endpoint.Endpoint
		decode    grpctranspot.DecodeRequestFunc
		encode    grpctranspot.EncodeResponseFunc
		before    []grpctranspot.ServerRequestFunc
		after     []grpctranspot.ServerResponseFunc
		finalizer []grpctranspot.ServerFinalizerFunc
		handleErr HandleErrFunc
	}

	HandleErrFunc = func(context.Context, error)
	ServiceOption func(r *service)
)

// NewService
func NewService(e endpoint.Endpoint, options ...ServiceOption) Service {
	s := &service{
		e: e,
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// ServiceBefore
func ServiceBefore(before ...grpctranspot.ServerRequestFunc) ServiceOption {
	return func(r *service) {
		r.before = append(r.before, before...)
	}
}

// ServiceAfter
func ServiceAfter(after ...grpctranspot.ServerResponseFunc) ServiceOption {
	return func(r *service) {
		r.after = append(r.after, after...)
	}
}

// ServiceDecode
func ServiceDecode(dec grpctranspot.DecodeRequestFunc) ServiceOption {
	return func(r *service) {
		r.decode = dec
	}
}

// ServiceEncode
func ServiceEncode(enc grpctranspot.EncodeResponseFunc) ServiceOption {
	return func(r *service) {
		r.encode = enc
	}
}

func ServiceHandleErr(errHandler HandleErrFunc) ServiceOption {
	return func(r *service) {
		r.handleErr = errHandler
	}
}

// Serve
func (s *service) Serve(ctx context.Context, request interface{}) (response interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}

	if len(s.finalizer) > 0 {
		defer func() {
			for _, f := range s.finalizer {
				f(ctx, err)
			}
		}()
	}

	for _, f := range s.before {
		ctx = f(ctx, md)
	}

	if s.decode != nil {
		request, err = s.decode(ctx, request)
		if err != nil {
			s.handleErr(ctx, err)
			return nil, err
		}
	}

	response, err = s.e(ctx, request)
	if err != nil {
		s.handleErr(ctx, err)
		return nil, err
	}

	var mdHeader, mdTrailer metadata.MD
	for _, f := range s.after {
		ctx = f(ctx, &mdHeader, &mdTrailer)
	}

	if s.encode != nil {
		response, err = s.encode(ctx, response)
		if err != nil {
			s.handleErr(ctx, err)
			return nil, err
		}
	}

	if len(mdHeader) > 0 {
		if err = grpc.SendHeader(ctx, mdHeader); err != nil {
			s.handleErr(ctx, err)
			return nil, err
		}
	}

	if len(mdTrailer) > 0 {
		if err = grpc.SetTrailer(ctx, mdTrailer); err != nil {
			s.handleErr(ctx, err)
			return nil, err
		}
	}

	return response, nil
}
