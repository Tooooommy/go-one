package redis

import "github.com/go-redis/redis/v8"

type RedisConfig struct {
}

// 集群
func InitRedis(cfd RedisConfig) *redis.Client {
	return nil
}
