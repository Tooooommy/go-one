package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/server"
)

type Config struct {
	server.Config
	Discov discov.Config
}

func (c Config) HaveEtcd() bool {
	return len(c.Discov.Hosts) > 0
}
