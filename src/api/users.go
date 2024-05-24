package api

import (
	"net/http"
	"time"

	"cameron.io/gin-server/src/auth"
	"cameron.io/gin-server/src/db"
	"cameron.io/gin-server/src/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RegisterUser(ctx *gin.Context) {
	var new_user models.User
	if err := ctx.ShouldBindJSON(&new_user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the User struct
	if err := validator.New().Struct(new_user); err != nil {
		// Validation failed, handle the error
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing_user, err := db.FindUserByEmail(ctx, new_user.Email)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return
	}
	if existing_user != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return
	}

	if new_user.Password, err = auth.HashPassword(new_user.Password); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"unexpected_error": err.Error()})
		return
	}
	new_user.CreatedAt = time.Now().UnixMilli()
	created_user, err := db.CreateUser(ctx, new_user)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, created_user)
}
