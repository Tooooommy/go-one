package sqlx

import (
	"github.com/Tooooommy/go-one/core/orm/mysqlx"
)

// GetCacheConn
func GetCacheConn(cfg mysqlx.Config) (*Client, error) {
	mc, err := mysqlx.GetCacheConn(cfg)
	if err != nil {
		return nil, err
	}
	return Connect(mc), nil
}
