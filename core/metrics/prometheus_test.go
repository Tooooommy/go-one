package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"testing"
	"time"
)

func TestProvider(t *testing.T) {
	m := NewMetrics("testnamespace", "testsubsystem")
	counter := m.NewCounter("testcounter")
	go func() {
		for {
			counter.Add(1)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe("localhost:8888", nil); err != nil {
		t.Error(err)
	}
}
