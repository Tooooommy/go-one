package sqlx

import (
	"github.com/Tooooommy/go-one/core/orm/mysqlx"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	cfg *mysqlx.Config
	orm *sqlx.DB
}

// Connect
func Connect(client *mysqlx.Client) *Client {
	cfg := client.CFG()
	raw := client.ORM()
	orm := sqlx.NewDb(raw, "mysql")
	return &Client{
		cfg: cfg,
		orm: orm,
	}
}

// ORM
func (c *Client) ORM() *sqlx.DB {
	return c.orm
}

// CFG
func (c *Client) CFG() *mysqlx.Config {
	return c.cfg
}

// Close
func (c *Client) Close() error {
	return c.orm.Close()
}
