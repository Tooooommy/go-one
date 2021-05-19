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
	host, port := resolverAddr(cfg.Address)
	cli, err := milvus.NewMilvusClient(
		context.Background(),
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
	err = client.Ping()
	return client, err
}

func (c *Client) Ping() error {
	if c.cli.IsConnected(context.Background()) == false {
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
