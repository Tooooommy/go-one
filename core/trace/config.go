package trace

import (
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"time"
)

type Config struct {
	Name                   string  `json:"name"`
	User                   string  `json:"user"`
	Password               string  `json:"password"`
	Disabled               bool    `json:"disabled"`
	RPCMetrics             bool    `json:"rpc_metrics"`
	Gen128Bit              bool    `json:"gen_128_bit"`
	SamplerType            string  `json:"sampler_type"`
	SamplerParam           float64 `json:"sampler_param"`
	SamplerEndpoint        string  `json:"sampler_endpoint"`
	SamplerRefreshInterval int     `json:"sampler_refresh_interval"`
	SamplerMaxOperations   int     `json:"sampler_max_operations"`
	ReporterMaxQueueSize   int     `json:"reporter_max_queue_size"`
	ReporterFlushInterval  int     `json:"reporter_flush_interval"`
	ReporterLogSpans       bool    `json:"reporter_log_spans"`
	ReporterHostPort       string  `json:"reporter_host_port"`
	ReporterReconDisabled  bool    `json:"reporter_recon_disabled"`
	ReporterReconInterval  int     `json:"reporter_recon_interval"`
	ReporterEndpoint       string  `json:"reporter_endpoint"`
}

func (cfg Config) JaegerConfig() *jaegercfg.Configuration {
	c := &jaegercfg.Configuration{
		ServiceName: cfg.Name,
		Disabled:    cfg.Disabled,
		RPCMetrics:  cfg.RPCMetrics,
		Gen128Bit:   cfg.Gen128Bit,
		Sampler: &jaegercfg.SamplerConfig{
			Type:                    cfg.SamplerType,
			Param:                   cfg.SamplerParam,
			SamplingServerURL:       cfg.SamplerEndpoint,
			SamplingRefreshInterval: time.Duration(cfg.SamplerRefreshInterval),
			MaxOperations:           cfg.SamplerMaxOperations,
		},
		Reporter: &jaegercfg.ReporterConfig{
			QueueSize:                  cfg.ReporterMaxQueueSize,
			BufferFlushInterval:        time.Duration(cfg.ReporterFlushInterval),
			LogSpans:                   cfg.ReporterLogSpans,
			LocalAgentHostPort:         cfg.ReporterHostPort,
			DisableAttemptReconnecting: cfg.ReporterReconDisabled,
			AttemptReconnectInterval:   time.Duration(cfg.ReporterReconInterval),
			CollectorEndpoint:          cfg.ReporterEndpoint,
			User:                       cfg.User,
			Password:                   cfg.Password,
		},
	}
	return c
}
