package routes

import (
	"cameron.io/gin-server/src/api"
	"github.com/gin-gonic/gin"
)

func AccountRoutes(r *gin.Engine) {
	r.POST("/accounts", api.PostAccount)
	r.GET("/accounts", api.GetAccount)
	r.POST("/accounts/:id", api.PostAccount)
}
