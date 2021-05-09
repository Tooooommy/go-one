package metrics

import (
	kitprom "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// 普罗米修斯
func NewPromxCounter(cfg Config, labelNames ...string) *kitprom.Counter {
	return kitprom.NewCounterFrom(prometheus.CounterOpts{
		Namespace:   cfg.Namespace,
		Subsystem:   cfg.Subsystem,
		Name:        cfg.Name,
		Help:        cfg.Help,
		ConstLabels: cfg.ConstLabels,
	}, labelNames)
}

func NewPromxGauge(cfg Config, labelNames ...string) *kitprom.Gauge {
	return kitprom.NewGaugeFrom(prometheus.GaugeOpts{
		Namespace:   cfg.Namespace,
		Subsystem:   cfg.Subsystem,
		Name:        cfg.Name,
		Help:        cfg.Help,
		ConstLabels: cfg.ConstLabels,
	}, labelNames)
}

func NewPromxHistogram(cfg Config, labelNames ...string) *kitprom.Histogram {
	return kitprom.NewHistogramFrom(prometheus.HistogramOpts{
		Namespace:   cfg.Namespace,
		Subsystem:   cfg.Subsystem,
		Name:        cfg.Name,
		Help:        cfg.Help,
		ConstLabels: cfg.ConstLabels,
		Buckets:     cfg.Buckets,
	}, labelNames)
}

func NewPromxSummary(cfg Config, labelNames ...string) *kitprom.Summary {
	return kitprom.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace:   cfg.Namespace,
		Subsystem:   cfg.Subsystem,
		Name:        cfg.Name,
		Help:        cfg.Help,
		ConstLabels: cfg.ConstLabels,
		Objectives:  cfg.Objectives,
		MaxAge:      time.Duration(cfg.MaxAge),
		AgeBuckets:  cfg.AgeBuckets,
		BufCap:      cfg.BufCap,
	}, labelNames)
}
