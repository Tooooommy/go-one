package breaker

import (
	"github.com/Tooooommy/go-one/server/config"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
)

func wrapHystrxConfig(cfg config.BreakerConfig) hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                cfg.Timeout,
		MaxConcurrentRequests:  cfg.MaxRequests,
		SleepWindow:            cfg.Interval,
		ErrorPercentThreshold:  cfg.ErrPerThreshold,
		RequestVolumeThreshold: cfg.ReqVolThreshold,
	}
}
func HystrixBreaker(cfg config.BreakerConfig) endpoint.Middleware {
	hystrix.ConfigureCommand(cfg.Name, wrapHystrxConfig(cfg))
	return circuitbreaker.Hystrix(cfg.Name)
}
