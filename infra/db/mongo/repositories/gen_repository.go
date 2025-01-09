package repositories

import (
	"cameron.io/gin-server/domain/repositories"
	"cameron.io/gin-server/infra/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GenRepository struct {
	collection *mongo.Collection
}

func NewGenRepository(collection mongo.Collection) repositories.GenRepository {
	return &GenRepository{collection: &collection}
}

func (gr *GenRepository) Insert(c *gin.Context, entity db.Obj) error {
	_, err := gr.collection.InsertOne(c, entity)
	return err
}

func (gr *GenRepository) Upsert(
	c *gin.Context,
	filter map[string]interface{},
	entity db.Obj) (db.Obj, error) {
	options := options.FindOneAndReplace()
	options.SetUpsert(true)

	var result db.Obj
	err := gr.collection.FindOneAndReplace(c, filter, entity, options).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
	}
	return result, nil
}

func (gr *GenRepository) FindAll(c *gin.Context, limit int) ([]db.Obj, error) {
	var results []db.Obj

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	cur, err := gr.collection.Find(c, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	for cur.Next(c) {
		var doc db.Obj
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

func (gr *GenRepository) FindById(c *gin.Context, id uuid.UUID) (db.Obj, error) {
	filter := bson.M{"_id": id}
	var result db.Obj
	if err := gr.collection.FindOne(c, filter).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (gr *GenRepository) Find(
	c *gin.Context,
	filter map[string]interface{}) (db.Obj, error) {
	var result db.Obj
	if err := gr.collection.FindOne(c, filter).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (gr *GenRepository) Delete(c *gin.Context, id uuid.UUID) (bool, error) {
	filter := bson.M{"_id": id}
	if err := gr.collection.FindOneAndDelete(c, filter).Err(); err != nil {
		if err != mongo.ErrNoDocuments {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
