package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"go-redis-sample/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.SetTrustedProxies(nil)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	userProfileRoutes := r.Group("/user-profile")
	{

		userProfileRoutes.GET("/:user-id", handlers.GetUserProfile)
		userProfileRoutes.POST("/:user-id", handlers.CreateUserProfile)
		userProfileRoutes.PUT("/:user-id", handlers.UpdateUserProfile)
		userProfileRoutes.DELETE("/:user-id", handlers.DeleteUserProfile)
	}
}
