package middleware

import "github.com/gin-gonic/gin"

// NoCache: 去掉浏览器缓存
func NoCacheHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")
		ctx.Next()
	}
}
