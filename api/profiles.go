package api

import (
	"net/http"

	"cameron.io/gin-server/models"
	"cameron.io/gin-server/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
	user_id := user.(map[string]interface{})["_id"].(string)
	if _, err := services.UpsertProfile(c, user_id, new_profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_upsert_error": err.Error()})
		return
	}

	c.Status(201)
}
