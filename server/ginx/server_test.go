package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestT(t *testing.T) {
	router := gin.New()
	router.Any("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	router.Run(":8080")
}
