package mysqlx

import (
	"database/sql"
	"github.com/Tooooommy/go-one/core/zapx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"time"
)

var (
	global *Client
	gdb    *gorm.DB
	xdb    *sqlx.DB
)

// Client
type Client struct {
	cfg Config
	raw *sql.DB
}

// ping
func (c *Client) ping(duration int) {
	defer func() {
		if result := recover(); result != nil {
			zapx.Error().Any("Recover Result", result).
				Msg("mysqlx ping function recover")
		}
		c.ping(duration)
	}()
	for {
		time.Sleep(time.Duration(duration) * time.Second)
		err := c.raw.Ping()
		if err != nil {
			c.raw = nil
			zapx.Error().Error(err).Msg("mysqlx database ping occurred error")
		}
	}
}

// NewClient
func NewClient(cfg Config) (*Client, error) {
	db, err := sql.Open("mysqlx", cfg.DSN())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	client := &Client{
		cfg: cfg,
		raw: db,
	}
	go client.ping(cfg.PingDuration)
	return client, nil
}

// DB
func (c *Client) DB() *sql.DB {
	return c.raw
}

// Config
func (c *Client) Config() Config {
	return c.cfg
}

// Init
func Init(client *Client) {
	global = client
}

// Global
func Global() *Client {
	return global
}
