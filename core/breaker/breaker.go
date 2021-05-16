package breaker

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/sony/gobreaker"
	handyBreaker "github.com/streadway/handy/breaker"
	"time"
)

type (
	Breaker struct {
		cfg Config
	}
)

func NewBreaker(cfg Config) *Breaker {
	return &Breaker{cfg: cfg}
}

func (b *Breaker) GoBreaker() *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(b.cfg.GetGoBreakerSettings())
}

func (b *Breaker) HandyBreaker() handyBreaker.Breaker {
	breaker := handyBreaker.NewBreaker(float64(b.cfg.ErrPerThreshold) / 100)
	if b.cfg.Timeout > 0 {
		breaker.WithCooldown(time.Duration(b.cfg.Timeout) * time.Millisecond)
	}
	if b.cfg.Interval > 0 {
		breaker.WithWindow(time.Duration(b.cfg.Interval) * time.Millisecond)
	}
	if b.cfg.ReqVolThreshold > 0 {
		breaker.WithMinObservation(uint(b.cfg.ReqVolThreshold))
	}
	return breaker
}

func (b *Breaker) HystrixBreaker() *hystrix.CircuitBreaker {
	hystrix.ConfigureCommand(b.cfg.Name, b.cfg.GetHystrixConfig())
	breaker, _, _ := hystrix.GetCircuit(b.cfg.Name)
	return breaker
}
