package handlers

import (
	"net/http"

	"go-redis-sample/internal/models"
	"go-redis-sample/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrUpdateUserProfile(c *gin.Context) {
	id := c.Param("user-id")

	var userProfile models.UserProfile
	if unmarshalJsonError := utils.ReadBodyAndUnmarshal(c.Request, &userProfile); unmarshalJsonError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	if err := models.SetUserProfile(id, userProfile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile"})
		return
	}

	c.JSON(http.StatusCreated, userProfile)
}

func GetUserProfile(c *gin.Context) {
	id := c.Param("user-id")
	userProfile, err := models.GetUserProfileById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func DeleteUserProfile(c *gin.Context) {
	id := c.Param("user-id")
	if err := models.DeleteUserProfile(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
