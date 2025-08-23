// Package main demonstrates practical concurrent processing
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID       int
	Duration time.Duration
	Result   string
}

// processTask simulates processing a task
func processTask(task Task) string {
	time.Sleep(task.Duration)
	return fmt.Sprintf("Task %d completed", task.ID)
}

// taskProcessor processes tasks concurrently
func taskProcessor(tasks []Task) []string {
	var wg sync.WaitGroup
	results := make([]string, len(tasks))

	// Process each task in a separate goroutine
	for i, task := range tasks {
		wg.Add(1)
		go func(i int, t Task) {
			defer wg.Done()
			results[i] = processTask(t)
			fmt.Printf("Completed: %s\n", results[i])
		}(i, task)
	}

	wg.Wait()
	return results
}

func main() {
	// Create some sample tasks
	tasks := []Task{
		{ID: 1, Duration: 200 * time.Millisecond},
		{ID: 2, Duration: 100 * time.Millisecond},
		{ID: 3, Duration: 300 * time.Millisecond},
		{ID: 4, Duration: 150 * time.Millisecond},
	}

	fmt.Println("Starting task processing...")
	start := time.Now()

	// Process tasks concurrently
	results := taskProcessor(tasks)

	// Print summary
	fmt.Printf("\nAll tasks completed in %v\n", time.Since(start))
	fmt.Println("\nResults summary:")
	for i, result := range results {
		fmt.Printf("Task %d: %s\n", i+1, result)
	}
}
