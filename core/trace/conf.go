package trace

import (
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"time"
)

type Conf struct {
	Name     string   `json:"name"`
	Disabled bool     `json:"disabled"`
	Sampler  Sampler  `json:"sampler"`
	Reporter Reporter `json:"reporter"`
}

type Sampler struct {
	Type            string  `json:"type"`
	Param           float64 `json:"param"`
	Endpoint        string  `json:"endpoint"`
	RefreshInterval int     `json:"refresh_interval"`
	MaxOperations   int     `json:"max_operations"`
}

type Reporter struct {
	MaxQueueSize  int    `json:"max_queue_size"`
	FlushInterval int    `json:"flush_interval"`
	LogSpans      bool   `json:"log_spans"`
	Address       string `json:"address"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	ReconDisabled bool   `json:"recon_disabled"`
	ReconInterval int    `json:"recon_interval"`
	Endpoint      string `json:"endpoint"`
}

func (cfg *Conf) JaegerConf() *jaegercfg.Configuration {
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
	return c
}
