package services

import (
	"cameron.io/gin-server/config"
	"cameron.io/gin-server/db"
	"cameron.io/gin-server/entities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var profileCollection *mongo.Collection = db.GetCollection(config.MongoConnection, "profile")

var (
	UpsertMode = true
)

func GetProfileByUserId(
	c *gin.Context,
	userObjId primitive.ObjectID,
) (bson.M, error) {
	filter := bson.M{"user": userObjId}

	var result bson.M
	if err := profileCollection.FindOne(c, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func UpsertProfile(
	c *gin.Context,
	userObjId primitive.ObjectID,
	profile entities.Profile,
) (bson.M, error) {
	filter := bson.M{"user": userObjId}
	options := options.FindOneAndReplaceOptions{Upsert: &UpsertMode}

	var result bson.M
	if err := profileCollection.FindOneAndReplace(c, filter, profile, &options).Decode(&result); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}
