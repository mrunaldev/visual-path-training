package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Create mock DB
	db := NewMockDB()

	// Create catalog service
	catalog, err := NewCatalogService(db, "localhost:6379")
	if err != nil {
		log.Fatal("Failed to create catalog service:", err)
	}
	defer catalog.Close()

	// Add some sample products
	products := []*Product{
		{
			ID:          "p1",
			Name:        "Laptop",
			Description: "High-performance laptop",
			Price:       999.99,
			Category:    "electronics",
		},
		{
			ID:          "p2",
			Name:        "Smartphone",
			Description: "Latest smartphone",
			Price:       699.99,
			Category:    "electronics",
		},
		{
			ID:          "p3",
			Name:        "T-Shirt",
			Description: "Cotton t-shirt",
			Price:       19.99,
			Category:    "clothing",
		},
	}

	// Save products
	for _, p := range products {
		if err := catalog.SaveProduct(p); err != nil {
			log.Printf("Failed to save product %s: %v\n", p.ID, err)
			continue
		}
		fmt.Printf("Saved product: %s\n", p.Name)
	}

	// Demonstrate cache-aside pattern
	fmt.Println("\nDemonstrating cache-aside pattern:")
	for i := 0; i < 2; i++ {
		product, err := catalog.GetProduct("p1")
		if err != nil {
			log.Printf("Failed to get product: %v\n", err)
			continue
		}
		fmt.Printf("Retrieved product: %s (should see cache hit on second attempt)\n", product.Name)
	}

	// Demonstrate category retrieval
	fmt.Println("\nDemonstrating category retrieval:")
	for i := 0; i < 2; i++ {
		products, err := catalog.GetProductsByCategory("electronics")
		if err != nil {
			log.Printf("Failed to get products: %v\n", err)
			continue
		}
		fmt.Printf("Found %d electronics products (should see cache hit on second attempt)\n", len(products))
		for _, p := range products {
			fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
		}
	}

	// Demonstrate cache expiration
	fmt.Println("\nDemonstrating cache expiration (waiting 2 seconds)...")
	time.Sleep(2 * time.Second)

	product, err := catalog.GetProduct("p1")
	if err != nil {
		log.Printf("Failed to get product: %v\n", err)
	} else {
		fmt.Printf("Retrieved product after delay: %s\n", product.Name)
	}
}
