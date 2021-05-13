package jaegerx

import (
	"github.com/Tooooommy/go-one/core/trace"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
)

// InitJaegerTracer
func InitJaegerTracer(cfg trace.Config, options ...jaegercfg.Option) (io.Closer, error) {
	c := cfg.JaegerConfig()
	options = append(options, jaegercfg.Logger(jaeger.StdLogger))
	return c.InitGlobalTracer(c.ServiceName, options...)
}
