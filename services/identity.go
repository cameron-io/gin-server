package services

import (
	"github.com/gin-gonic/gin"
)

func GetUserIdFromClaims(c *gin.Context) string {
	user, _ := c.Get("identity")
	return user.(map[string]interface{})["id"].(string)
}
