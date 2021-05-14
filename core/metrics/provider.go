package metrics

import (
	kitmetrics "github.com/go-kit/kit/metrics"
	kitprovider "github.com/go-kit/kit/metrics/provider"
)

// https://github.com/zserge/metric
// https://github.com/propan/expvardash
type (
	// Metrics
	Metrics struct {
		provider  kitprovider.Provider
		counter   kitmetrics.Counter
		gauge     kitmetrics.Gauge
		histogram kitmetrics.Histogram
	}
)

// NewMetrics
func NewMetrics(namespace, subsystem, name string) *Metrics {
	provider := kitprovider.NewExpvarProvider()
	if len(namespace) != 0 && len(subsystem) != 0 {
		provider = kitprovider.NewPrometheusProvider(namespace, subsystem)
	}
	counter := provider.NewCounter(name)
	gauge := provider.NewGauge(name)
	histogram := provider.NewHistogram(name, 50)
	return &Metrics{
		provider:  provider,
		counter:   counter,
		gauge:     gauge,
		histogram: histogram,
	}
}

// With
func (m *Metrics) With(labelValues ...string) *Metrics {
	m.counter = m.counter.With(labelValues...)
	m.gauge = m.gauge.With(labelValues...)
	m.histogram = m.histogram.With(labelValues...)
	return m
}

// Add
func (m *Metrics) Add(delta float64) *Metrics {
	m.counter.Add(delta)
	return m
}

// Metrics
func (m *Metrics) Set(value float64) *Metrics {
	m.gauge.Set(value)
	return m
}

func (m *Metrics) Incr(value float64) *Metrics {
	m.gauge.Add(value)
	return m
}

func (m *Metrics) Decr(value float64) *Metrics {
	m.gauge.Add(-value)
	return m
}

// Observe
func (m *Metrics) Observe(value float64) *Metrics {
	m.histogram.Observe(value)
	return m
}

// Counter
func (m *Metrics) Counter() kitmetrics.Counter {
	return m.counter
}

// Gauge
func (m *Metrics) Gauge() kitmetrics.Gauge {
	return m.gauge
}

// Histogram
func (m *Metrics) Histogram() kitmetrics.Histogram {
	return m.histogram
}
