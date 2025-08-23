// Package main demonstrates basic goroutine concepts
package main

import (
	"fmt"
	"time"
)

// simpleGreeting prints a greeting after a delay
func simpleGreeting(name string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Hello, %s!\n", name)
}

// counter prints numbers up to n with a delay
func counter(id, n int) {
	for i := 1; i <= n; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Counter %d: %d\n", id, i)
	}
}

func main() {
	// Sequential execution
	fmt.Println("=== Sequential Execution ===")
	simpleGreeting("Alice")
	simpleGreeting("Bob")

	fmt.Println("\n=== Concurrent Execution ===")
	// Concurrent execution using goroutines
	go simpleGreeting("Alice")
	go simpleGreeting("Bob")

	// Multiple concurrent counters
	go counter(1, 3)
	go counter(2, 3)

	// Wait for goroutines to finish
	// In real applications, use proper synchronization
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\nAll done!")
}
