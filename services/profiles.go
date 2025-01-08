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

func GetAllProfiles(c *gin.Context) ([]bson.M, error) {
	var results []bson.M

	findOptions := options.Find()
	findOptions.SetLimit(10)

	cur, err := profileCollection.Find(c, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(c) {
		var doc bson.M
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		results = append(results, doc)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func UpsertProfile(
	c *gin.Context,
	userObjId primitive.ObjectID,
	profile entities.Profile,
) (bson.M, error) {
	filter := bson.M{"user": userObjId}
	options := options.FindOneAndReplace()
	options.SetUpsert(true)

	var result bson.M
	if err := profileCollection.FindOneAndReplace(c, filter, profile, options).Decode(&result); err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}
