package services

import (
	"cameron.io/gin-server/config"
	"cameron.io/gin-server/db"
	"cameron.io/gin-server/models"
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

func UpsertProfile(
	c *gin.Context,
	user_obj_id primitive.ObjectID,
	profile models.Profile,
) (bson.M, error) {
	filter := bson.M{"user": user_obj_id}
	options := options.FindOneAndReplaceOptions{Upsert: &UpsertMode}

	var result bson.M
	if result := profileCollection.FindOneAndReplace(c, filter, profile, &options); result.Err() != nil {
		if result.Err() != mongo.ErrNoDocuments {
			return nil, result.Err()
		}
	}
	return result, nil
}
