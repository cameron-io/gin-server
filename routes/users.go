package routes

import (
	"cameron.io/gin-server/api"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", api.RegisterUser)
	r.POST("/login", api.AuthUser)
}
