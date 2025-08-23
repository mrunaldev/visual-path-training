// Package patterns demonstrates advanced concurrency patterns
package patterns

import (
	"context"
	"sync"
	"time"
)

// Task represents work to be done
type Task struct {
	ID       int
	Duration time.Duration
}

// Result represents the outcome of a task
type Result struct {
	TaskID int
	Data   string
}

// WorkPool manages a pool of workers
type WorkPool struct {
	numWorkers int
	tasks      chan Task
	results    chan Result
	done       chan struct{}
}

// NewWorkPool creates a new work pool
func NewWorkPool(numWorkers int) *WorkPool {
	return &WorkPool{
		numWorkers: numWorkers,
		tasks:      make(chan Task),
		results:    make(chan Result),
		done:       make(chan struct{}),
	}
}

// Start begins the worker pool
func (wp *WorkPool) Start(ctx context.Context) {
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < wp.numWorkers; i++ {
		wg.Add(1)
		go wp.worker(ctx, i, &wg)
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(wp.results)
		close(wp.done)
	}()
}

// worker processes tasks
func (wp *WorkPool) worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task, ok := <-wp.tasks:
			if !ok {
				return
			}
			// Process task
			time.Sleep(task.Duration) // Simulate work
			wp.results <- Result{
				TaskID: task.ID,
				Data:   "Completed",
			}

		case <-ctx.Done():
			return
		}
	}
}

// Submit adds a task to the pool
func (wp *WorkPool) Submit(task Task) {
	wp.tasks <- task
}

// Results returns the results channel
func (wp *WorkPool) Results() <-chan Result {
	return wp.results
}

// Stop gracefully shuts down the pool
func (wp *WorkPool) Stop() {
	close(wp.tasks)
	<-wp.done
}

// Example usage in main package:
/*
func main() {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Create and start work pool
    pool := NewWorkPool(3)
    pool.Start(ctx)

    // Submit tasks
    go func() {
        for i := 1; i <= 5; i++ {
            pool.Submit(Task{
                ID:       i,
                Duration: time.Duration(i) * 100 * time.Millisecond,
            })
        }
        pool.Stop()
    }()

    // Collect results
    for result := range pool.Results() {
        fmt.Printf("Task %d completed\n", result.TaskID)
    }
}
*/
