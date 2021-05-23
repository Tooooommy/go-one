package meili

import (
	"github.com/meilisearch/meilisearch-go"
)

type (
	Client struct {
		cfg Config
		orm meilisearch.ClientInterface
	}
)

// NewClient
func NewClient(cfg Config) (*Client, error) {
	cli := meilisearch.NewClient(meilisearch.Config{
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
func (c *Client) ORM() meilisearch.ClientInterface {
	return c.orm
}

// CFG
func (c *Client) CFG() Config {
	return c.cfg
}
