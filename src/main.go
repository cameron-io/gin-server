package main

import (
	"context"
	"log"

	"cameron.io/gin-server/src/api"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This function runs before we call our main function and connects to our MongoDB database. If it cannot connect, the application stops.
func init() {
	if _, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

func main() {
	r := gin.Default()

	r.POST("/accounts", api.PostAccount)

	r.Run("localhost:5000")
}
