package repositories

import (
	"cameron.io/gin-server/domain/data"
	"cameron.io/gin-server/domain/i_repositories"
	db "cameron.io/gin-server/infra/db/mongo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenRepository struct {
	collection *mongo.Collection
}

func NewGenRepository(table string) i_repositories.GenRepository {
	collection := db.GetDbCollection(table)
	return &GenRepository{collection: collection}
}

func (gr *GenRepository) Insert(c *gin.Context, entity interface{}) error {
	_, err := gr.collection.InsertOne(c, entity)
	return err
}

func (gr *GenRepository) Upsert(
	c *gin.Context,
	filter map[string]interface{},
	entity interface{}) (data.Obj, error) {
	options := options.FindOneAndReplace()
	options.SetUpsert(true)

	var result data.Obj
	err := gr.collection.FindOneAndReplace(c, filter, entity, options).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}

func (gr *GenRepository) FindAll(c *gin.Context, limit int) ([]data.Obj, error) {
	var results []data.Obj

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	cur, err := gr.collection.Find(c, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(c) {
		var doc data.Obj
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

func (gr *GenRepository) FindById(c *gin.Context, id uuid.UUID) (data.Obj, error) {
	filter := bson.M{"_id": id}
	var result data.Obj
	if err := gr.collection.FindOne(c, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (gr *GenRepository) Find(
	c *gin.Context,
	filter map[string]interface{}) (data.Obj, error) {
	var result data.Obj
	if err := gr.collection.FindOne(c, filter).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (gr *GenRepository) Delete(c *gin.Context, filter map[string]any) (bool, error) {
	if err := gr.collection.FindOneAndDelete(c, filter).Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
