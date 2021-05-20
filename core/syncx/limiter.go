package syncx

import (
	"errors"
)

// Fork from go-zero

// ErrLimitReturn indicates that the more than borrowed elements were returned.
var ErrLimitReturn = errors.New("discarding limited token, resource pool is full, someone returned multiple times")
var ErrLimitAllow = errors.New("limiter resource is full, cannot allow")

type Limiter struct {
	pool chan struct{}
}

func NewLimiter(n int) Limiter {
	return Limiter{
		pool: make(chan struct{}, n),
	}
}

func (l Limiter) Allow() error {
	select {
	case l.pool <- struct{}{}:
		return nil
	default:
		return ErrLimitAllow
	}
}

func (l Limiter) Wait() {
	l.pool <- struct{}{}
}

func (l Limiter) Return() error {
	select {
	case <-l.pool:
		return nil
	default:
		return ErrLimitReturn
	}
}

func (l Limiter) ReportResult(result error) {
	_ = l.Return()
}
