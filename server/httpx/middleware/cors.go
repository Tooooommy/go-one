package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func CrossOrigins(origins ...string) http.Handler {
	options := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}
	if len(origins) > 0 {
		options.AllowedOrigins = origins
	}
	return http.HandlerFunc(cors.New(options).HandlerFunc)
}
