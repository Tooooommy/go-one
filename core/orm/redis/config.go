package redis

import (
	"crypto/tls"
	"github.com/Tooooommy/go-one/core/syncx"
	"github.com/go-redis/redis/v8"
)

type RedisType string

const (
	NodeType    RedisType = "node"
	ClusterType RedisType = "cluster"
)

type Config struct {
	Type        RedisType `json:"type"`
	Address     []string  `json:"address"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Database    int       `json:"database"`
	MaxRetries  int       `json:"max_retries"`
	PoolSize    int       `json:"pool_size"`
	MinIdleConn int       `json:"min_idle_conn"`
	Limit       int       `json:"limit"`
	Tls         bool      `json:"tls"`
}

func DefaultConfig() *Config {
	return &Config{
		Type:        NodeType,
		Address:     []string{"127.0.0.1:6379"},
		Database:    0,
		MaxRetries:  3,
		PoolSize:    10,
		MinIdleConn: 3,
		Limit:       0,
	}
}

func (cfg *Config) TLSConfig() *tls.Config {
	if cfg.Tls {
		return &tls.Config{
			InsecureSkipVerify: false,
		}
	}
	return nil
}

func (cfg *Config) RedisOptions() *redis.Options {
	opt := &redis.Options{
		Addr:         cfg.Address[0],
		Username:     cfg.Username,
		Password:     cfg.Password,
		DB:           cfg.Database,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConn,
		TLSConfig:    cfg.TLSConfig(),
	}
	if cfg.Limit > 0 {
		opt.Limiter = syncx.NewLimiter(cfg.Limit)
	}
	return opt
}

func (cfg *Config) ClusterOptions() *redis.ClusterOptions {
	opt := &redis.ClusterOptions{
		Addrs:        cfg.Address,
		Username:     cfg.Username,
		Password:     cfg.Password,
		MaxRetries:   cfg.MaxRetries,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConn,
		TLSConfig:    cfg.TLSConfig(),
	}
	return opt
}
