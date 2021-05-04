package ratelimit

import (
	"context"
	"errors"
	"github.com/Tooooommy/go-one/core/syncx"
	ep "github.com/Tooooommy/go-one/server/kitx/endpoint"
	"github.com/go-kit/kit/endpoint"
)

var ErrReachMaxLimit = errors.New("max limit is reach")

func MaxLimiter(n int) endpoint.Middleware {
	if n > 0 {
		limiter := syncx.NewLimit(n)
		return func(next endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (response interface{}, err error) {
				if limiter.TryBorrow() {
					response, err := next(ctx, request)
					if err != nil {
						return nil, err
					}
					err = limiter.Return()
					return response, err
				} else {
					return nil, ErrReachMaxLimit
				}
			}
		}
	} else {
		return ep.NopMiddleware()
	}
}
