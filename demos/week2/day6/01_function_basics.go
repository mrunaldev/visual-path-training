// Package main demonstrates basic function concepts in Go
package main

import (
	"fmt"
)

// greet is a simple function that takes a name parameter
// and returns a greeting string
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// add demonstrates a function with multiple parameters
func add(a, b int) int {
	return a + b
}

// swap demonstrates multiple return values
func swap(x, y string) (string, string) {
	return y, x
}

// sum demonstrates variadic functions
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Basic function call
	message := greet("Gopher")
	fmt.Println(message)

	// Function with multiple parameters
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// Function with multiple return values
	first, second := swap("hello", "world")
	fmt.Printf("Swapped: %s %s\n", first, second)

	// Variadic function calls
	fmt.Println("Sum:", sum(1, 2, 3))
	fmt.Println("Sum:", sum(10, 20, 30, 40))
}
