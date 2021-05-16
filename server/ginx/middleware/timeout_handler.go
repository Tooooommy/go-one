package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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
