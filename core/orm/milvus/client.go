package milvus

import (
	"context"
	"errors"
	"github.com/Tooooommy/go-one/core/task"
	"github.com/milvus-io/milvus-sdk-go/milvus"
)

var (
	ErrMilvusPing = errors.New("milvus ping occurred error")
)

type (
	Config struct {
		Host         string `json:"host"`
		Port         string `json:"port"`
		PingDuration int64  `json:"ping_duration"`
	}

	Client struct {
		cfg Config
		cli milvus.MilvusClient
	}
)

func (cfg Config) NewClient() (*Client, error) {
	return NewClient(cfg)
}

func NewClient(cfg Config) (*Client, error) {
	cli, err := milvus.NewMilvusClient(
		context.Background(),
		milvus.ConnectParam{
			IPAddress: cfg.Host,
			Port:      cfg.Port,
		},
	)
	if err != nil {
		return nil, err
	}
	client := &Client{
		cfg: cfg,
		cli: cli,
	}
	client.ping()
	return client, nil
}

func (c *Client) ping() {
	go task.TickHandler(c.cfg.PingDuration, func() error {
		if c.cli.IsConnected(context.Background()) == false {
			c.cli = nil
			return ErrMilvusPing
		}
		return nil
	})
}
