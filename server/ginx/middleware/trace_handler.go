package middleware

import (
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/gin-gonic/gin"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
)

// TraceHandler: 开启链路追踪
func TraceHandler(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		span := opentracing.SpanFromContext(ctx)
		if span == nil {
			span = opentracing.GlobalTracer().StartSpan(name)
			defer span.Finish()
		}
		ctx = opentracing.ContextWithSpan(ctx, span)
		request := kitopentracing.ContextToHTTP(opentracing.GlobalTracer(), zapx.KitL())
		c.Request = c.Request.WithContext(request(ctx, c.Request))
		c.Next()
	}
}
