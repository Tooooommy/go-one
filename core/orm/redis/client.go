package redis

import (
	"context"
	"github.com/Tooooommy/go-one/core/syncx"
	"strings"

	"github.com/go-redis/redis/v8"
)

type (
	Client interface {
		Conn() (redis.UniversalClient, error)
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

func (c *client) getConn() (redis.UniversalClient, error) {
	addr := strings.Join(c.cfg.Address, ",")
	val, ok := manager.Get(addr)
	if ok {
		return val.(redis.UniversalClient), nil
	}
	cli := redis.NewUniversalClient(c.cfg.UniversalOptions())
	cli.AddHook(&TracingHook{})
	err := cli.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	manager.Set(addr, cli)
	return cli, err
}

// Conn
func (c *client) Conn() (redis.UniversalClient, error) {
	return c.getConn()
}
