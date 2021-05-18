package mongo

import (
	"fmt"
	"strings"
)

type Config struct {
	Username        string            `json:"username"`
	Password        string            `json:"password"`
	Address         []string          `json:"address"`
	Database        string            `json:"database"`
	Options         map[string]string `json:"options"`
	PingDuration    int               `json:"ping_duration"`
	MaxConnIdleTime int               `json:"max_conn_idle_time"`
	MaxPoolSize     uint64            `json:"max_pool_size"`
	MinPoolSize     uint64            `json:"min_pool_size"`
}

func (cfg Config) DSN() string {
	username := cfg.Username
	password := cfg.Password
	address := strings.Join(cfg.Address, ",")
	database := "test"

	if len(cfg.Address) <= 0 {
		address = "127.0.0.1:27017"
	}
	if cfg.Database != "" {
		database = cfg.Database
	}
	var opts []string
	for k, v := range cfg.Options {
		opts = append(opts, k+"="+v)
	}
	opt := strings.Join(opts, "&")
	return fmt.Sprintf("mongdb://%s:%s@%s/%s?%s", username, password, address, database, opt)
}
