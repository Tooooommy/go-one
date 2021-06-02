package discov

import (
	"context"
	"github.com/Tooooommy/go-one/core/zapx"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	grpctranspot "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"io"
	"time"
)

type (
	Invoker interface {
		Invoke(context.Context, sd.Instancer, int, time.Duration, interface{}) (interface{}, error)
	}

	invoker struct {
		encode  grpctranspot.EncodeRequestFunc
		decode  grpctranspot.DecodeResponseFunc
		options []grpctranspot.ClientOption
		method  string
		service string
	}
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
func NewInvoker(service, method string, options ...InvokerOption) Invoker {
	invoker := &invoker{
		service: service,
		method:  method,
		encode:  defaultDecode,
		decode:  defaultEncode,
		options: []grpctranspot.ClientOption{},
	}
	for _, opt := range options {
		opt(invoker)
	}
	return invoker
}

// SetEncode
func SetEncode(encode grpctranspot.EncodeRequestFunc) InvokerOption {
	return func(invoker *invoker) {
		invoker.encode = encode
	}
}

// SetDecode
func SetDecode(decode grpctranspot.DecodeResponseFunc) InvokerOption {
	return func(invoker *invoker) {
		invoker.decode = decode
	}
}

// SetOption
func SetOptions(options ...grpctranspot.ClientOption) InvokerOption {
	return func(i *invoker) {
		i.options = append(i.options, options...)
	}
}

// Invoke
func (i *invoker) Invoke(
	ctx context.Context,
	instancer sd.Instancer,
	retries int,
	timeout time.Duration,
	request interface{},
) (interface{}, error) {
	factory := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.DialContext(
			ctx,
			instance,
			grpc.WithInsecure(), // TODO: 暂未添加
		)
		if err != nil {
			return nil, nil, err
		}
		i.options = append(i.options, grpctranspot.ClientBefore(kitjwt.ContextToGRPC()))
		client := grpctranspot.NewClient(
			conn,
			i.service,
			i.method,
			i.encode,
			i.decode,
			request,
			i.options...,
		)
		return client.Endpoint(), conn, nil
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
