package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov/etcdx"
	"github.com/Tooooommy/go-one/server"
)

type Config struct {
	server.Config
	Etcd etcdx.Config `json:"discovery"`
}
