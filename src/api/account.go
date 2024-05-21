package api

import (
	"net/http"

	"cameron.io/gin-server/src/db"
	"cameron.io/gin-server/src/models"
	"github.com/gin-gonic/gin"
)

func PostAccount(ctx *gin.Context) {
	var new_account models.Account
	if err := ctx.ShouldBindJSON(&new_account); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing_user, err := db.FindUserByEmail(ctx, new_account.Email)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_find_error": err.Error()})
		return
	}
	if existing_user != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return
	}

	created_user, err := db.CreateAccount(ctx, new_account)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"db_create_error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, created_user)
}
