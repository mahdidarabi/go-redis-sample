package persistence

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	"go-redis-sample/internal/config"
)

func NewRedisClient(redisConfig config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Username: redisConfig.Username,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
}
