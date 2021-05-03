package ginx

import (
	"net/http"
)

type Router interface {
	With(...http.Handler) Router
	Handle(string, string, http.Handler)
	HandleFunc(string, string, http.HandlerFunc)
	SetNoFound(http.Handler)
	SetNoMethod(http.Handler)
}
