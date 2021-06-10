package redis

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
)

type RedisType string

type Conf struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Address  []string `json:"address"`
	Master   string   `json:"master"`
	Database int      `json:"database"`
	Tls      bool     `json:"tls"`
}

func (cfg *Conf) TLSConfig() *tls.Config {
	if cfg.Tls {
		return &tls.Config{
			InsecureSkipVerify: false,
		}
	}
	return nil
}

func (cfg *Conf) UniversalOptions() *redis.UniversalOptions {
	opt := &redis.UniversalOptions{
		Addrs:     cfg.Address,
		Username:  cfg.Username,
		Password:  cfg.Password,
		DB:        cfg.Database,
		TLSConfig: cfg.TLSConfig(),
	}
	return opt
}

func (cfg *Conf) NewClient() Client {
	return NewClient(cfg)
}
