package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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
