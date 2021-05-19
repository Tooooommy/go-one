package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

type Client struct {
	cfg Config
	orm *mongo.Client
	one sync.Once
}

func NewClient(cfg Config) (*Client, error) {
	cli, err := newMongoClient(cfg)
	if err != nil {
		return nil, err
	}
	client := &Client{
		cfg: cfg,
		orm: cli,
	}
	client.one.Do(func() {
		err = client.Ping()
	})
	return client, err
}

// ping
func (c *Client) Ping() error {
	return c.orm.Ping(context.Background(), readpref.Primary())
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

	return mongo.NewClient(opt)
}

func (c *Client) ORM() *mongo.Client {
	return c.orm
}

func (c *Client) CFG() Config {
	return c.cfg
}
