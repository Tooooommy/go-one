package meili

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
)

// Pool
type (
	Config struct {
		Address string `json:"address"`
		ApiKey  string `json:"api_key"`
	}
	Client struct {
		cfg Config
		cli meilisearch.ClientInterface
	}
)

// DSN
func (cfg Config) DSN() string {
	return fmt.Sprintf("%s%s", cfg.Address, cfg.ApiKey)
}

// NewClient
func (cfg Config) NewClient() (*Client, error) {
	return NewClient(cfg)
}

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
