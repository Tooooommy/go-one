package discov

import (
	"context"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
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
		encode    grpctranspot.EncodeRequestFunc
		decode    grpctranspot.DecodeResponseFunc
		before    []grpctranspot.ClientRequestFunc
		after     []grpctranspot.ClientResponseFunc
		finalizer []grpctranspot.ClientFinalizerFunc
	}

	InvokeFunc    func(*grpc.ClientConn, interface{}, ...grpc.CallOption) (interface{}, error)
	InvokerOption func(*invoker)
)

var (
	defaultEncode = func(ctx context.Context, request interface{}) (interface{}, error) {
		return request, nil
	}
	defaultDecode = func(ctx context.Context, response interface{}) (interface{}, error) {
		return response, nil
	}
)

// NewInvoker
func NewInvoker(options ...InvokerOption) Invoker {
	invoker := &invoker{
		encode: defaultDecode,
		decode: defaultEncode,
	}
	for _, opt := range options {
		opt(invoker)
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

		request, err = i.encode(ctx, request)
		if err != nil {
			return nil, err
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

		response, err = i.decode(ctx, response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}

// Invoke
func (i *invoker) Invoke(ctx context.Context, instancer sd.Instancer, retries int,
	timeout time.Duration, request interface{}) (interface{}, error) {
	factory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		// TODO: 暂未添加
		conn, err := grpc.DialContext(ctx, instance, grpc.WithInsecure())
		if err != nil {
			return nil, nil, err
		}
		return i.MakeEndpoint(conn), conn, nil
	}

	endpointer := sd.NewEndpointer(instancer, factory, zapx.KitL())
	e := lb.Retry(retries, timeout, lb.NewRoundRobin(endpointer))
	return e(ctx, request)
}

func (c *Client) NewInstancer(prefix string) (sd.Instancer, error) {
	if c.cfg.HaveEtcd() {
		cli, err := c.getClient()
		if err != nil {
			return nil, err
		}
		return etcdv3.NewInstancer(cli, prefix, zapx.KitL())
	} else {
		return sd.FixedInstancer([]string{prefix}), nil
	}
}
