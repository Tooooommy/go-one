package sqlx

import (
	"github.com/Tooooommy/go-one/core/orm/mysqlx"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	cfg mysqlx.Config
	orm *sqlx.DB
}

func Connect(client *mysqlx.Client) *Client {
	cfg := client.CFG()
	raw := client.ORM()
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

func (c *Client) ORM() *sqlx.DB {
	return c.orm
}

func (c *Client) CFG() mysqlx.Config {
	return c.cfg
}
