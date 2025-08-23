package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewBlog creates a new blog instance
func NewBlog(uri string) (*Blog, error) {
	// Create client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping database
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	// Get database
	db := client.Database("blog_demo")

	// Create indexes
	if err := createIndexes(db); err != nil {
		return nil, fmt.Errorf("failed to create indexes: %v", err)
	}

	return &Blog{
		db: db,
	}, nil
}

// createIndexes creates necessary indexes
func createIndexes(db *mongo.Database) error {
	// Create unique index on username
	userModel := mongo.IndexModel{
		Keys:    bson.D{primitive.E{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := db.Collection("users").Indexes().CreateOne(context.TODO(), userModel)
	if err != nil {
		return fmt.Errorf("failed to create username index: %v", err)
	}

	// Create index on tags
	postModel := mongo.IndexModel{
		Keys: bson.D{primitive.E{Key: "tags", Value: 1}},
	}
	_, err = db.Collection("posts").Indexes().CreateOne(context.TODO(), postModel)
	if err != nil {
		return fmt.Errorf("failed to create tags index: %v", err)
	}

	return nil
}

func main() {
	// Create blog instance
	blog, err := NewBlog("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// Run demo
	if err := runDemo(blog); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Demo completed successfully!")
}
