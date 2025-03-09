package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "go-redis-sample/docs"

	"go-redis-sample/internal/config"
	"go-redis-sample/internal/middleware"
	"go-redis-sample/internal/models"
	"go-redis-sample/internal/persistence"
	"go-redis-sample/internal/routes"
)

// @title Go Redis Sample API
// @version 1.0
// @description This is a sample API using Go, Gin, and Redis
// @BasePath /
func main() {
	godotenv.Load()

	appConfig := config.LoadConfig()

	redisClient := persistence.NewRedisClient(appConfig.RedisConfig)
	models.SetCacheExpirationTime(time.Duration(appConfig.CacheExpirationTime) * time.Second)
	models.SetRedisClient(redisClient)

	logLevel := appConfig.LogLevel
	if logLevel == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Use(middleware.LoggingMiddleware())

	serverAddr := fmt.Sprintf(":%s", appConfig.Port)
	log.Println("Server is running on port " + serverAddr)
	r.Run(serverAddr)
}
