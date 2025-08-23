// Package main demonstrates recover functionality in Go
package main

import (
	"fmt"
	"log"
)

// recoverFromPanic demonstrates how to use recover
func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Printf("Recovered from panic: %v\n", r)
	}
}

// processItem demonstrates panic recovery in a real scenario
func processItem(item interface{}) {
	defer recoverFromPanic()

	// Type assertion that might panic
	value := item.(string)
	fmt.Printf("Processing string: %s\n", value)
}

func main() {
	// Example 1: Recovering from a type assertion panic
	log.Println("Processing string item...")
	processItem("hello")

	log.Println("Processing integer item...")
	processItem(42) // This will panic but recover

	log.Println("Program continues normally")
}
