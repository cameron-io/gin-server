package data

import (
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Obj map[string]any

func ConvToUuid(str string) (primitive.ObjectID, error) {
	switch os.Getenv("DB_ENGINE") {
	case "mongodb":
		return primitive.ObjectIDFromHex(str)
	default:
		return [12]byte{}, errors.New("unknown db type")
	}
}
