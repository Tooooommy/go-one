package breaker

import (
	"github.com/Tooooommy/go-one/server/conf"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
)

func wrapHystrxConfig(cfg conf.BreakerConfig) hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                cfg.Timeout,
		MaxConcurrentRequests:  cfg.MaxRequests,
		SleepWindow:            cfg.Interval,
		ErrorPercentThreshold:  cfg.ErrPerThreshold,
		RequestVolumeThreshold: cfg.ReqVolThreshold,
	}
}
func HystrixBreaker(cfg conf.BreakerConfig) endpoint.Middleware {
	hystrix.ConfigureCommand(cfg.Name, wrapHystrxConfig(cfg))
	return circuitbreaker.Hystrix(cfg.Name)
}
