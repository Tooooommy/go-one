package meili

import (
	meili "github.com/meilisearch/meilisearch-go"
)

type (
	Client struct {
		cfg Config
		orm meili.ClientInterface
	}
	ClientOption func(*Config)
)

// NewClient
func NewClient(cfg Config) (*Client, error) {
	cli := meili.NewClient(meili.Config{
		Host:   cfg.Address,
		APIKey: cfg.ApiKey,
	})
	err := cli.Health().Get()
	return &Client{
		cfg: cfg,
		orm: cli,
	}, err
}

// ORM
func (c *Client) ORM() meili.ClientInterface {
	return c.orm
}

// CFG
func (c *Client) CFG() Config {
	return c.cfg
}
