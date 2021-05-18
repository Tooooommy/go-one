package sqlx

import (
	"github.com/Tooooommy/go-one/core/orm/mysqlx"
	"github.com/jmoiron/sqlx"
)

var (
	global *Client
)

type Client struct {
	cfg mysqlx.Config
	orm *sqlx.DB
}

func Connect(client *mysqlx.Client) *Client {
	cfg := client.Config()
	raw := client.DB()
	orm := sqlx.NewDb(raw, "mysql")
	return &Client{
		cfg: cfg,
		orm: orm,
	}
}

func NewClient(cfg mysqlx.Config) (*Client, error) {
	cli, err := mysqlx.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	client := Connect(cli)
	return client, nil
}

func Init(client *Client) {
	global = client
}

func Global() *Client {
	return global
}

func GetSqlxAuto() *sqlx.DB {
	return global.orm
}
