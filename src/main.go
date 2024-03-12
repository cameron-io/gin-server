package main

import (
	controller "cameron.io/albums/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/albums", controller.PostAlbums)
	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumById)

	router.Run("localhost:8080")
}
