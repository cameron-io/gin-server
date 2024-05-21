package main

import (
	"context"
	"log"
	"os"

	"cameron.io/gin-server/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	if _, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(
			"mongodb://"+
				os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+
				"@"+
				os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+
				"/"+
				os.Getenv("DB_NAME")),
	); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

func main() {
	r := gin.Default()

	routes.AccountRoutes(r)

	r.Run("localhost:" + os.Getenv("SERVER_PORT"))
}
