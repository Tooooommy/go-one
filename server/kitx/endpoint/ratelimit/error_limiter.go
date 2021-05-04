package ratelimit

import (
	ep "github.com/Tooooommy/go-one/server/kitx/endpoint"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
)

func ErrorLimiter(n int) endpoint.Middleware {
	if n > 0 {
		limiter := rate.NewLimiter(1, n)
		return ratelimit.NewErroringLimiter(limiter)
	} else {
		return ep.NopMiddleware()
	}
}
