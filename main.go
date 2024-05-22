package main

import (
	"os"

	"cameron.io/gin-server/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoutes(r)

	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
