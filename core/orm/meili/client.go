package meili

import (
	"github.com/meilisearch/meilisearch-go"
)

// Pool
type (
	Client struct {
		cfg Config
		cli meilisearch.ClientInterface
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
		cli: cli,
	}, err
}

func (c *Client) ORM() meilisearch.ClientInterface {
	return c.cli
}

func (c *Client) CFG() Config {
	return c.cfg
}
