package trace

import (
	"github.com/Tooooommy/go-one/core/metrics"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
)

// InitJaegerTracer
func InitJaegerTracer(cfg *Config, metrics *metrics.Metrics) (io.Closer, error) {
	c := cfg.JaegerConfig()
	var options []jaegercfg.Option
	options = append(options, jaegercfg.Logger(&logger{}))
	if metrics != nil {
		options = append(options, jaegercfg.Metrics(NewFactory(c.ServiceName, metrics)))
	}
	return c.InitGlobalTracer(c.ServiceName, options...)
}
