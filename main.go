package main

import (
	"os"

	"cameron.io/gin-server/middleware"
	"cameron.io/gin-server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.TokenAuthentication())

	routes.UserRoutes(r)

	r.SetTrustedProxies(nil)
	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
