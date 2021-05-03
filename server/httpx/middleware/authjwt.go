package middleware

import (
	"github.com/Tooooommy/go-one/server/httpx"
	"net/http"
)

// 可以校验时间、校验ip、校验设备
func AuthJwt(secret, preSecret string, opts ...httpx.ParseOption) func(http.Handler) http.Handler {
	parser := httpx.NewTokenParser(httpx.TokenConfig{
		Secret:    secret,
		PreSecret: preSecret,
	}, opts...)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := parser.ParseToken(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}

			next.ServeHTTP(w, r)
		})
	}
}
