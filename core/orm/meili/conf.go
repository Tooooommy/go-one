package meili

import "fmt"

type Conf struct {
	Address  string `json:"address"`
	ApiKey   string `json:"api_key"`
	PoolSize int    `json:"pool_size"`
	MaxConn  int    `json:"max_conn"`
	MaxIdle  int    `json:"max_idle"`
	Timeout  int    `json:"timeout"`
}

// DSN
func (cfg *Conf) DSN() string {
	return fmt.Sprintf("%s%s", cfg.Address, cfg.ApiKey)
}

// NewClient
func (cfg *Conf) NewClient() Client {
	return NewClient(cfg)
}
