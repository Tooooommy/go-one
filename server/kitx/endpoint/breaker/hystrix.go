package breaker

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
)

func wrapHystrxConfig(cfg Config) hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                cfg.Timeout,
		MaxConcurrentRequests:  cfg.MaxRequests,
		SleepWindow:            cfg.Interval,
		ErrorPercentThreshold:  cfg.ErrPerThreshold,
		RequestVolumeThreshold: cfg.ReqVolThreshold,
	}
}
func HystrixBreaker(cfg Config) endpoint.Middleware {
	hystrix.ConfigureCommand(cfg.Name, wrapHystrxConfig(cfg))
	return circuitbreaker.Hystrix(cfg.Name)
}
