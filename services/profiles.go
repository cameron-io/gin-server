package services

import (
	"cameron.io/gin-server/config"
	"cameron.io/gin-server/db"
	"cameron.io/gin-server/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var profileCollection *mongo.Collection = db.GetCollection(config.MongoConnection, "profile")

var (
	UpsertMode = true
)

func UpsertProfile(
	c *gin.Context,
	user_id string,
	profileFields models.Profile,
) (bson.M, error) {
	filter := bson.M{"user": user_id}
	update := bson.M{"$set": profileFields}
	options := options.FindOneAndUpdateOptions{Upsert: &UpsertMode}

	var result bson.M
	if err := profileCollection.FindOneAndUpdate(c, filter, update, &options).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
