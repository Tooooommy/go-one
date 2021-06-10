package breaker

import (
	kitbreaker "github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/sony/gobreaker"
)

func Breaker(breaker *gobreaker.CircuitBreaker) endpoint.Middleware {
	return kitbreaker.Gobreaker(breaker)
}
