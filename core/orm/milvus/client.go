package milvus

import (
	"context"
	"errors"
	"github.com/milvus-io/milvus-sdk-go/milvus"
)

var (
	ErrMilvusPing = errors.New("milvus ping occurred error")
)

type Client struct {
	cfg Config
	orm milvus.MilvusClient
}

// NewClient
func NewClient(cfg Config) (*Client, error) {
	ctx := context.Background()
	host, port := resolverAddr(cfg.Address)
	cli, err := milvus.NewMilvusClient(ctx,
		milvus.ConnectParam{IPAddress: host, Port: port})
	if err != nil {
		return nil, err
	}
	client := &Client{
		cfg: cfg,
		orm: cli,
	}
	err = client.Ping(ctx)
	return client, err
}

// Ping
func (c *Client) Ping(ctx context.Context) error {
	if c.orm.IsConnected(ctx) == false {
		return ErrMilvusPing
	}
	return nil
}

// ORM
func (c *Client) ORM() milvus.MilvusClient {
	return c.orm
}

// CFG
func (c *Client) CFG() Config {
	return c.cfg
}
