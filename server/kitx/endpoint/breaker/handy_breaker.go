package breaker

import (
	"github.com/Tooooommy/go-one/server/config"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	handyBreaker "github.com/streadway/handy/breaker"
	"time"
)

func HandyBreaker(cfg config.BreakerConfig) endpoint.Middleware {
	cb := handyBreaker.NewBreaker(float64(cfg.ErrPerThreshold) / 100)
	cb.WithCooldown(time.Duration(cfg.Timeout) * time.Millisecond)
	cb.WithWindow(time.Duration(cfg.Interval) * time.Millisecond)
	cb.WithMinObservation(uint(cfg.ReqVolThreshold))
	return circuitbreaker.HandyBreaker(cb)
}
