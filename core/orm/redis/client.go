package redis

import (
	"context"
	"fmt"
	"io"

	"github.com/go-redis/redis/v8"
)

type (
	Node interface {
		redis.Cmdable
		io.Closer
		redis.Scripter
		redis.UniversalClient
	}

	Client struct {
		cfg *Config
		orm Node
	}
	Option func(cfg *Config)
)

// NewClient
func NewClient(ctx context.Context, options ...Option) (*Client, error) {
	cfg := DefaultConfig()
	for _, opt := range options {
		opt(cfg)
	}

	var cli Node
	switch cfg.Type {
	case NodeType:
		cli = redis.NewClient(cfg.RedisOptions())
		cli.AddHook(&TracingHook{})
	case ClusterType:
		opt := cfg.ClusterOptions()
		opt.NewClient = func(opt *redis.Options) *redis.Client {
			node := redis.NewClient(opt)
			node.AddHook(&TracingHook{})
			return node
		}
		cli = redis.NewClusterClient(opt)
	default:
		return nil, fmt.Errorf("redis type '%s' is not supported", cfg.Type)
	}
	client := &Client{cfg: cfg, orm: cli}
	err := client.Ping(ctx)
	return client, err
}

// Ping
func (c *Client) Ping(ctx context.Context) error {
	return c.orm.Ping(ctx).Err()
}

// ORM
func (c *Client) ORM() Node {
	return c.orm
}

// CFG
func (c *Client) CFG() *Config {
	return c.cfg
}

// Close
func (c *Client) Close() error {
	return c.orm.Close()
}
