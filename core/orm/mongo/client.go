package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	mgoptions "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type (
	Client struct {
		cfg *Config
		orm *mongo.Client
	}
	Option func(cfg *Config)
)

func NewClient(options ...Option) (*Client, error) {
	cfg := DefaultConfig()
	for _, opt := range options {
		opt(cfg)
	}

	opt := mgoptions.Client().ApplyURI(cfg.DSN())
	if cfg.MaxConnIdleTime > 0 {
		opt.SetMaxConnIdleTime(time.Duration(cfg.MaxConnIdleTime) * time.Millisecond)
	}
	if cfg.MaxPoolSize > 0 {
		opt.SetMaxPoolSize(cfg.MaxPoolSize)
	}
	if cfg.MinPoolSize > 0 {
		opt.SetMinPoolSize(cfg.MinPoolSize)
	}

	cli, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}

	client := &Client{cfg: cfg, orm: cli}
	err = client.Ping()
	return client, err
}

// Ping
func (c *Client) Ping() error {
	return c.orm.Ping(context.Background(), readpref.Primary())
}

func (c *Client) Close() error {
	return c.orm.Disconnect(context.Background())
}

// ORM
func (c *Client) ORM() *mongo.Client {
	return c.orm
}

// CFG
func (c *Client) CFG() *Config {
	return c.cfg
}
