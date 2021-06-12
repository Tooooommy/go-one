package metrics

import (
	kmetrics "github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/discard"
)

type nop struct{}

func NewNop() Metrics {
	return &nop{}
}

func (n *nop) NewCounter(name string) kmetrics.Counter {
	return discard.NewCounter()
}

func (n *nop) NewGauge(name string) kmetrics.Gauge {
	return discard.NewGauge()
}

func (n *nop) NewHistogram(name string, buckets int) kmetrics.Histogram {
	return discard.NewHistogram()
}

func (n *nop) Stop() {}

func (n *nop) Namespace() string {
	return "nop_metrics_namespace"
}

func (n *nop) Subsystem() string {
	return "nop_metrics_subsystem"
}
