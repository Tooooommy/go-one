package mongo

import (
	"context"
	"errors"
	"github.com/Tooooommy/go-one/core/task"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	ErrMongoPing = errors.New("mongo ping occurred error")
)

type Client struct {
	cfg Config
	cli *mongo.Client
}

func NewClient(cfg Config) (*Client, error) {
	cli, err := newMongoClient(cfg)
	if err != nil {
		return nil, err
	}
	client := &Client{
		cfg: cfg,
		cli: cli,
	}
	client.ping()
	return client, nil
}

// ping
func (c *Client) ping() {
	go task.TickHandler(c.cfg.PingDuration, func() error {
		err := c.cli.Ping(context.Background(), readpref.Primary())
		if err != nil {
			c.cli = nil
			return ErrMongoPing
		}
		return nil
	})

}

// newMongoClient
func newMongoClient(cfg Config) (*mongo.Client, error) {
	opt := options.Client().ApplyURI(cfg.DSN())
	if cfg.MaxConnIdleTime > 0 {
		opt.SetMaxConnIdleTime(time.Duration(cfg.MaxConnIdleTime) * time.Millisecond)
	}
	if cfg.MaxPoolSize > 0 {
		opt.SetMaxPoolSize(cfg.MaxPoolSize)
	}
	if cfg.MinPoolSize > 0 {
		opt.SetMinPoolSize(cfg.MinPoolSize)
	}

	return mongo.Connect(context.Background(), opt)
}
