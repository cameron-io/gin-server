package main

import (
	"os"

	"cameron.io/gin-server/src/middleware"
	"cameron.io/gin-server/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.TokenAuthentication())

	routes.UserRoutes(r)

	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
