package db

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUserByEmail(ctx *gin.Context, client *mongo.Client, email string) bson.M {
	collection := getCollection(client)

	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// create a search filer
	filter := bson.D{
		{
			Key:   "email",
			Value: email,
		},
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

func getCollection(client *mongo.Client) *mongo.Collection {
	return client.Database("gopher").Collection("account")
}
