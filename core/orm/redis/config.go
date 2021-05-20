package redis

import (
	"crypto/tls"
	"fmt"
)

type (
	Config struct {
		Address            string `json:"address"`
		Username           string `json:"username"`
		Password           string `json:"password"`
		Database           int    `json:"database"`
		MaxRetries         int    `json:"max_retries"`
		DialTimeout        int    `json:"dial_timeout"`
		ReadTimeout        int    `json:"read_timeout"`
		WriteTimeout       int    `json:"write_timeout"`
		PoolSize           int    `json:"pool_size"`
		MinIdleConns       int    `json:"min_idle_conns"`
		MaxConnAge         int    `json:"max_conn_age"`
		PoolTimeout        int    `json:"pool_timeout"`
		IdleTimeout        int    `json:"idle_timeout"`
		IdleCheckFrequency int    `json:"idle_check_frequency"`
		Limit              int    `json:"limit"`
	}
)

func (cfg Config) DSN() string {
	return fmt.Sprintf("redis://%s:%s@%s/%d", cfg.Username, cfg.Password, cfg.Address, cfg.Database)
}

func (cfg Config) TLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}
