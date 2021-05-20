package milvus

import "github.com/Tooooommy/go-one/core/syncx"

var (
	manager = syncx.NewConnManger()
)

// GetCacheConn
func GetCacheConn(cfg Config) (*Client, error) {
	key := cfg.DSN()
	val, ok := manager.Get(key)
	if ok {
		return val.(*Client), nil
	}
	client, err := cfg.NewClient()
	if err != nil {
		return nil, err
	}
	manager.Set(key, client)
	return client, nil
}
