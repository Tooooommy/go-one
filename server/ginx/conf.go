package ginx

import (
	"github.com/Tooooommy/go-one/server/conf"
)

type ServerConf struct {
	conf.Conf
	MaxConns int   `json:"max_conns"`
	MaxBytes int64 `json:"max_bytes"`
	Timeout  int64 `json:"timeout"`
}
