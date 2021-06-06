package ginx

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestT(t *testing.T) {
	router := gin.New()
	router.Run(":8080")
}
