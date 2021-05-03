package ginx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestT(t *testing.T) {
	router := gin.New()
	router.Any("/test", func(ctx *gin.Context) {
		fmt.Printf("%+v\n", ctx.Params)
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	router.Any("/param/:id", func(ctx *gin.Context) {
		fmt.Printf("%+v\n", ctx.Params)
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	router.Run(":8080")
}
