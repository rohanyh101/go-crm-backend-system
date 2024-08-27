package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBInstance initializes a new MongoDB client instance
func DBInstance() *mongo.Client {
	// Load environment variables from .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Printf("error loading .env file: %v", err)
	// }

	// Get MongoDB URI from environment variables
	MONGO_URI := os.Getenv("MONGO_URI")
	if MONGO_URI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	// Create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("error connecting to MongoDB: %v", err)
	}

	// Ping MongoDB to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

// OpenCollection opens a collection from the specified database
func OpenCollection(databaseName, collectionName string) *mongo.Collection {
	client := DBInstance()

	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}
