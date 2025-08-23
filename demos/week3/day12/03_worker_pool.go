// Package main demonstrates channel communication patterns
package main

import (
	"fmt"
	"sync"
)

// WorkPool demonstrates a worker pool pattern
type WorkPool struct {
	jobs    chan int
	results chan int
	wg      sync.WaitGroup
}

// NewWorkPool creates a new work pool with the specified number of workers
func NewWorkPool(numWorkers int) *WorkPool {
	wp := &WorkPool{
		jobs:    make(chan int, numWorkers),
		results: make(chan int, numWorkers),
	}

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}

	return wp
}

// worker processes jobs
func (wp *WorkPool) worker(id int) {
	defer wp.wg.Done()

	for job := range wp.jobs {
		// Simulate processing
		result := job * 2
		fmt.Printf("Worker %d processed job %d -> %d\n", id, job, result)
		wp.results <- result
	}
}

// Close shuts down the work pool
func (wp *WorkPool) Close() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

func main() {
	// Create a work pool with 3 workers
	pool := NewWorkPool(3)

	// Send jobs
	go func() {
		for i := 1; i <= 6; i++ {
			pool.jobs <- i
		}
		pool.Close()
	}()

	// Collect results
	var results []int
	for result := range pool.results {
		results = append(results, result)
	}

	fmt.Println("\nAll results:", results)
}
