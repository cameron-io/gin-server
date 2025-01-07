package api

import (
	"net/http"

	"cameron.io/gin-server/models"
	"cameron.io/gin-server/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpsertProfile(c *gin.Context) {
	var new_profile models.Profile

	if err := c.ShouldBindJSON(&new_profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.New().Struct(new_profile); err != nil {
		// Validation failed, handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("identity")
	user_id := user.(map[string]interface{})["id"].(string)

	user_obj_id, conv_err := primitive.ObjectIDFromHex(user_id)
	if conv_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"conv_error": conv_err.Error()})
		return
	}
	new_profile.User = user_obj_id

	profile, err := services.UpsertProfile(c, user_obj_id, new_profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_upsert_error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, profile)
}
