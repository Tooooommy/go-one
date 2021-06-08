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

type Conf struct {
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

func DefaultConf() *Conf {
	return &Conf{
		Type:        NodeType,
		Address:     []string{"127.0.0.1:6379"},
		Database:    0,
		MaxRetries:  3,
		PoolSize:    100,
		MinIdleConn: 64,
		Limit:       0,
	}
}

func (cfg *Conf) TLSConfig() *tls.Config {
	if cfg.Tls {
		return &tls.Config{
			InsecureSkipVerify: false,
		}
	}
	return nil
}

func (cfg *Conf) RedisOptions() *redis.Options {
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

func (cfg *Conf) ClusterOptions() *redis.ClusterOptions {
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

func (cfg *Conf) NewClient() Client {
	return NewClient(cfg)
}
