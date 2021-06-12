package metrics

import (
	kmetrics "github.com/go-kit/kit/metrics"
	kprovider "github.com/go-kit/kit/metrics/provider"
)

// https://github.com/zserge/metric
// https://github.com/propan/expvardash
type (
	// metrics
	metrics struct {
		namespace string
		subsystem string
		provider  kprovider.Provider
	}
)

// NewMetrics
func NewMetrics(namespace, subsystem string) kprovider.Provider {
	provider := kprovider.NewPrometheusProvider(namespace, subsystem)
	return &metrics{
		provider:  provider,
		namespace: namespace,
		subsystem: subsystem,
	}
}

// NewCounter
func (m *metrics) NewCounter(name string) kmetrics.Counter {
	return m.provider.NewCounter(name)
}

// NewGauge
func (m *metrics) NewGauge(name string) kmetrics.Gauge {
	return m.provider.NewGauge(name)
}

// NewHistogram
func (m *metrics) NewHistogram(name string, buckets int) kmetrics.Histogram {
	return m.provider.NewHistogram(name, buckets)
}

// Stop
func (m *metrics) Stop() {
	m.provider.Stop()
}

// Namespace
func (m *metrics) Namespace() string {
	return m.namespace
}

// Subsystem
func (m *metrics) Subsystem() string {
	return m.subsystem
}
