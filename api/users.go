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

func UserRoutes(r *gin.Engine) {
	r.POST("/api/accounts/register", func(ctx *gin.Context) {
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

		// Check if user already exists
		existing_user, err := services.FindUserByEmail(ctx, new_user.Email)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
			return
		}
		if existing_user != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
			return
		}

		// Hash password
		if new_user.Password, err = utils.HashPassword(new_user.Password); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"unexpected_error": err.Error()})
			return
		}

		// Create new user
		new_user.CreatedAt = time.Now().UnixMilli()

		if _, err := services.CreateUser(ctx, new_user); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
			return
		}

		created_user, err := services.FindUserByEmail(ctx, new_user.Email)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusCreated, created_user)
	})
}
