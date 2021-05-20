package meili

import "fmt"

type Config struct {
	Address string `json:"address"`
	ApiKey  string `json:"api_key"`
}

// DSN
func (cfg Config) DSN() string {
	return fmt.Sprintf("%s%s", cfg.Address, cfg.ApiKey)
}

// NewClient
func (cfg Config) NewClient() (*Client, error) {
	return NewClient(cfg)
}
