package main

import (
	"os"

	"cameron.io/gin-server/api"
	"cameron.io/gin-server/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.TokenAuthentication())

	api.UserRoutes(r)

	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
