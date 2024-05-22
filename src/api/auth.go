package api

import (
	"net/http"

	"cameron.io/gin-server/src/db"
	"cameron.io/gin-server/src/models"
	"cameron.io/gin-server/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Authenticate(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := validator.New().Struct(user); err != nil {
		// Validation failed, handle the error
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existing_user, err := db.FindUserByEmail(ctx, user.Email)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return
	}
	if existing_user == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"msg": "invalid credentials."})
		return
	}
	is_match := utils.CheckPasswordHash(user.Password)
	if !is_match {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "invalid credentials."})
		return
	}

	ctx.Status(http.StatusOK)

	// TODO: Return HttpOnly JWT Cookie
}
