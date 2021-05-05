package trace

import (
	"github.com/go-kit/kit/endpoint"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
)

func Tracing(name string) endpoint.Middleware {
	return kitopentracing.TraceServer(
		opentracing.GlobalTracer(),
		name,
	)
}
