// Package main demonstrates advanced function concepts in Go
package main

import (
	"fmt"
)

// init function is called before main
func init() {
	fmt.Println("Initializing...")
}

// closure returns an anonymous function
func closure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// multipleReturns demonstrates returning multiple values
func multipleReturns(x int) (int, error) {
	if x < 0 {
		return 0, fmt.Errorf("negative number not allowed")
	}
	return x * 2, nil
}

// deferDemo shows defer keyword usage
func deferDemo() {
	defer fmt.Println("This prints last")
	fmt.Println("This prints first")
}

func main() {
	// Anonymous function example
	func() {
		fmt.Println("Anonymous function called")
	}()

	// Closure example
	counter := closure()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2

	// Multiple returns with error handling
	if result, err := multipleReturns(-5); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Defer example
	deferDemo()
}
