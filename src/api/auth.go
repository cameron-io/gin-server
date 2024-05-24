package api

import (
	"net/http"

	"cameron.io/gin-server/src/db"
	"cameron.io/gin-server/src/models"
	"cameron.io/gin-server/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AuthUser(ctx *gin.Context) {
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
	existing_user, err := db.FindUserByEmail(ctx, user_auth.Email)
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
}
