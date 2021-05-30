package milvus

import (
	"strings"
)

type Config struct {
	Address  string `json:"address"`
	PoolSize int    `json:"pool_size"`
	MaxConn  int    `json:"max_conn"`
	MaxIdle  int    `json:"max_idle"`
	Timeout  int    `json:"timeout"`
}

func DefaultConfig() *Config {
	return &Config{
		Address:  "127.0.0.1:19530",
		PoolSize: 5,
		MaxConn:  5,
		MaxIdle:  5,
		Timeout:  5,
	}
}

// SetAddress set milvus client address[host:port]
func SetAddress(address string) Option {
	return func(cfg *Config) {
		cfg.Address = address
	}
}

// SetPoolSize set pool initialize cap
func SetPoolSize(size int) Option {
	return func(cfg *Config) {
		cfg.PoolSize = size
	}
}

// SetMaxConn set connection max
func SetMaxConn(max int) Option {
	return func(cfg *Config) {
		cfg.MaxConn = max
	}
}

// SetTimeout set idle connection timeout
func SetTimeout(timeout int) Option {
	return func(cfg *Config) {
		cfg.Timeout = timeout
	}
}

// DSN unique for client
func (cfg *Config) DSN() string {
	return cfg.Address
}

func resolverAddr(address string) (host, port string) {
	ss := strings.Split(address, ":")
	if len(ss) >= 2 {
		host = ss[0]
		port = ss[1]
	} else if len(ss[0]) == 1 {
		port = "19530"
		if ss[0] == "" {
			host = "localhost"
		}
	}
	return
}
