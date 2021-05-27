package mysqlx

import (
	"github.com/Tooooommy/go-one/core/syncx"
)

var (
	manager = syncx.NewManger()
)

// GetCacheConn
func GetCacheConn(cfg Config) (*Client, error) {
	key := cfg.DSN()
	val, exist := manager.Get(key)
	if exist {
		return val.(*Client), nil
	}
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	manager.Set(key, client)
	return client, nil
}
