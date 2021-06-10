package breaker

import (
	"github.com/sony/gobreaker"
	"math"
	"time"
)

type (
	Option func(setting gobreaker.Settings)
)

func MaxRequests(max uint32) Option {
	return func(setting gobreaker.Settings) {
		setting.MaxRequests = max
	}
}

func Interval(interval time.Duration) Option {
	return func(setting gobreaker.Settings) {
		setting.Interval = interval
	}
}

func Timeout(timeout time.Duration) Option {
	return func(setting gobreaker.Settings) {
		setting.Timeout = timeout
	}
}

func ReadyToTrip(trip func(counts gobreaker.Counts) bool) Option {
	return func(setting gobreaker.Settings) {
		setting.ReadyToTrip = trip
	}
}

func OnStateChange(change func(name string, from gobreaker.State, to gobreaker.State)) Option {
	return func(setting gobreaker.Settings) {
		setting.OnStateChange = change
	}
}

func defaultReadyToTrip(errPercent uint32) func(counts gobreaker.Counts) bool {
	return func(counts gobreaker.Counts) bool {
		return counts.TotalFailures/counts.Requests*100 > errPercent
	}
}

func NewBreaker(name string, options ...Option) *gobreaker.CircuitBreaker {
	setting := gobreaker.Settings{
		Name:        name,
		MaxRequests: math.MaxUint32,
	}
	for _, opt := range options {
		opt(setting)
	}
	return gobreaker.NewCircuitBreaker(setting)
}

func NewNextBreaker(name string, options ...Option) *gobreaker.TwoStepCircuitBreaker {
	setting := gobreaker.Settings{
		Name:        name,
		MaxRequests: math.MaxUint32,
	}
	for _, opt := range options {
		opt(setting)
	}
	return gobreaker.NewTwoStepCircuitBreaker(setting)
}
