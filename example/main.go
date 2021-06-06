package main

import (
	"github.com/Tooooommy/go-one/server/conf"
	"github.com/Tooooommy/go-one/server/ginx"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := ginx.Conf{
		Conf: conf.Conf{
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
		ginx.WithEngine(router),
	)
	err := svr.Start()
	if err != nil {
		panic(err)
	}
}
