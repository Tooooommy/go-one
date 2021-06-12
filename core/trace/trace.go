package trace

import (
	"github.com/Tooooommy/go-one/core/metrics"
	"github.com/Tooooommy/go-one/core/trace/hooks"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
)

// InitTracer
func InitTracer(cfg *Conf, metrics *metrics.Metrics) (io.Closer, error) {
	c := cfg.JaegerConf()
	var options []jaegercfg.Option

	options = append(options,
		jaegercfg.Logger(hooks.NewLogger()),
		jaegercfg.Metrics(hooks.NewFactory(c.ServiceName, metrics)),
		jaegercfg.Gen128Bit(true),
		jaegercfg.PoolSpans(true),
	)
	return c.InitGlobalTracer(c.ServiceName, options...)
}
