package breaker

import (
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/sony/gobreaker"
	"time"
)

func wrapGoBreakerSettings(cfg Config) gobreaker.Settings {
	return gobreaker.Settings{
		Name:        cfg.Name,
		MaxRequests: uint32(cfg.MaxRequests),
		Interval:    time.Duration(cfg.Interval) * time.Millisecond,
		Timeout:     time.Duration(cfg.Timeout) * time.Millisecond,
		ReadyToTrip: readyToTrip(cfg.ErrPerThreshold),
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			// TODO: 告警
			zapx.Info().String("name", name).
				Int("from_state", int(from)).
				Int("to_state", int(to)).
				Msg("状态发生变化")
		},
	}
}

func readyToTrip(errPerThreshold int) func(counts gobreaker.Counts) bool {
	if errPerThreshold > 0 {
		return func(counts gobreaker.Counts) bool {
			total := counts.TotalFailures + counts.TotalSuccesses
			return counts.TotalFailures/total*100 > uint32(errPerThreshold)
		}
	}
	return nil
}

func GoBreaker(cfg Config) endpoint.Middleware {
	breaker := gobreaker.NewCircuitBreaker(wrapGoBreakerSettings(cfg))
	return circuitbreaker.Gobreaker(breaker)
}
