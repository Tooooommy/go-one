package trace

import (
	"github.com/go-kit/kit/endpoint"
	ktrace "github.com/go-kit/kit/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
)

// TraceServer
func TraceServer(name string) endpoint.Middleware {
	return ktrace.TraceServer(opentracing.GlobalTracer(), name)
}

// TracClient
func TracClient(name string) endpoint.Middleware {
	return ktrace.TraceClient(opentracing.GlobalTracer(), name)
}
