package services

import (
	"context"

	"cameron.io/gin-server/config"
	"cameron.io/gin-server/db"
	"cameron.io/gin-server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.GetCollection(config.MongoConnection, "user")

func FindUserByEmail(c *gin.Context, email string) (bson.M, error) {
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

func CreateUser(c *gin.Context, new_user models.User) (*mongo.InsertOneResult, error) {
	return userCollection.InsertOne(context.TODO(), new_user)
}

func DeleteUserByEmail(c *gin.Context, email string) (bool, error) {
	filter := bson.M{
		"email": email,
	}
	res := userCollection.FindOneAndDelete(context.TODO(), filter)
	if err := res.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
