package middleware

import "github.com/gin-gonic/gin"

// NopHandler: 空中间件
func NopHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
