package services

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserIdFromClaims(c *gin.Context) (primitive.ObjectID, error) {
	user, _ := c.Get("identity")
	userId := user.(map[string]interface{})["id"].(string)

	return primitive.ObjectIDFromHex(userId)
}
