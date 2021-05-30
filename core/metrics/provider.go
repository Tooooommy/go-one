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
		provider kitprovider.Provider
	}
)

// NewMetrics
func NewMetrics(namespace, subsystem string) *Metrics {
	provider := kitprovider.NewExpvarProvider()
	if len(namespace) != 0 && len(subsystem) != 0 {
		provider = kitprovider.NewPrometheusProvider(namespace, subsystem)
	}
	return &Metrics{
		provider: provider,
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
func (m *Metrics) NewHistogram(name string) metrics.Histogram {
	return m.provider.NewHistogram(name, 50)
}

// Stop
func (m *Metrics) Stop() {
	m.provider.Stop()
}
