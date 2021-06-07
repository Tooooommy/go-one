package redis

import (
	"context"
	"fmt"
	"github.com/Tooooommy/go-one/core/syncx"
	"io"
	"strings"

	"github.com/go-redis/redis/v8"
)

type (
	Node interface {
		redis.Cmdable
		io.Closer
		redis.Scripter
		redis.UniversalClient
	}

	Client interface {
		Conn() Node
	}

	client struct {
		cfg *Conf
		orm Node
	}
	Option func(cfg *Conf)
)

var (
	manager = syncx.NewManager()
)

// NewClient
func NewClient(cfg *Conf) *client {
	return &client{cfg: cfg}
}

func (c *client) getConn() (Node, error) {
	addr := strings.Join(c.cfg.Address, ",")
	val, ok := manager.Get(addr)
	if ok {
		return val.(Node), nil
	}
	var node Node
	switch c.cfg.Type {
	case NodeType:
		node = redis.NewClient(c.cfg.RedisOptions())
		node.AddHook(&TracingHook{})
	case ClusterType:
		opt := c.cfg.ClusterOptions()
		opt.NewClient = func(opt *redis.Options) *redis.Client {
			node := redis.NewClient(opt)
			node.AddHook(&TracingHook{})
			return node
		}
		node = redis.NewClusterClient(opt)
	default:
		return nil, fmt.Errorf("redis type '%s' is not supported", c.cfg.Type)
	}
	err := node.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	manager.Set(addr, node)
	return node, err
}

// Conn
func (c *client) Conn() (Node, error) {
	return c.getConn()
}
