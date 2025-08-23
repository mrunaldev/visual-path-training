package main

import (
	"fmt"

	"gorm.io/gorm"
)

// runDemo demonstrates various GORM features
func runDemo(db *gorm.DB) error {
	// Create categories
	categories := []Category{
		{Name: "Electronics", Description: "Electronic devices and gadgets"},
		{Name: "Books", Description: "Books and e-books"},
		{Name: "Clothing", Description: "Apparel and accessories"},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			return fmt.Errorf("error creating category: %v", err)
		}
	}
	fmt.Println("Created categories")

	// Create products
	products := []Product{
		{
			Name:        "Laptop",
			Description: "High-performance laptop",
			Price:       999.99,
			Stock:       10,
			Categories:  []Category{categories[0]}, // Electronics
		},
		{
			Name:        "Go Programming Book",
			Description: "Learn Go programming language",
			Price:       49.99,
			Stock:       50,
			Categories:  []Category{categories[1]}, // Books
		},
		{
			Name:        "T-Shirt",
			Description: "Cotton t-shirt",
			Price:       19.99,
			Stock:       100,
			Categories:  []Category{categories[2]}, // Clothing
		},
	}

	for _, product := range products {
		if err := db.Create(&product).Error; err != nil {
			return fmt.Errorf("error creating product: %v", err)
		}
	}
	fmt.Println("Created products")

	// Create users
	users := []User{
		{
			Name:    "John Doe",
			Email:   "john@example.com",
			Address: "123 Main St",
		},
		{
			Name:    "Jane Smith",
			Email:   "jane@example.com",
			Address: "456 Oak Ave",
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("error creating user: %v", err)
		}
	}
	fmt.Println("Created users")

	// Add items to cart
	cart := CartItem{
		UserID:    users[0].ID,
		ProductID: products[0].ID,
		Quantity:  1,
	}
	if err := db.Create(&cart).Error; err != nil {
		return fmt.Errorf("error creating cart item: %v", err)
	}
	fmt.Println("Added items to cart")

	// Create order with transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		// Check stock
		var product Product
		if err := tx.First(&product, products[0].ID).Error; err != nil {
			return err
		}
		if product.Stock < cart.Quantity {
			return fmt.Errorf("insufficient stock")
		}

		// Create order
		order := Order{
			UserID: users[0].ID,
			Status: "pending",
			Items: []OrderItem{
				{
					ProductID: products[0].ID,
					Quantity:  cart.Quantity,
					Price:     products[0].Price,
				},
			},
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// Update stock
		if err := tx.Model(&product).Update("stock", product.Stock-cart.Quantity).Error; err != nil {
			return err
		}

		// Clear cart
		if err := tx.Delete(&cart).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error processing order: %v", err)
	}
	fmt.Println("Created order")

	// Add review
	review := Review{
		UserID:    users[0].ID,
		ProductID: products[0].ID,
		Rating:    5,
		Comment:   "Great product!",
	}
	if err := db.Create(&review).Error; err != nil {
		return fmt.Errorf("error creating review: %v", err)
	}
	fmt.Println("Added review")

	// Demonstrate queries
	fmt.Println("\nDemonstrating queries:")

	// Find product with reviews and categories
	var productWithDetails Product
	err = db.Preload("Reviews.User").
		Preload("Categories").
		First(&productWithDetails, products[0].ID).Error
	if err != nil {
		return fmt.Errorf("error querying product: %v", err)
	}

	fmt.Printf("\nProduct: %s\n", productWithDetails.Name)
	fmt.Printf("Categories: %d\n", len(productWithDetails.Categories))
	fmt.Printf("Reviews: %d\n", len(productWithDetails.Reviews))

	// Find user orders with items
	var userWithOrders User
	err = db.Preload("Orders.Items.Product").
		First(&userWithOrders, users[0].ID).Error
	if err != nil {
		return fmt.Errorf("error querying user: %v", err)
	}

	fmt.Printf("\nUser: %s\n", userWithOrders.Name)
	fmt.Printf("Orders: %d\n", len(userWithOrders.Orders))

	return nil
}
