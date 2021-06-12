package metrics

import (
	kmetrics "github.com/go-kit/kit/metrics"
	kprovider "github.com/go-kit/kit/metrics/provider"
)

type expvar struct {
	provider kprovider.Provider
}

func NewExpvar() Metrics {
	provider := kprovider.NewExpvarProvider()
	return &expvar{provider: provider}
}

func (e *expvar) NewCounter(name string) kmetrics.Counter {
	return e.provider.NewCounter(name)
}

func (e *expvar) NewGauge(name string) kmetrics.Gauge {
	return e.provider.NewGauge(name)
}

func (e *expvar) NewHistogram(name string, buckets int) kmetrics.Histogram {
	return e.provider.NewHistogram(name, buckets)
}

func (e *expvar) Stop() {
	e.provider.Stop()
}

func (e *expvar) Namespace() string {
	return ""
}

func (e *expvar) Subsystem() string {
	return ""
}
