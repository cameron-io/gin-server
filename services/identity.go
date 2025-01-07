package services

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserIdFromClaims(c *gin.Context) (primitive.ObjectID, error) {
	user, _ := c.Get("identity")
	user_id := user.(map[string]interface{})["id"].(string)

	return primitive.ObjectIDFromHex(user_id)
}
