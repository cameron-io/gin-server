package db

import (
	"context"

	"cameron.io/gin-server/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a global variable to hold our MongoDB connection
var mongoClient *mongo.Client
var collection = mongoClient.Database("gopher").Collection("account")

func FindUserByEmail(ctx *gin.Context, email string) bson.M {
	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// create a search filer
	filter := bson.D{
		{Key: "email", Value: email},
	}

	// retrieve all the documents that match the filter
	var result bson.M
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		panic(err)
	}

	return result
}

func CreateAccount(ctx *gin.Context, new_account models.Account) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), new_account)
}
