package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"go-redis-sample/internal/config"
	"go-redis-sample/internal/middleware"
	"go-redis-sample/internal/models"
	"go-redis-sample/internal/persistence"
	"go-redis-sample/internal/routes"
)

func main() {
	godotenv.Load()

	appConfig := config.LoadConfig()

	redisClient := persistence.NewRedisClient(appConfig.RedisConfig)
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
