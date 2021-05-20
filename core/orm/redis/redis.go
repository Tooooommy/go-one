package redis

import (
	"context"
	"github.com/Tooooommy/go-one/core/syncx"
	"github.com/go-redis/redis/v8"
	"time"
)

type (
	Client struct {
		cfg Config
		cli *redis.Client
	}
)

func NewClient(cfg Config) (*Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:               cfg.Address,
		Username:           cfg.Username,
		Password:           cfg.Password,
		DB:                 cfg.Database,
		MaxRetries:         cfg.MaxRetries,
		DialTimeout:        time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:        time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:       time.Duration(cfg.WriteTimeout) * time.Second,
		PoolSize:           cfg.PoolSize,
		MinIdleConns:       cfg.MinIdleConns,
		MaxConnAge:         time.Duration(cfg.MaxConnAge) * time.Second,
		PoolTimeout:        time.Duration(cfg.PoolTimeout) * time.Second,
		IdleTimeout:        time.Duration(cfg.IdleTimeout) * time.Second,
		IdleCheckFrequency: time.Duration(cfg.IdleCheckFrequency) * time.Second,
		TLSConfig:          cfg.TLSConfig(),
		Limiter:            syncx.NewLimiter(cfg.Limit),
	})
	client := &Client{
		cfg: cfg,
		cli: cli,
	}
	err := client.cli.Ping(context.Background()).Err()
	return client, err
}
