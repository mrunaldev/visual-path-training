// Package main demonstrates panic handling in Go
package main

import (
	"fmt"
)

// dangerousOperation demonstrates a function that might panic
func dangerousOperation(shouldPanic bool) {
	if shouldPanic {
		panic("something went terribly wrong!")
	}
	fmt.Println("Operation completed successfully")
}

// Example of array out of bounds panic
func accessArray(index int) {
	arr := []int{1, 2, 3}
	fmt.Printf("Value at index %d: %d\n", index, arr[index])
}

func main() {
	// Safe operation
	fmt.Println("Executing safe operation...")
	dangerousOperation(false)

	// Panic example
	fmt.Println("\nExecuting dangerous operation...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	dangerousOperation(true)
	fmt.Println("This line won't be executed")
}
