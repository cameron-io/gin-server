package db

import (
	"context"

	"cameron.io/gin-server/src/config"
	"cameron.io/gin-server/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountCollection *mongo.Collection = config.GetCollection(config.DB, "account")

func FindUserByEmail(ctx *gin.Context, email string) (bson.M, error) {
	// retrieve single and multiple documents with a specified filter using FindOne() and Find()
	// create a search filer
	filter := bson.D{
		{Key: "email", Value: email},
	}

	// retrieve all the documents that match the filter
	var result bson.M
	if err := accountCollection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func CreateAccount(ctx *gin.Context, new_account models.Account) (*mongo.InsertOneResult, error) {
	return accountCollection.InsertOne(context.TODO(), new_account)
}
