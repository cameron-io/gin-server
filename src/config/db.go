package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find .env file
var _ = godotenv.Load()

// Client instance
var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	client, conn_err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(os.Getenv("DB_URI")),
	)
	if conn_err != nil {
		log.Fatal("Could not connect to MongoDB")
	}

	//ping the database
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("gopher").Collection(collectionName)
	return collection
}
