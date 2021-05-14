package ginx

import (
	"github.com/Tooooommy/go-one/core/metrics"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// Nop: 空中间件
func Nop() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// CorsHandler: 跨域请求中间件
func CorsHandler(origins ...string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowHeaders: []string{
			"Content-Type",
			"x-requested-by",
			"*",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// NoCache: 去掉浏览器缓存
func NoCacheHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")
		ctx.Next()
	}
}

// GunzipHandler: gzip压缩处理
func GunzipHandler(level int, options ...gzip.Option) gin.HandlerFunc {
	return gzip.Gzip(level, options...)
}

var TimeoutReason = "Request Timeout"

// TimeoutHandler: 请求超时处理
func TimeoutHandler(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if duration > 0 {
			http.TimeoutHandler(
				http.HandlerFunc(func(http.ResponseWriter, *http.Request) { c.Next() }),
				duration*time.Millisecond,
				TimeoutReason,
			).ServeHTTP(c.Writer, c.Request)
		}
	}
}

// RequestIDHandler: 生成唯一值
func RequestIDHandler() gin.HandlerFunc {
	return requestid.New()
}

var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

// RealIPHandler: 获取真实IP
func RealIPHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.Request.Header.Get(xRealIP)
		if ip == "" {
			if xff := c.Request.Header.Get(xForwardedFor); xff != "" {
				i := strings.Index(xff, ", ")
				if i == -1 {
					i = len(xff)
				}
				ip = xff[:i]
			}
		}
		c.Request.RemoteAddr = ip
		c.Next()
	}
}

// RecoveryHandler: panic恢复
func RecoveryHandler(c *gin.Context) {
	defer func() {
		if result := recover(); result != nil {
			zapx.Error().Any("Recovery Result: %+v", result).Msg(string(debug.Stack()))
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}

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

// MetricsHandler: 开启普罗米修斯
func MetricsHandler(m *metrics.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		n := time.Now()
		defer func() {
			labels := []string{c.Request.Method, c.Request.URL.Path, strconv.Itoa(c.Writer.Status())}
			m.With(labels...).Add(1).Observe(float64(time.Since(n).Milliseconds()))
		}()
		c.Next()
	}
}
