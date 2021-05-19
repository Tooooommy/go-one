package gormx

import "github.com/Tooooommy/go-one/core/orm/mysqlx"

func GetCahceConn(cfg mysqlx.Config) (*Client, error) {
	client, err := mysqlx.GetCacheConn(cfg)
	if err != nil {
		return nil, err
	}
	return Connect(client)
}
