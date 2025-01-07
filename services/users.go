package services

import (
	"context"

	"cameron.io/gin-server/config"
	"cameron.io/gin-server/db"
	"cameron.io/gin-server/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CreateUser(c *gin.Context, new_user entities.User) (*mongo.InsertOneResult, error) {
	return userCollection.InsertOne(context.TODO(), new_user)
}

func DeleteUserByID(c *gin.Context, id primitive.ObjectID) (bool, error) {
	ctx := context.TODO()
	if err := profileCollection.FindOneAndDelete(ctx, bson.M{"user": id}).Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
	}
	if err := userCollection.FindOneAndDelete(ctx, bson.M{"_id": id}).Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
