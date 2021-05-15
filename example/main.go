package main

import (
	"github.com/Tooooommy/go-one/server"
	"github.com/Tooooommy/go-one/server/ginx"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := ginx.Config{
		Config: server.Config{
			Name: "go-one",
			Host: "127.0.0.1",
			Port: 8080,
		},
		MaxConns: 1000,
		MaxBytes: 1000,
		Timeout:  10000,
	}

	router.GET("/get")

	svr := ginx.NewServer(
		ginx.WithConfig(cfg),
		ginx.WithGinEngine(router),
	)
	err := svr.Start()
	if err != nil {
		panic(err)
	}
}
