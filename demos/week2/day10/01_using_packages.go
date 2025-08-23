// Package main demonstrates using the calculator package
package main

import (
	"fmt"
	"log"

	"calculator"
)

func main() {
	// Using the calculator package
	a, b := 10.0, 5.0

	sum := calculator.Add(a, b)
	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, sum)

	diff := calculator.Subtract(a, b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, diff)

	product := calculator.Multiply(a, b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, product)

	quotient := calculator.Divide(a, b)
	fmt.Printf("%.2f / %.2f = %.2f\n", a, b, quotient)

	// Test division by zero
	result := calculator.Divide(a, 0)
	log.Printf("Division by zero returns: %.2f\n", result)
}
