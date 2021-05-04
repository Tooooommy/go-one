package ginx

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

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
		c.Next()
	}
}
