package repositories

import (
	"cameron.io/gin-server/domain/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindUserByEmail(c *gin.Context, email string) (bson.M, error)
	CreateUser(c *gin.Context, new_user entities.User) (*mongo.InsertOneResult, error)
	DeleteUserByID(c *gin.Context, userId string) (bool, error)
}
