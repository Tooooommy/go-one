package middleware

import (
	"github.com/Tooooommy/go-one/core/metrics"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

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
