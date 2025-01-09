package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type Obj map[string]any

func ConvToUuid(str string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(str)
}
