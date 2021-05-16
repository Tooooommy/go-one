package middleware

import (
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

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
