package redis

import (
	"github.com/Tooooommy/go-one/core/syncx"
)

var (
	manager = syncx.NewManger()
)

func GetCacheConn(cfg Config) (*Client, error) {
	key := cfg.DSN()
	val, ok := manager.Get(key)
	if ok {
		return val.(*Client), nil
	}
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	manager.Set(key, client)
	return client, err
}
