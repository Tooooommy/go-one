package ginx

import (
	"github.com/Tooooommy/go-one/core/logx"
	"github.com/Tooooommy/go-one/core/metrics"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Nop() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// Cors: 跨域请求
func Cors(origins ...string) gin.HandlerFunc {
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
func NoCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")
		ctx.Next()
	}
}

// Gunzip: gzip压缩处理
func Gunzip(level int, options ...gzip.Option) gin.HandlerFunc {
	return gzip.Gzip(level, options...)
}

var TimeoutReason = "Request Timeout"

// Timeout: 请求超时处理
func Timeout(duration time.Duration) gin.HandlerFunc {
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

// RequestId: 生成唯一值
func RequestId() gin.HandlerFunc {
	return requestid.New()
}

var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

// GetRealIp: 获取真实IP
func RealIp() gin.HandlerFunc {
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

// Recovery: panic恢复
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			logx.Error().Any("Recovery Error: %+v", err).Msg("GIN HTTP Panic")
			c.AbortWithStatus(http.StatusInternalServerError)
		}()
		c.Next()
	}
}

// StartTracing: 开启链路追踪
func StartTracing(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(name) > 0 {
			rf := kitopentracing.ContextToHTTP(opentracing.GlobalTracer(), logx.KitL())
			ctx := c.Request.Context()
			span := opentracing.SpanFromContext(ctx)
			if span == nil {
				span = opentracing.GlobalTracer().StartSpan(name)
				defer span.Finish()
			}
			ctx = opentracing.ContextWithSpan(ctx, span)
			c.Request = c.Request.WithContext(rf(ctx, c.Request))
		}
		c.Next()
	}
}

// StartPromxMetrics: 开启开启普罗米修斯
func StartPromxMetrics(cfg metrics.Config) gin.HandlerFunc {
	if len(cfg.Name) != 0 || len(cfg.Namespace) != 0 || len(cfg.Subsystem) != 0 {
		counter := metrics.NewPromxCounter(cfg)
		histogram := metrics.NewPromxHistogram(cfg)
		return func(c *gin.Context) {
			n := time.Now()
			defer func() {
				labels := []string{c.Request.Method, c.Request.Method, strconv.Itoa(c.Writer.Status())}
				counter.With(labels...).Add(1)
				histogram.With(labels...).Observe(float64(time.Since(n).Milliseconds()))
			}()
			c.Next()
		}
	}
	return Nop()
}
