package middleware

import (
	"errors"
	"net/http"
	"time"
)

var ErrRequestTimeout = errors.New("request timeout")

func TimeoutHandler(duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if duration > 0 {
			return http.TimeoutHandler(next, duration, ErrRequestTimeout.Error())
		}
		return next
	}
}
