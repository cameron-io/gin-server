package services

import (
	"cameron.io/gin-server/application/interfaces"
	"cameron.io/gin-server/domain/entities"
	db "cameron.io/gin-server/infra/db/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Move to repository, add as dependency
var profileCollection *mongo.Collection = db.GetDbCollection("profile")

type ProfileService struct{}

func NewProfileService() interfaces.ProfileService {
	return &ProfileService{}
}

func (s *ProfileService) GetProfileByUserId(
	c *gin.Context,
	userId string,
) (bson.M, error) {
	id, conv_err := primitive.ObjectIDFromHex(userId)
	if conv_err != nil {
		return nil, conv_err
	}
	filter := bson.M{"user": id}

	var result bson.M
	if err := profileCollection.FindOne(c, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *ProfileService) GetAllProfiles(c *gin.Context) ([]bson.M, error) {
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

func (s *ProfileService) UpsertProfile(
	c *gin.Context,
	userId string,
	profile entities.Profile,
) (bson.M, error) {
	id, conv_err := primitive.ObjectIDFromHex(userId)
	if conv_err != nil {
		return nil, conv_err
	}

	profile.User = id

	filter := bson.M{"user": id}
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
