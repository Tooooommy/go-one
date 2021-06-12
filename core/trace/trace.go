package trace

import (
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// InitTracer
func InitTracer(cfg *Conf, options ...jaegercfg.Option) (io.Closer, error) {
	c := &jaegercfg.Configuration{
		ServiceName: cfg.Name,
		Disabled:    cfg.Disabled,
		RPCMetrics:  true,
		Gen128Bit:   true,
		Sampler: &jaegercfg.SamplerConfig{
			Type:                    cfg.Sampler.Type,
			Param:                   cfg.Sampler.Param,
			SamplingServerURL:       cfg.Sampler.Endpoint,
			SamplingRefreshInterval: time.Duration(cfg.Sampler.RefreshInterval),
			MaxOperations:           cfg.Sampler.MaxOperations,
		},
		Reporter: &jaegercfg.ReporterConfig{
			QueueSize:                  cfg.Reporter.MaxQueueSize,
			BufferFlushInterval:        time.Duration(cfg.Reporter.FlushInterval),
			LogSpans:                   cfg.Reporter.LogSpans,
			LocalAgentHostPort:         cfg.Reporter.Address,
			DisableAttemptReconnecting: cfg.Reporter.ReconDisabled,
			AttemptReconnectInterval:   time.Duration(cfg.Reporter.ReconInterval),
			CollectorEndpoint:          cfg.Reporter.Endpoint,
			User:                       cfg.Reporter.Username,
			Password:                   cfg.Reporter.Password,
		},
	}
	return c.InitGlobalTracer(c.ServiceName, options...)
}
