package ginx

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	router := gin.New()
	router.Use(
		func(c *gin.Context) {
			http.TimeoutHandler(
				http.HandlerFunc(func(http.ResponseWriter, *http.Request) { c.Next() }),
				1*time.Millisecond,
				TimeoutReason,
			).ServeHTTP(c.Writer, c.Request)
		},
	).GET("/test", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("GET", time.Now().Unix())
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	router.Run(":8080")
}
