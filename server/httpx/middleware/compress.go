package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

const (
	contentEncoding = "Content-Encoding"
	gunzip          = "zip"
)

func GzipHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get(contentEncoding), gunzip) {
			reader, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			r.Body = reader
		}
		next.ServeHTTP(w, r)
	})
}
