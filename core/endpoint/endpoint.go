package endpoint

import "github.com/go-kit/kit/endpoint"

// NopMiddleware
func NopMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return next
	}
}
