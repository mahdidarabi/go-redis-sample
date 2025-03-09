package config

import (
	"fmt"
	"strings"

	"go-redis-sample/internal/utils"
)

type AppConfig struct {
	Port     string
	LogLevel string

	CacheExpirationTime int

	RedisConfig RedisConfig
}

type RedisConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       int
}

// LoadConfig loads application configuration from environment variables
func LoadConfig() *AppConfig {
	port := utils.GetEnvWithDefault("HTTP_PORT", "3000")
	logLevel := utils.GetEnvWithDefault("LOG_LEVEL", "info")

	cacheExpirationTime := utils.ParseStringToInt(utils.GetEnvWithDefault("CACHE_EXPIRATION_TIME", "120"))

	// redisConfig := RedisConfig{
	// 	Host:     utils.GetEnvWithDefault("REDIS_HOST", "localhost"),
	// 	Port:     utils.GetEnvWithDefault("REDIS_PORT", "6379"),
	// 	Username: utils.GetEnvWithDefault("REDIS_USERNAME", "default"),
	// 	Password: utils.GetEnvWithDefault("REDIS_PASSWORD", ""),
	// 	DB:       utils.ParseStringToInt(utils.GetEnvWithDefault("REDIS_DB", "0")),
	// }

	redisHost := utils.GetEnvWithDefault("REDIS_HOST", "localhost")
	redisPort := utils.GetEnvWithDefault("REDIS_PORT", "6379")
	redisUsername := utils.GetEnvWithDefault("REDIS_USERNAME", "default")
	redisPassword := utils.GetEnvWithDefault("REDIS_PASSWORD", "")
	redisDB := utils.ParseStringToInt(utils.GetEnvWithDefault("REDIS_DB", "0"))

	fmt.Printf("AppConfig: Port: %s, LogLevel: %s, CacheExpirationTime: %d, RedisHost: %s, RedisPort: %s, RedisUsername: %s, RedisPassword: %s, RedisDB: %d\n", port, logLevel, cacheExpirationTime, redisHost, redisPort, redisUsername, strings.Repeat("*", len(redisPassword)), redisDB)

	return &AppConfig{
		port,
		logLevel,
		cacheExpirationTime,
		RedisConfig{
			Host:     redisHost,
			Port:     redisPort,
			Username: redisUsername,
			Password: redisPassword,
			DB:       redisDB,
		},
	}
}
