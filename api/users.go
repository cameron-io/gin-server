package api

import (
	"net/http"
	"time"

	"cameron.io/gin-server/models"
	"cameron.io/gin-server/services"
	"cameron.io/gin-server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func GetUserInfo(c *gin.Context) {
	user, _ := c.Get("identity")
	c.JSON(http.StatusOK, user)
}

func RegisterUser(c *gin.Context) {
	var new_user models.User

	if err := c.ShouldBindJSON(&new_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the User struct
	if err := validator.New().Struct(new_user); err != nil {
		// Validation failed, handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	existing_user, err := services.FindUserByEmail(c, new_user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return
	}
	if existing_user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return
	}

	// Hash password
	if new_user.Password, err = utils.HashPassword(new_user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"unexpected_error": err.Error()})
		return
	}

	// Create new user
	new_user.CreatedAt = time.Now().UnixMilli()

	if _, err := services.CreateUser(c, new_user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteUser(c *gin.Context) {
	user_id, claims_err := services.GetUserIdFromClaims(c)
	if claims_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": claims_err.Error()})
		return
	}

	res, err := services.DeleteUserByID(c, user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !res {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}
