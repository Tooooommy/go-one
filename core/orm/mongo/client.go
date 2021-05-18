package mongo

import (
	"context"
	"github.com/Tooooommy/go-one/core/zapx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	global *Client
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
	go client.ping(cfg.PingDuration)
	return client, nil
}

// ping
func (c *Client) ping(duration int) {
	defer func() {
		if result := recover(); result != nil {
			zapx.Error().Any("Recover Result", result).
				Msg("mongo ping function recover")
		}
		c.ping(duration)
	}()
	for {
		time.Sleep(time.Duration(duration) * time.Second)
		err := c.cli.Ping(context.Background(), readpref.Primary())
		if err != nil {
			c.cli, err = newMongoClient(c.cfg)
			zapx.Error().Error(err).Msg("mongo database ping occurred error")
		}
	}
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

// Init
func Init(client *Client) {
	global = client
}

// global
func Global() *Client {
	return global
}

func GetMongoAuto() *mongo.Client {
	return global.cli
}
