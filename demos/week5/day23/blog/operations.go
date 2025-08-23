package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser adds a new user
func (b *Blog) CreateUser(username, email, password string) (*User, error) {
	user := User{
		Username: username,
		Email:    email,
		Password: password, // In production, hash the password
		Created:  time.Now(),
	}

	result, err := b.db.Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return &user, nil
}

// CreatePost creates a new blog post
func (b *Blog) CreatePost(title, content string, authorID primitive.ObjectID, tags []string) (*Post, error) {
	post := Post{
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		Tags:      tags,
		Comments:  []Comment{},
		Created:   time.Now(),
		Updated:   time.Now(),
		ViewCount: 0,
	}

	result, err := b.db.Collection("posts").InsertOne(context.TODO(), post)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %v", err)
	}

	post.ID = result.InsertedID.(primitive.ObjectID)
	return &post, nil
}

// AddComment adds a comment to a post
func (b *Blog) AddComment(postID, authorID primitive.ObjectID, content string) error {
	comment := Comment{
		ID:       primitive.NewObjectID(),
		Content:  content,
		AuthorID: authorID,
		Created:  time.Now(),
	}

	update := bson.D{primitive.E{Key: "$push", Value: bson.D{primitive.E{Key: "comments", Value: comment}}}}
	result, err := b.db.Collection("posts").UpdateOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: postID}},
		update,
	)

	if err != nil {
		return fmt.Errorf("failed to add comment: %v", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}

// IncrementViews increments the view count of a post
func (b *Blog) IncrementViews(postID primitive.ObjectID) error {
	update := bson.D{primitive.E{Key: "$inc", Value: bson.D{primitive.E{Key: "viewCount", Value: 1}}}}
	result, err := b.db.Collection("posts").UpdateOne(
		context.TODO(),
		bson.D{primitive.E{Key: "_id", Value: postID}},
		update,
	)

	if err != nil {
		return fmt.Errorf("failed to increment views: %v", err)
	}
	if result.ModifiedCount == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}

// GetPostsByTag finds posts with a specific tag
func (b *Blog) GetPostsByTag(tag string) ([]Post, error) {
	cursor, err := b.db.Collection("posts").Find(
		context.TODO(),
		bson.D{primitive.E{Key: "tags", Value: tag}},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %v", err)
	}
	defer cursor.Close(context.TODO())

	var posts []Post
	if err := cursor.All(context.TODO(), &posts); err != nil {
		return nil, fmt.Errorf("failed to decode posts: %v", err)
	}

	return posts, nil
}

// GetPopularPosts gets posts with most views
func (b *Blog) GetPopularPosts(limit int64) ([]Post, error) {
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "viewCount", Value: -1}}).SetLimit(limit)
	cursor, err := b.db.Collection("posts").Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %v", err)
	}
	defer cursor.Close(context.TODO())

	var posts []Post
	if err := cursor.All(context.TODO(), &posts); err != nil {
		return nil, fmt.Errorf("failed to decode posts: %v", err)
	}

	return posts, nil
}

// GetUserPosts gets all posts by a specific user
func (b *Blog) GetUserPosts(userID primitive.ObjectID) ([]Post, error) {
	cursor, err := b.db.Collection("posts").Find(
		context.TODO(),
		bson.D{primitive.E{Key: "authorId", Value: userID}},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %v", err)
	}
	defer cursor.Close(context.TODO())

	var posts []Post
	if err := cursor.All(context.TODO(), &posts); err != nil {
		return nil, fmt.Errorf("failed to decode posts: %v", err)
	}

	return posts, nil
}

// GetPostStats gets statistics about posts using aggregation
func (b *Blog) GetPostStats() ([]bson.M, error) {
	pipeline := mongo.Pipeline{
		bson.D{primitive.E{Key: "$unwind", Value: "$tags"}},
		bson.D{primitive.E{Key: "$group", Value: bson.D{
			primitive.E{Key: "_id", Value: "$tags"},
			primitive.E{Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}},
			primitive.E{Key: "avgViews", Value: bson.D{primitive.E{Key: "$avg", Value: "$viewCount"}}},
		}}},
		bson.D{primitive.E{Key: "$sort", Value: bson.D{primitive.E{Key: "count", Value: -1}}}},
	}

	cursor, err := b.db.Collection("posts").Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, fmt.Errorf("failed to aggregate posts: %v", err)
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, fmt.Errorf("failed to decode results: %v", err)
	}

	return results, nil
}
