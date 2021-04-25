package middleware

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"
)

var ErrServerPanic = errors.New("server panic")

func RecoverHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if result := recover(); result != nil {
				// TODO:
				log.Printf("%v\n%s", result, debug.Stack())
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, ErrServerPanic.Error(), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
