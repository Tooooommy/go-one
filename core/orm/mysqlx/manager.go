package mysqlx

import (
	"github.com/Tooooommy/go-one/core/syncx"
)

var (
	manager = syncx.NewConnManger()
)

// GetCacheConn
func GetCacheConn(cfg Config) (*Client, error) {
	dsn := cfg.DSN()
	val, exist := manager.Get(dsn)
	if exist {
		return val.(*Client), nil
	}
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	manager.Set(dsn, client)
	return client, nil
}
