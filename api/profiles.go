package api

import (
	"net/http"

	"cameron.io/gin-server/entities"
	"cameron.io/gin-server/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func UpsertProfile(c *gin.Context) {
	var new_profile entities.Profile

	if err := c.ShouldBindJSON(&new_profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.New().Struct(new_profile); err != nil {
		// Validation failed, handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, conv_err := services.GetUserIdFromClaims(c)
	if conv_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"conv_error": conv_err.Error()})
		return
	}
	new_profile.User = user_id

	profile, err := services.UpsertProfile(c, user_id, new_profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_upsert_error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}
