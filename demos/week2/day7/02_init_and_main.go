// Package main demonstrates the init and main function relationship
package main

import (
	"fmt"
)

var globalVar = initGlobalVar()

func initGlobalVar() int {
	fmt.Println("Initializing global variable")
	return 42
}

func init() {
	fmt.Println("First init function")
}

func init() {
	fmt.Println("Second init function")
}

func main() {
	fmt.Println("Main function - Global var:", globalVar)

	// Underscore operator example
	nums := []int{1, 2, 3, 4}
	for _, num := range nums {
		fmt.Println("Number:", num)
	}
}
