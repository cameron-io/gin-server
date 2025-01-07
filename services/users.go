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

func FindUserByEmail(ctx *gin.Context, email string) (bson.M, error) {
	filter := bson.D{
		{Key: "email", Value: email},
	}
	var result bson.M
	if err := userCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(ctx *gin.Context, new_user models.User) (*mongo.InsertOneResult, error) {
	return userCollection.InsertOne(context.TODO(), new_user)
}
