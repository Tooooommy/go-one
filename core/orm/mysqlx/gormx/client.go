package gormx

import (
	"github.com/Tooooommy/go-one/core/orm/mysqlx"
	"github.com/Tooooommy/go-one/core/zapx/gormx"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Client struct {
	cfg mysqlx.Config
	orm *gorm.DB
}

func Connect(client *mysqlx.Client) (*Client, error) {
	cfg := client.CFG()
	db := client.ORM()
	log := logger.New(gormx.NewLogger(zapcore.InfoLevel), logger.Config{
		SlowThreshold:             0,
		Colorful:                  true,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Info,
	})
	orm, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysqlx",
		Conn:       db,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: cfg.Prefix,
		},
		Logger:   log,
		ConnPool: db,
	})
	return &Client{
		cfg: cfg,
		orm: orm,
	}, err
}

// NewClient
func NewClient(cfg mysqlx.Config) (*Client, error) {
	cli, err := mysqlx.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return Connect(cli)
}

// ORM
func (c *Client) ORM() *gorm.DB {
	return c.orm
}

// CFG
func (c *Client) CFG() mysqlx.Config {
	return c.cfg
}

func (c *Client) Close() error {
	db, err := c.orm.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
