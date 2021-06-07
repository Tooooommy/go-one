package mongo

import (
	"context"
	"github.com/Tooooommy/go-one/core/syncx"
	"go.mongodb.org/mongo-driver/mongo"
	mgoptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type (
	Client interface {
		Conn() (*mongo.Client, error)
	}

	client struct {
		cfg *Conf
	}
)

var (
	manager = syncx.NewManager()
)

// NewClient
func NewClient(cfg *Conf) Client {
	return &client{cfg: cfg}
}

func (c *client) getConn() (*mongo.Client, error) {
	dsn := c.cfg.DSN()
	val, ok := manager.Get(dsn)
	if ok {
		return val.(*mongo.Client), nil
	}
	opt := mgoptions.Client().ApplyURI(dsn)
	if c.cfg.MaxConnIdleTime > 0 {
		opt.SetMaxConnIdleTime(time.Duration(c.cfg.MaxConnIdleTime) * time.Millisecond)
	}
	if c.cfg.MaxPoolSize > 0 {
		opt.SetMaxPoolSize(c.cfg.MaxPoolSize)
	}
	if c.cfg.MinPoolSize > 0 {
		opt.SetMinPoolSize(c.cfg.MinPoolSize)
	}

	cli, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}
	err = cli.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	manager.Set(dsn, cli)
	return cli, err
}

// Conn
func (c *client) Conn() (*mongo.Client, error) {
	return c.getConn()
}
