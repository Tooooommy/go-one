package tools

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"strings"
)

const (
	TraceContextHeaderName   = "Uber-Trace-Id"
	TraceBaggageHeaderPrefix = "Uberctx-"
)

var (
	ErrNotExistTraceSpan = errors.New("trace span is not exist")
)

func ExtractTraceKeyFromCtx(ctx context.Context) (string, string, error) {
	span := opentracing.SpanFromContext(ctx)
	return ExtractTraceKeyFromSpan(span)
}

func ExtractTraceKeyFromSpan(span opentracing.Span) (traceKey string, traceName string, err error) {
	if span == nil {
		err = ErrNotExistTraceSpan
		return
	}
	carrier := opentracing.HTTPHeadersCarrier{}
	err = opentracing.GlobalTracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
	if err != nil {
		return
	}
	err = carrier.ForeachKey(func(key, val string) error {
		if key == TraceContextHeaderName {
			traceKey = val
		} else if strings.HasPrefix(key, TraceBaggageHeaderPrefix) {
			traceName = val
		}
		return nil
	})
	return
}
