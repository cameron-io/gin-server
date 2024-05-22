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

// Client instance
// Find .env file
var _ = godotenv.Load()
var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	client, conn_err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(
			"mongodb://"+
				os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+
				"@"+
				os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+
				"/"+
				os.Getenv("DB_NAME")),
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
