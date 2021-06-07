package mongo

import (
	"fmt"
	"strings"
)

type Conf struct {
	Username        string            `json:"username"`
	Password        string            `json:"password"`
	Address         []string          `json:"address"`
	Database        string            `json:"database"`
	Options         map[string]string `json:"options"`
	MaxConnIdleTime int               `json:"max_conn_idle_time"`
	MaxPoolSize     uint64            `json:"max_pool_size"`
	MinPoolSize     uint64            `json:"min_pool_size"`
}

func DefaultConf() *Conf {
	return &Conf{
		Username:        "admin",
		Password:        "admin",
		Address:         []string{"127.0.0.1:27017"},
		Database:        "test",
		MaxConnIdleTime: 60,
		MaxPoolSize:     100,
		MinPoolSize:     10,
	}
}

func (cfg *Conf) DSN() string {
	address := strings.Join(cfg.Address, ",")
	var opts []string
	for k, v := range cfg.Options {
		opts = append(opts, k+"="+v)
	}
	opt := strings.Join(opts, "&")
	return fmt.Sprintf("mongdb://%s:%s@%s/%s?%s", cfg.Username, cfg.Password, address, cfg.Database, opt)
}

func (cfg *Conf) NewClient() Client {
	return NewClient(cfg)
}
