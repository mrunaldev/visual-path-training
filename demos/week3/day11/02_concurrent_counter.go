// Package main demonstrates concurrent execution with shared data
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is a thread-safe counter
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment adds 1 to the counter safely
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count safely
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// worker increments the counter multiple times
func worker(id int, counter *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		counter.Increment()
		fmt.Printf("Worker %d: counter = %d\n", id, counter.Value())
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Create a safe counter
	counter := &SafeCounter{}

	// Create a WaitGroup to wait for all goroutines
	var wg sync.WaitGroup

	// Launch multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, counter, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Printf("\nFinal counter value: %d\n", counter.Value())
}
