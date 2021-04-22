package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Handle(http.MethodGet, "/test", func(context *gin.Context) {

	})
}
