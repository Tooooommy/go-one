package middleware

import (
	"errors"
	"golang.org/x/time/rate"
	"net/http"
)

var ErrRequestRateLimit = errors.New("request rate limit")

// 限制
func RatelimitHandler(n int) func(http.Handler) http.Handler {
	if n <= 0 {
		return func(next http.Handler) http.Handler {
			return next
		}
	} else {
		limiter := rate.NewLimiter(1, n)
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if limiter.Allow() {
					next.ServeHTTP(w, r)
				} else {
					http.Error(w, ErrRequestRateLimit.Error(), http.StatusServiceUnavailable)
				}
			})
		}
	}
}
