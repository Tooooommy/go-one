package mysqlx

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

// Client
type Client struct {
	cfg Config
	orm *sql.DB
	one sync.Once
}

// Ping
func (c *Client) Ping() error {
	return c.orm.Ping()
}

// NewClient
func NewClient(cfg Config) (*Client, error) {
	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	client := &Client{
		cfg: cfg,
		orm: db,
	}
	client.one.Do(func() {
		err = client.Ping()
	})
	return client, err
}

// ORM
func (c *Client) ORM() *sql.DB {
	return c.orm
}

// CFG
func (c *Client) CFG() Config {
	return c.cfg
}
