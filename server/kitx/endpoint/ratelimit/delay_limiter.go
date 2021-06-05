package ratelimit

import (
	ep "github.com/Tooooommy/go-one/server/kitx/endpoint"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
)

func DelayLimiter(limit float64) endpoint.Middleware {
	if limit > 0 {
		limiter := rate.NewLimiter(rate.Limit(limit), 1)
		return ratelimit.NewDelayingLimiter(limiter)
	} else {
		return ep.NopMiddleware()
	}
}
