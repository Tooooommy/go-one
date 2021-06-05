package metrics

import (
	"github.com/go-kit/kit/metrics"
	kitprovider "github.com/go-kit/kit/metrics/provider"
)

// https://github.com/zserge/metric
// https://github.com/propan/expvardash
type (
	// Metrics
	Metrics struct {
		provider  kitprovider.Provider
		namespace string
		subsystem string
	}
)

// NewMetrics
func NewMetrics(namespace, subsystem string) kitprovider.Provider {
	var provider kitprovider.Provider
	if len(namespace) != 0 && len(subsystem) != 0 {
		provider = kitprovider.NewPrometheusProvider(namespace, subsystem)
	} else {
		provider = kitprovider.NewExpvarProvider()
	}
	return &Metrics{
		provider:  provider,
		namespace: namespace,
		subsystem: subsystem,
	}
}

// NewCounter
func (m *Metrics) NewCounter(name string) metrics.Counter {
	return m.provider.NewCounter(name)
}

// NewGauge
func (m *Metrics) NewGauge(name string) metrics.Gauge {
	return m.provider.NewGauge(name)
}

// NewHistogram
func (m *Metrics) NewHistogram(name string, buckets int) metrics.Histogram {
	return m.provider.NewHistogram(name, buckets)
}

// Stop
func (m *Metrics) Stop() {
	m.provider.Stop()
}
