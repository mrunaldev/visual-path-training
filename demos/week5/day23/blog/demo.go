package main

import (
	"fmt"
)

func runDemo(blog *Blog) error {
	// The blog instance is now passed as a parameter

	// Create a user
	user, err := blog.CreateUser("johndoe", "john@example.com", "secretpassword")
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	fmt.Printf("Created user: %s\n", user.Username)

	// Create some posts
	post1, err := blog.CreatePost(
		"Introduction to MongoDB",
		"MongoDB is a popular NoSQL database...",
		user.ID,
		[]string{"mongodb", "database", "nosql"},
	)
	if err != nil {
		return fmt.Errorf("failed to create post: %v", err)
	}
	fmt.Printf("Created post: %s\n", post1.Title)

	post2, err := blog.CreatePost(
		"Working with Go and MongoDB",
		"In this tutorial we'll learn how to use MongoDB with Go...",
		user.ID,
		[]string{"mongodb", "golang", "tutorial"},
	)
	if err != nil {
		return fmt.Errorf("failed to create post: %v", err)
	}
	fmt.Printf("Created post: %s\n", post2.Title)

	// Add some comments
	err = blog.AddComment(post1.ID, user.ID, "Great introduction!")
	if err != nil {
		return fmt.Errorf("failed to add comment: %v", err)
	}
	fmt.Println("Added comment to post1")

	// Simulate some views
	for i := 0; i < 5; i++ {
		if err := blog.IncrementViews(post1.ID); err != nil {
			return fmt.Errorf("failed to increment views: %v", err)
		}
	}
	for i := 0; i < 3; i++ {
		if err := blog.IncrementViews(post2.ID); err != nil {
			return fmt.Errorf("failed to increment views: %v", err)
		}
	}

	// Get posts by tag
	posts, err := blog.GetPostsByTag("mongodb")
	if err != nil {
		return fmt.Errorf("failed to get posts by tag: %v", err)
	}
	fmt.Printf("\nFound %d posts with tag 'mongodb':\n", len(posts))
	for _, p := range posts {
		fmt.Printf("- %s (views: %d)\n", p.Title, p.ViewCount)
	}

	// Get popular posts
	popular, err := blog.GetPopularPosts(2)
	if err != nil {
		return fmt.Errorf("failed to get popular posts: %v", err)
	}
	fmt.Printf("\nPopular posts:\n")
	for _, p := range popular {
		fmt.Printf("- %s (views: %d)\n", p.Title, p.ViewCount)
	}

	// Get post statistics
	stats, err := blog.GetPostStats()
	if err != nil {
		return fmt.Errorf("failed to get post stats: %v", err)
	}
	fmt.Printf("\nPost statistics by tag:\n")
	for _, stat := range stats {
		fmt.Printf("- Tag: %v, Count: %v, Avg Views: %v\n",
			stat["_id"], stat["count"], stat["avgViews"])
	}

	return nil
}
