package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs incoming requests and the time taken to process them.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)

		log.Printf("[%s] %s %s - %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP(), elapsed)
	}
}
