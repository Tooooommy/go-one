package trace

import (
	"github.com/Tooooommy/go-one/core/metrics"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
)

// InitTracer
func InitTracer(cfg *Config, metrics *metrics.Metrics) (io.Closer, error) {
	c := cfg.JaegerConfig()
	var options []jaegercfg.Option

	options = append(options,
		jaegercfg.Logger(&logger{}),
		// jaegercfg.Metrics(NewFactory(c.ServiceName, metrics)),
		jaegercfg.Gen128Bit(true),
		jaegercfg.PoolSpans(true),
	)
	return c.InitGlobalTracer(c.ServiceName, options...)
}
