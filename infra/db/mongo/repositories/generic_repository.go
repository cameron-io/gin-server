package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenericRepository struct {
	collection *mongo.Collection
}

func (gr *GenericRepository) Insert(c *gin.Context, entity interface{}) error {
	_, err := gr.collection.InsertOne(c, entity.(bson.M))
	return err
}

func (gr *GenericRepository) Upsert(
	c *gin.Context,
	filter interface{},
	entity interface{}) (interface{}, error) {
	bsonfilter := filter.(bson.M)
	bsonEntity := entity.(bson.M)

	options := options.FindOneAndReplace()
	options.SetUpsert(true)

	var result bson.M
	err := gr.collection.FindOneAndReplace(c, bsonfilter, bsonEntity, options).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}

func (gr *GenericRepository) FindAll(c *gin.Context, limit int) (interface{}, error) {
	var results []bson.M

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	cur, err := gr.collection.Find(c, bson.D{}, findOptions)
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

func (gr *GenericRepository) Find(c *gin.Context, filter any) (interface{}, error) {
	bsonFilter := filter.(bson.M)
	var result bson.M
	if err := gr.collection.FindOne(c, bsonFilter).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (gr *GenericRepository) Delete(c *gin.Context, uuid uuid.UUID) (bool, error) {
	filter := bson.M{"_id": uuid}
	if err := gr.collection.FindOneAndDelete(c, filter).Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
