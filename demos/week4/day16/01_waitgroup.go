// Package sync_examples demonstrates WaitGroup usage
package sync_examples

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID int
}

// ProcessTasks demonstrates basic WaitGroup usage
func ProcessTasks(tasks []Task) {
	var wg sync.WaitGroup

	// Add number of tasks to wait for
	wg.Add(len(tasks))

	for _, task := range tasks {
		// Launch a goroutine for each task
		go func(t Task) {
			defer wg.Done() // Decrement counter when done
			processTask(t)
		}(task)
	}

	// Wait for all tasks to complete
	wg.Wait()
}

// processTask simulates work
func processTask(task Task) {
	fmt.Printf("Processing task %d\n", task.ID)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Completed task %d\n", task.ID)
}

// Example usage in main package:
/*
func main() {
	// Create sample tasks
	tasks := make([]Task, 5)
	for i := range tasks {
		tasks[i] = Task{ID: i + 1}
	}

	fmt.Println("Starting task processing...")
	ProcessTasks(tasks)
	fmt.Println("All tasks completed!")
}
*/
