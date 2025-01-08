package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = godotenv.Load()

var (
	ctx      context.Context = context.Background()
	dbName   string          = os.Getenv("DB_NAME")
	dbUri    string          = os.Getenv("DB_URI")
	dbClient *mongo.Client   = newDbConnection()
)

func GetDbCollection(collectionName string) *mongo.Collection {
	return dbClient.Database(dbName).Collection(collectionName)
}

func newDbConnection() *mongo.Client {
	clientOptions := options.Client()
	clientOptions.ApplyURI(dbUri)
	client, connErr := mongo.Connect(ctx, clientOptions)
	if connErr != nil {
		log.Fatal("Could not connect to MongoDB")
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client
}
