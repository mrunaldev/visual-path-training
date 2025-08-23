// Package main demonstrates basic error handling patterns in Go
package main

import (
	"errors"
	"fmt"
)

// CustomError demonstrates creating custom error types
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

// divideNumbers shows basic error handling
func divideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	// Basic error handling
	result, err := divideNumbers(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %v\n", result)
	}

	// Custom error example
	customErr := &CustomError{
		Code:    500,
		Message: "internal calculation error",
	}
	fmt.Printf("Custom error: %v\n", customErr)
}
