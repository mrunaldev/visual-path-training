package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Create order processor
	processor, err := NewOrderProcessor("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to create order processor:", err)
	}
	defer processor.Close()

	// Start order processing in a goroutine
	go func() {
		if err := processor.ProcessOrders(); err != nil {
			log.Printf("Error processing orders: %v", err)
		}
	}()

	// Generate and submit sample orders
	products := []struct {
		id    string
		price float64
	}{
		{"LAPTOP", 999.99},
		{"PHONE", 599.99},
		{"TABLET", 299.99},
		{"WATCH", 199.99},
	}

	for i := 1; i <= 5; i++ {
		// Create random order
		items := make([]Item, 0)
		totalAmount := 0.0

		// Add 1-3 random items
		numItems := rand.Intn(3) + 1
		for j := 0; j < numItems; j++ {
			product := products[rand.Intn(len(products))]
			quantity := rand.Intn(3) + 1
			items = append(items, Item{
				ProductID: product.id,
				Quantity:  quantity,
				Price:     product.price,
			})
			totalAmount += product.price * float64(quantity)
		}

		order := Order{
			ID:          fmt.Sprintf("ORD-%d", i),
			CustomerID:  fmt.Sprintf("CUST-%d", rand.Intn(100)+1),
			Items:       items,
			TotalAmount: totalAmount,
		}

		// Submit order
		if err := processor.SubmitOrder(order); err != nil {
			log.Printf("Failed to submit order %s: %v", order.ID, err)
			continue
		}
		fmt.Printf("Submitted order %s for customer %s with %d items (total: $%.2f)\n",
			order.ID, order.CustomerID, len(items), totalAmount)

		// Wait a bit before next order
		time.Sleep(time.Second * 2)
	}

	// Wait for processing to complete
	fmt.Println("\nWaiting for order processing to complete...")
	time.Sleep(time.Second * 10)
}
