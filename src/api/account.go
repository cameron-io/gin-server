package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAccount(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusCreated, ctx)
}

func GetAccount(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, ctx)
}

func GetAccountById(ctx *gin.Context) {
	// id := ctx.Param("id")
	ctx.IndentedJSON(http.StatusOK, ctx)
}
