package routes

import (
	"github.com/gin-gonic/gin"

	"go-redis-sample/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.SetTrustedProxies(nil)

	userProfileRoutes := r.Group("/user-profile")
	{
		userProfileRoutes.POST("/:user-id", handlers.CreateOrUpdateUserProfile)
		userProfileRoutes.GET("/:user-id", handlers.GetUserProfile)
		userProfileRoutes.PUT("/:user-id", handlers.CreateOrUpdateUserProfile)
		userProfileRoutes.DELETE("/:user-id", handlers.DeleteUserProfile)
	}
}
