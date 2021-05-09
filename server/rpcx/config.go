package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/server"
)

type Config struct {
	server.Config
	Discov discov.Config
}
