package api

import (
	"net/http"

	"cameron.io/gin-server/entities"
	"cameron.io/gin-server/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func GetCurrentUserProfile(c *gin.Context) {
	userId, claimErr := services.GetUserIdFromClaims(c)
	if claimErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": claimErr.Error()})
		return
	}
	profile, dbErr := services.GetProfileByUserId(c, userId)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
	}

	c.JSON(http.StatusOK, profile)
}

func UpsertProfile(c *gin.Context) {
	var newProfile entities.Profile

	if err := c.ShouldBindJSON(&newProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.New().Struct(newProfile); err != nil {
		// Validation failed, handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, convErr := services.GetUserIdFromClaims(c)
	if convErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"conv_error": convErr.Error()})
		return
	}
	newProfile.User = userId

	profile, err := services.UpsertProfile(c, userId, newProfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_upsert_error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

func GetAllProfiles(c *gin.Context) {
	profiles, err := services.GetAllProfiles(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_all_profiles_error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profiles)
}
