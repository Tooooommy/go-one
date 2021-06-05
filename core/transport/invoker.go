package transport

import (
	"context"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	grpctranspot "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

type (
	Invoker interface {
		Invoke(context.Context, sd.Instancer, int, time.Duration, interface{}) (interface{}, error)
	}

	invoker struct {
		invoke    InvokeFunc
		factory   func(context.Context) sd.Factory
		encode    grpctranspot.EncodeRequestFunc
		decode    grpctranspot.DecodeResponseFunc
		before    []grpctranspot.ClientRequestFunc
		after     []grpctranspot.ClientResponseFunc
		finalizer []grpctranspot.ClientFinalizerFunc
	}

	InvokeFunc    func(*grpc.ClientConn, interface{}, ...grpc.CallOption) (interface{}, error)
	InvokerOption func(*invoker)
)

// NewInvoker
func NewInvoker(invoke InvokeFunc, options ...InvokerOption) Invoker {
	invoker := &invoker{
		invoke: invoke,
	}
	for _, opt := range options {
		opt(invoker)
	}

	invoker.factory = func(ctx context.Context) sd.Factory {
		return func(instance string) (endpoint.Endpoint, io.Closer, error) {
			// TODO: 暂未添加
			conn, err := grpc.DialContext(ctx, instance, grpc.WithInsecure())
			if err != nil {
				return nil, nil, err
			}
			return invoker.MakeEndpoint(conn), conn, nil
		}
	}

	return invoker
}

// InvokerFunc
func InvokerFunc(invoke InvokeFunc) InvokerOption {
	return func(i *invoker) {
		i.invoke = invoke
	}
}

// InvokerBefore
func InvokerBefore(before ...grpctranspot.ClientRequestFunc) InvokerOption {
	return func(i *invoker) {
		i.before = append(i.before, before...)
	}
}

// InvokerAfter
func InvokerAfter(after ...grpctranspot.ClientResponseFunc) InvokerOption {
	return func(i *invoker) {
		i.after = append(i.after, after...)
	}
}

// InvokerFinalizer
func InvokerFinalizer(finalizer ...grpctranspot.ClientFinalizerFunc) InvokerOption {
	return func(i *invoker) {
		i.finalizer = append(i.finalizer, finalizer...)
	}
}

// InvokerEncode
func InvokerEncode(encode grpctranspot.EncodeRequestFunc) InvokerOption {
	return func(invoker *invoker) {
		invoker.encode = encode
	}
}

// InvokerDecode
func InvokerDecode(decode grpctranspot.DecodeResponseFunc) InvokerOption {
	return func(invoker *invoker) {
		invoker.decode = decode
	}
}

// MakeEndpoint
func (i *invoker) MakeEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		if i.finalizer != nil {
			defer func() {
				for _, f := range i.finalizer {
					f(ctx, err)
				}
			}()
		}

		if i.encode != nil {
			request, err = i.encode(ctx, request)
			if err != nil {
				return nil, err
			}
		}

		md := &metadata.MD{}
		for _, f := range i.before {
			ctx = f(ctx, md)
		}
		ctx = metadata.NewOutgoingContext(ctx, *md)

		var header, trailer metadata.MD
		response, err = i.invoke(conn, request, grpc.Header(&header), grpc.Trailer(&trailer))
		if err != nil {
			return nil, err
		}

		for _, f := range i.after {
			ctx = f(ctx, header, trailer)
		}

		if i.decode != nil {
			response, err = i.decode(ctx, response)
			if err != nil {
				return nil, err
			}
		}

		return response, nil
	}
}

// Invoke
func (i *invoker) Invoke(ctx context.Context, instancer sd.Instancer, retries int, timeout time.Duration, request interface{}) (interface{}, error) {
	endpointer := sd.NewEndpointer(instancer, i.factory(ctx), zapx.KitL())
	e := lb.Retry(retries, timeout, lb.NewRoundRobin(endpointer))
	return e(ctx, request)
}
