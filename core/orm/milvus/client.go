package milvus

import (
	"context"
	"errors"
	"github.com/milvus-io/milvus-sdk-go/milvus"
	"github.com/silenceper/pool"
	"time"
)

var (
	ErrConnectClosed = errors.New("connection is closed")
	ErrClientInvalid = errors.New("client invalid")
)

type (
	Client struct {
		cfg  *Config
		pool pool.Pool
	}

	// Option set client config
	Option func(*Config)
)

// NewClient
func NewClient(ctx context.Context, options ...Option) (*Client, error) {
	cfg := DefaultConfig()
	for _, opt := range options {
		opt(cfg)
	}
	p, err := newPool(ctx, cfg)
	if err != nil {
		return nil, err
	}
	client := &Client{cfg: cfg, pool: p}
	return client, nil
}

// Config
func (c *Client) CFG() *Config {
	return c.cfg
}

// ORM
func (c *Client) ORM() (milvus.MilvusClient, error) {
	mc, err := c.pool.Get()
	if err != nil {
		return nil, err
	}
	if cli, ok := mc.(milvus.MilvusClient); ok {
		return cli, nil
	} else {
		return nil, ErrClientInvalid
	}
}

// newPool
func newPool(ctx context.Context, cfg *Config) (pool.Pool, error) {
	return pool.NewChannelPool(&pool.Config{
		InitialCap: cfg.PoolSize,
		MaxCap:     cfg.MaxConn,
		MaxIdle:    cfg.MaxIdle,
		Factory: func() (interface{}, error) {
			return newClient(ctx, cfg.Address)
		},
		Close: func(i interface{}) error {
			if client, ok := i.(milvus.MilvusClient); !ok {
				return ErrClientInvalid
			} else {
				return client.Disconnect(ctx)
			}
		},
		Ping: func(i interface{}) error {
			if client, ok := i.(milvus.MilvusClient); !ok {
				return ErrClientInvalid
			} else if !client.IsConnected(ctx) {
				return ErrConnectClosed
			}
			return nil
		},
		IdleTimeout: time.Duration(cfg.Timeout) * time.Second,
	})
}

// newClient
func newClient(ctx context.Context, address string) (milvus.MilvusClient, error) {
	host, port := resolverAddr(address)
	param := milvus.ConnectParam{IPAddress: host, Port: port}
	cli, err := milvus.NewMilvusClient(ctx, param)
	if err != nil {
		return nil, err
	}
	if cli.IsConnected(ctx) == false {
		return nil, ErrConnectClosed
	}
	return cli, err
}
