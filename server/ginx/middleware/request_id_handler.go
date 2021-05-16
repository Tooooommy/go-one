package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

// RequestIDHandler: 生成唯一值
func RequestIDHandler() gin.HandlerFunc {
	return requestid.New()
}
