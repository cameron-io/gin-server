package routes

import (
	"net/http"

	"cameron.io/gin-server/api"
	"cameron.io/gin-server/models"
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

		created_user, err := api.UserCreate(ctx, new_user)

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

		token, err := api.UserAuth(ctx, user_auth)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"token_error": err.Error()})
			return
		}

		ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)
		ctx.Status(http.StatusOK)
	})
}
