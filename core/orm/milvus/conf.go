package milvus

import (
	"strings"
)

type Conf struct {
	Address  string `json:"address"`
	PoolSize int    `json:"pool_size"`
	MaxConn  int    `json:"max_conn"`
	MaxIdle  int    `json:"max_idle"`
	Timeout  int    `json:"timeout"`
}

func DefaultConf() *Conf {
	return &Conf{
		Address:  "127.0.0.1:19530",
		PoolSize: 5,
		MaxConn:  5,
		MaxIdle:  5,
		Timeout:  5,
	}
}

// DSN unique for client
func (cfg *Conf) DSN() string {
	return cfg.Address
}

func resolverAddr(address string) (host, port string) {
	ss := strings.Split(address, ":")
	if len(ss) >= 2 {
		host = ss[0]
		port = ss[1]
	} else if len(ss[0]) == 1 {
		port = "19530"
		if ss[0] == "" {
			host = "localhost"
		}
	}
	return
}
