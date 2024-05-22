package routes

import (
	"cameron.io/gin-server/src/api"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", api.RegisterUser)
}
