package middleware

import (
	"errors"
	"log"
	"net/http"
)

var ErrContentTooLarge = errors.New("content too large")

func MaxContent(n int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		if n <= 0 {
			return next
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > n {
				// TODO:
				log.Print("")
				http.Error(w, ErrContentTooLarge.Error(), http.StatusRequestEntityTooLarge)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
