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
	cli milvus.MilvusClient
}

func NewClient(cfg Config) (*Client, error) {
	ctx := context.Background()
	host, port := resolverAddr(cfg.Address)
	cli, err := milvus.NewMilvusClient(
		ctx,
		milvus.ConnectParam{
			IPAddress: host,
			Port:      port,
		},
	)
	if err != nil {
		return nil, err
	}
	client := &Client{
		cfg: cfg,
		cli: cli,
	}
	err = client.Ping(ctx)
	return client, err
}

func (c *Client) Ping(ctx context.Context) error {
	if c.cli.IsConnected(ctx) == false {
		return ErrMilvusPing
	}
	return nil
}
