package controller

import (
	"net/http"

	"cameron.io/albums/data"
	"cameron.io/albums/src/models"
	"github.com/gin-gonic/gin"
)

func PostAlbums(ctx *gin.Context) {
	var newAlbum models.Album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, data.GetAlbums())
}

func GetAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range data.GetAlbums() {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
