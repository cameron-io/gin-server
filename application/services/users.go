package services

import (
	"context"

	"cameron.io/gin-server/domain/entities"
	"cameron.io/gin-server/domain/interfaces"
	"cameron.io/gin-server/infra/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.GetDbCollection("user")

type UserService struct{}

func NewUserService() interfaces.UserService {
	return &UserService{}
}

func (s *UserService) FindUserByEmail(c *gin.Context, email string) (bson.M, error) {
	filter := bson.M{
		"email": email,
	}
	var result bson.M
	if err := userCollection.FindOne(c, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (s *UserService) CreateUser(c *gin.Context, new_user entities.User) (*mongo.InsertOneResult, error) {
	return userCollection.InsertOne(context.TODO(), new_user)
}

func (s *UserService) DeleteUserByID(c *gin.Context, userId string) (bool, error) {
	id, conv_err := primitive.ObjectIDFromHex(userId)
	if conv_err != nil {
		return false, conv_err
	}
	if err := profileCollection.FindOneAndDelete(c, bson.M{"user": id}).Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
	}
	if err := userCollection.FindOneAndDelete(c, bson.M{"_id": id}).Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
