package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// GunzipHandler: gzip压缩处理
func GunzipHandler(level int, options ...gzip.Option) gin.HandlerFunc {
	return gzip.Gzip(level, options...)
}
