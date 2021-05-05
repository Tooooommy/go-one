package ratelimit

import (
	ep "github.com/Tooooommy/go-one/server/kitx/endpoint"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"time"
)

func DelayLimiter(interval, burst int) endpoint.Middleware {
	if burst > 0 {
		limiter := rate.NewLimiter(rate.Every(time.Second*time.Duration(interval)), burst)
		return ratelimit.NewDelayingLimiter(limiter)
	} else {
		return ep.NopMiddleware()
	}
}
