package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Blog represents our blog application
type Blog struct {
	db *mongo.Database
}

// User represents a blog user
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Created  time.Time          `bson:"created"`
}

// Post represents a blog post
type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	AuthorID  primitive.ObjectID `bson:"authorId"`
	Tags      []string           `bson:"tags"`
	Comments  []Comment          `bson:"comments"`
	Created   time.Time          `bson:"created"`
	Updated   time.Time          `bson:"updated"`
	ViewCount int                `bson:"viewCount"`
}

// Comment represents a comment on a blog post
type Comment struct {
	ID       primitive.ObjectID `bson:"_id"`
	Content  string             `bson:"content"`
	AuthorID primitive.ObjectID `bson:"authorId"`
	Created  time.Time          `bson:"created"`
}

// createIndexes sets up the required indexes for our collections
func (b *Blog) createIndexes() error {
	// Create username index for users collection
	userIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{primitive.E{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{primitive.E{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := b.db.Collection("users").Indexes().CreateMany(context.TODO(), userIndexes)
	if err != nil {
		return err
	}

	// Create indexes for posts collection
	postIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{primitive.E{Key: "authorId", Value: 1}},
		},
		{
			Keys: bson.D{primitive.E{Key: "tags", Value: 1}},
		},
		{
			Keys: bson.D{primitive.E{Key: "viewCount", Value: -1}},
		},
	}
	_, err = b.db.Collection("posts").Indexes().CreateMany(context.TODO(), postIndexes)
	if err != nil {
		return err
	}

	return nil
}
