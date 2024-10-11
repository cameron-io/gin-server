package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = godotenv.Load()

// Client instance
var MongoConnection *mongo.Client = connectDB()

func connectDB() *mongo.Client {
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
	log.Println("Connected to MongoDB")
	return client
}
