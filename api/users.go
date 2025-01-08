package api

import (
	"net/http"
	"time"

	"cameron.io/gin-server/domain/entities"
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
	var newUser entities.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the User struct
	if err := validator.New().Struct(newUser); err != nil {
		// Validation failed, handle the error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	existingUser, err := services.FindUserByEmail(c, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return
	}

	// Hash password
	if newUser.Password, err = utils.HashPassword(newUser.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"unexpected_error": err.Error()})
		return
	}

	// Create new user
	newUser.CreatedAt = time.Now().UnixMilli()

	if _, err := services.CreateUser(c, newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func DeleteUser(c *gin.Context) {
	userId := services.GetUserIdFromClaims(c)

	res, err := services.DeleteUserByID(c, userId)
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
