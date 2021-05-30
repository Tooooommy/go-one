package meili

import "fmt"

type Config struct {
	Address  string `json:"address"`
	ApiKey   string `json:"api_key"`
	PoolSize int    `json:"pool_size"`
	MaxConn  int    `json:"max_conn"`
	MaxIdle  int    `json:"max_idle"`
	Timeout  int    `json:"timeout"`
}

// DSN
func (cfg *Config) DSN() string {
	return fmt.Sprintf("%s%s", cfg.Address, cfg.ApiKey)
}

// NewClient
func (cfg *Config) NewClient() (*Client, error) {
	return NewClient(cfg)
}
