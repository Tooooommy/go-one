package redis

import (
	"crypto/tls"
	"fmt"
)

type RedisType string

const (
	minIdleConns           = 8
	maxRetries             = 3
	NodeType     RedisType = "node"
	ClusterType  RedisType = "cluster"
)

type (
	Config struct {
		RedisType RedisType `json:"redis_type"`
		Address   []string  `json:"address"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Database  int       `json:"database"`
		Limit     int       `json:"limit"`
	}
)

func (cfg Config) DSN() string {
	address := cfg.Address[0]
	return fmt.Sprintf("redis://%s:%s@%s/%d", cfg.Username, cfg.Password, address, cfg.Database)
}

func (cfg Config) TLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}
