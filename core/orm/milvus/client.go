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
	Client interface {
		Conn() (milvus.MilvusClient, error)
	}
	client struct {
		cfg  *Conf
		pool pool.Pool
	}

	// Option set client config
	Option func(*Conf)
)

// NewClient
func NewClient(cfg *Conf) (Client, error) {
	client := &client{cfg: cfg}
	err := client.initPool()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// ORM
func (c *client) Conn() (milvus.MilvusClient, error) {
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
func (c *client) initPool() (err error) {
	c.pool, err = pool.NewChannelPool(&pool.Config{
		InitialCap: c.cfg.PoolSize,
		MaxCap:     c.cfg.MaxConn,
		MaxIdle:    c.cfg.MaxIdle,
		Factory: func() (interface{}, error) {
			return c.getClient()
		},
		Close: func(i interface{}) error {
			if client, ok := i.(milvus.MilvusClient); !ok {
				return ErrClientInvalid
			} else {
				return client.Disconnect(context.Background())
			}
		},
		Ping: func(i interface{}) error {
			if client, ok := i.(milvus.MilvusClient); !ok {
				return ErrClientInvalid
			} else if !client.IsConnected(context.Background()) {
				return ErrConnectClosed
			}
			return nil
		},
		IdleTimeout: time.Duration(c.cfg.Timeout) * time.Second,
	})
	return
}

// getClient
func (c *client) getClient() (milvus.MilvusClient, error) {
	ctx := context.Background()
	host, port := resolverAddr(c.cfg.Address)
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
