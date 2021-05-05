package trace

import (
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
)

type JaegerOption func(configuration *jaegercfg.Configuration)

// InitJaegerTracer
func InitJaegerTracer(cfg Config, options ...JaegerOption) (io.Closer, error) {
	c := cfg.JaegerConfig()
	for _, opt := range options {
		opt(c)
	}
	return c.InitGlobalTracer(c.ServiceName,
		jaegercfg.Logger(jaeger.StdLogger))
}
