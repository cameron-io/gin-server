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
		created_user, err := services.CreateUser(ctx, new_user)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, created_user)
	})

	r.POST("/api/accounts/login", func(ctx *gin.Context) {
		var user_auth models.Auth
		if err := ctx.ShouldBindJSON(&user_auth); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := validator.New().Struct(user_auth); err != nil {
			// Validation failed, handle the error
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		existing_user, err := services.FindUserByEmail(ctx, user_auth.Email)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
			return
		}
		if existing_user == nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"msg": "invalid credentials."})
			return
		}
		is_match := utils.CheckPasswordHash(user_auth.Password)
		if !is_match {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "invalid credentials."})
			return
		}

		token, token_err := utils.CreateToken(user_auth.Email)
		if token_err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"token_error": token_err.Error()})
			return
		}
		ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)
		ctx.Status(http.StatusOK)
	})
}
