package jaegerx

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestNewJaegerTracer(t *testing.T) {
	carrier := opentracing.HTTPHeadersCarrier(http.Header{})

	closer, err := InitJaegerTracer(&Config{
		Name: "go-one",
		Sampler: Sampler{
			Type:  "const",
			Param: 1,
		},
		Reporter: Reporter{
			Address:  "127.0.0.1:6831",
			LogSpans: true,
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	tracer := opentracing.GlobalTracer()
	span1 := tracer.StartSpan("span1")
	defer span1.Finish()
	ext.HTTPMethod.Set(span1, "get")
	ext.HTTPUrl.Set(span1, "www.test.hello.com")
	host, portString, err := net.SplitHostPort("localhost")
	if err == nil {
		ext.PeerHostname.Set(span1, host)
		if port, err := strconv.Atoi(portString); err == nil {
			ext.PeerPort.Set(span1, uint16(port))
		}
	} else {
		ext.PeerHostname.Set(span1, "localhost")
	}

	// There's nothing we can do with any errors here.
	if err = tracer.Inject(
		span1.Context(),
		opentracing.HTTPHeaders,
		carrier,
	); err != nil {
		panic(err)
	}
	span1.SetBaggageItem("span1", "span1")
	time.Sleep(10 * time.Millisecond)

	ctx := opentracing.ContextWithSpan(context.Background(), span1)
	span2, ctx := opentracing.StartSpanFromContext(ctx, "span2")
	defer span2.Finish()
	span2.LogKV("span2", "span2")
	time.Sleep(10 * time.Millisecond)

	span3, ctx := opentracing.StartSpanFromContext(ctx, "span3")
	defer span3.Finish()
	span3.SetTag("span3", "span3")
	time.Sleep(10 * time.Millisecond)

	span4 := opentracing.StartSpan("span4", opentracing.ChildOf(span3.Context()))
	defer span4.Finish()
	span5 := opentracing.StartSpan("span5", opentracing.ChildOf(span3.Context()))
	defer span5.Finish()

}
