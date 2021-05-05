package trace

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"testing"
	"time"
)

func TestNewJaegerTracer(t *testing.T) {
	closer, err := InitJaegerTracer(Config{
		Name:             "go-one",
		SamplerType:      "const",
		SamplerParam:     1,
		ReporterHostPort: "127.0.0.1:6831",
		ReporterLogSpans: true,
	})
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	span1 := opentracing.GlobalTracer().StartSpan("span1")
	defer span1.Finish()
	span1.SetBaggageItem("span1", "span1")
	time.Sleep(10 * time.Millisecond)

	ctx := opentracing.ContextWithSpan(context.Background(), span1)
	span2, ctx := opentracing.StartSpanFromContext(ctx, "span2")
	defer span2.Finish()
	span2.LogKV("span2", "span2")
	time.Sleep(10 * time.Millisecond)

	ctx = opentracing.ContextWithSpan(ctx, span2)
	span3, ctx := opentracing.StartSpanFromContext(ctx, "span3")
	defer span3.Finish()
	span3.SetTag("span3", "span3")
	time.Sleep(10 * time.Millisecond)

}
