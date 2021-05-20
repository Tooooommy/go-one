package milvus

import (
	"context"
	"errors"
	"github.com/milvus-io/milvus-sdk-go/milvus"
	"strings"
)

var (
	ErrMilvusPing = errors.New("milvus ping occurred error")
)

type (
	Config struct {
		Address string `json:"address"`
	}

	Client struct {
		cfg Config
		cli milvus.MilvusClient
	}
)

func (cfg Config) DSN() string {
	return cfg.Address
}

func (cfg Config) NewClient() (*Client, error) {
	return NewClient(cfg)
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

func resolverAddr(address string) (host, port string) {
	ss := strings.Split(address, ":")
	if len(ss) >= 2 {
		host = ss[0]
		port = ss[1]
	} else if len(ss[0]) == 1 {
		port = "19530"
		if ss[0] == "" {
			host = "localhost"
		}
	}
	return
}
