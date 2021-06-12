package hooks

import (
	"github.com/Tooooommy/go-one/core/metrics"
	kmetrics "github.com/go-kit/kit/metrics"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	xkit "github.com/uber/jaeger-lib/metrics/go-kit"
)

func NewMetrics(metrics metrics.Metrics) jaegercfg.Option {
	return jaegercfg.Metrics(xkit.Wrap(metrics.Namespace(), &factory{metrics: metrics}))
}

type factory struct {
	metrics metrics.Metrics
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
