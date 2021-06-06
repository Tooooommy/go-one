package rpcx

import (
	"github.com/Tooooommy/go-one/core/discov"
	"github.com/Tooooommy/go-one/server/conf"
)

type (
	ServerConf struct {
		conf.Conf
		Etcd discov.Config `json:"etcd"`
	}

	ClientConf struct {
		Retries int           `json:"retries"`
		Timeout int64         `json:"timeout"`
		Token   string        `json:"token"`
		Address string        `json:"address"`
		Etcd    discov.Config `json:"etcd"`
	}
)
