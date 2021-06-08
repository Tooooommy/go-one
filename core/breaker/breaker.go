package breaker

import (
	"github.com/sony/gobreaker"
)

type Option func(cfg *Conf)

func NewBreaker(options ...Option) *gobreaker.CircuitBreaker {
	cfg := DefaultConf()
	for _, opt := range options {
		opt(cfg)
	}
	return gobreaker.NewCircuitBreaker(cfg.GetGoBreakerSettings())
}

func NewTwoStepBreaker(cfg *Conf) *gobreaker.TwoStepCircuitBreaker {
	return gobreaker.NewTwoStepCircuitBreaker(cfg.GetGoBreakerSettings())
}
