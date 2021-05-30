package mysqlx

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Client
type (
	Client struct {
		cfg *Config
		orm *sql.DB
	}

	Option func(cfg *Config)
)

// NewClient
func NewClient(options ...Option) (*Client, error) {
	cfg := DefaultConfig()
	for _, opt := range options {
		opt(cfg)
	}

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		return nil, err
	}

	if cfg.ConnMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	}
	if cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	client := &Client{cfg: cfg, orm: db}
	err = client.Ping()
	return client, err
}

// ORM
func (c *Client) ORM() *sql.DB {
	return c.orm
}

// CFG
func (c *Client) CFG() *Config {
	return c.cfg
}

// Ping
func (c *Client) Ping() error {
	return c.orm.Ping()
}

// Close
func (c *Client) Close() error {
	return c.orm.Close()
}
