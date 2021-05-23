package redis

import (
	"context"
	"fmt"
	"github.com/Tooooommy/go-one/core/syncx"
	"github.com/go-redis/redis/v8"
	"io"
)

type (
	Node interface {
		redis.Cmdable
		io.Closer
		redis.Scripter
		redis.UniversalClient
	}
	Client struct {
		cfg Config
		orm Node
	}
)

// NewClient
func NewClient(cfg Config) (*Client, error) {
	var cli Node
	switch cfg.RedisType {
	case NodeType:
		cli = redis.NewClient(&redis.Options{
			Addr:         cfg.Address[0],
			Username:     cfg.Username,
			Password:     cfg.Password,
			DB:           cfg.Database,
			MaxRetries:   maxRetries,
			MinIdleConns: minIdleConns,
			TLSConfig:    cfg.TLSConfig(),
			Limiter:      syncx.NewLimiter(cfg.Limit),
		})
	case ClusterType:
		cli = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        cfg.Address,
			Username:     cfg.Username,
			Password:     cfg.Password,
			MaxRetries:   maxRetries,
			MinIdleConns: minIdleConns,
			TLSConfig:    cfg.TLSConfig(),
		})
	default:
		return nil, fmt.Errorf("redis type '%s' is not supported", cfg.RedisType)
	}

	client := &Client{
		cfg: cfg,
		orm: cli,
	}
	err := client.Ping(context.Background())
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
func (c *Client) CFG() Config {
	return c.cfg
}

// Close
func (c *Client) Close() error {
	return c.orm.Close()
}
