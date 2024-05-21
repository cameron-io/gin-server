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

	if db.FindUserByEmail(ctx, new_account.Email) != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "user already exists"})
		return
	}

	result, err := db.CreateAccount(ctx, new_account)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, result)
}

func GetAccount(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, ctx)
}

func GetAccountById(ctx *gin.Context) {
	// id := ctx.Param("id")
	ctx.IndentedJSON(http.StatusOK, ctx)
}
