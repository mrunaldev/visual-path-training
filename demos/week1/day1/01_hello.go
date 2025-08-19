// Package main demonstrates a basic Go program structure
package main

import (
	"fmt"
	"log"
	"os"
)

// main is the entry point of the program.
// It demonstrates basic output and error handling.
func main() {
	// Print a greeting message
	message := "Hello, Go!"
	
	// Write to stdout, checking for errors
	if _, err := fmt.Println(message); err != nil {
		log.Printf("Error printing message: %v\n", err)
		os.Exit(1)
	}
}
