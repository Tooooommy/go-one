package jaegerx

import (
	"github.com/Tooooommy/go-one/core/metrics"
	kmetrics "github.com/go-kit/kit/metrics"
	xmetrics "github.com/uber/jaeger-lib/metrics"
	xkit "github.com/uber/jaeger-lib/metrics/go-kit"
)

func NewFactory(namespace string, metrics *metrics.Metrics) xmetrics.Factory {
	return xkit.Wrap(namespace, &factory{
		metrics: metrics,
	})
}

type factory struct {
	metrics *metrics.Metrics
	buckets int
}

func (f *factory) Counter(name string) kmetrics.Counter {
	return f.metrics.NewCounter(name)
}

func (f *factory) Gauge(name string) kmetrics.Gauge {
	return f.metrics.NewGauge(name)
}

func (f *factory) Histogram(name string) kmetrics.Histogram {
	return f.metrics.NewHistogram(name, 50)
}

func (f *factory) Capabilities() xkit.Capabilities {
	return xkit.Capabilities{Tagging: false}
}
