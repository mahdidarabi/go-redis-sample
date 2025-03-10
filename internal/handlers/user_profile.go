package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-redis-sample/internal/models"
	"go-redis-sample/internal/utils"
)

// GetUserProfile godoc
// @Summary Get a user profile
// @Description Get a user profile from Redis
// @Tags user-profile
// @Produce json
// @Param user-id path string true "User ID"
// @Success 200 {object} models.UserProfile
// @Failure 404 {object} object "Not Found"
// @Router /user-profile/{user-id} [get]
func GetUserProfile(c *gin.Context) {
	id := c.Param("user-id")
	userProfile, err := models.GetUserProfileById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

// CreateUserProfile godoc
// @Summary Create a user profile
// @Description Create a user profile in Redis
// @Tags user-profile
// @Accept json
// @Produce json
// @Param user-id path string true "User ID"
// @Param userProfile body models.UserProfile true "User Profile"
// @Success 201 {object} models.UserProfile
// @Failure 500 {object} object "Internal Server Error"
// @Router /user-profile/{user-id} [post]
func CreateUserProfile(c *gin.Context) {
	handleCreateOrUpdateUserProfile(c, "CREATE")
}

// UpdateUserProfile godoc
// @Summary Update a user profile
// @Description Update a user profile in Redis
// @Tags user-profile
// @Accept json
// @Produce json
// @Param user-id path string true "User ID"
// @Param userProfile body models.UserProfile true "User Profile"
// @Success 200 {object} models.UserProfile
// @Failure 500 {object} object "Internal Server Error"
// @Router /user-profile/{user-id} [put]
func UpdateUserProfile(c *gin.Context) {
	handleCreateOrUpdateUserProfile(c, "UPDATE")
}

// DeleteUserProfile godoc
// @Summary Delete a user profile
// @Description Delete a user profile from Redis
// @Tags user-profile
// @Param user-id path string true "User ID"
// @Success 204
// @Failure 404 {object} object "Not Found"
// @Router /user-profile/{user-id} [delete]
func DeleteUserProfile(c *gin.Context) {
	id := c.Param("user-id")
	if err := models.DeleteUserProfile(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func handleCreateOrUpdateUserProfile(c *gin.Context, mode string) {
	id := c.Param("user-id")

	userProfileExisting, err := models.GetUserProfileById(id)
	if err != nil && err.Error() != "redis: nil" { // Check if the error is not due to a non-existing profile
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if user profile exists"})
		return
	}

	responseCode := 0

	if mode == "CREATE" {
		responseCode = http.StatusCreated
		if userProfileExisting != (models.UserProfile{}) {
			responseCode = http.StatusConflict
			c.JSON(responseCode, gin.H{"error": "User profile already exists"})
			return
		}
	} else if mode == "UPDATE" {
		responseCode = http.StatusOK
		if userProfileExisting == (models.UserProfile{}) {
			responseCode = http.StatusNotFound
			c.JSON(responseCode, gin.H{"error": "User profile not found"})
			return
		}
	}

	var userProfile models.UserProfile
	if unmarshalJsonError := utils.ReadBodyAndUnmarshal(c.Request, &userProfile); unmarshalJsonError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	if err := models.SetUserProfile(id, userProfile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mutate user profile"})
		return
	}
	c.JSON(responseCode, userProfile)
}
