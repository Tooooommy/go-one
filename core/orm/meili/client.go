package meili

import (
	"github.com/Tooooommy/go-one/core/syncx"
	meili "github.com/meilisearch/meilisearch-go"
)

type (
	Client interface {
		Conn() (meili.ClientInterface, error)
	}
	client struct {
		cfg *Conf
	}
)

var (
	manager = syncx.NewManager()
)

// NewClient
func NewClient(cfg *Conf) Client {
	return &client{cfg: cfg}
}

// getClient
func (c *client) getClient() (meili.ClientInterface, error) {
	val, ok := manager.Get(c.cfg.Address)
	if ok {
		return val.(meili.ClientInterface), nil
	}
	cli := meili.NewClient(meili.Config{
		Host:   c.cfg.Address,
		APIKey: c.cfg.ApiKey,
	})
	err := cli.Health().Get()
	if err != nil {
		return nil, err
	}
	manager.Set(c.cfg.Address, cli)
	return cli, nil
}

// Conn
func (c *client) Conn() (meili.ClientInterface, error) {
	return c.getClient()
}
