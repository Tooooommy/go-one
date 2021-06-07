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
func (c *Conf) DSN() string {
	return fmt.Sprintf("%s%s", c.Address, c.ApiKey)
}

// NewClient
func (c *Conf) NewClient() Client {
	return NewClient(c)
}
