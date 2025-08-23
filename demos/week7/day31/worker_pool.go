package main

import (
	"context"
	"fmt"
	"time"
)

// WorkerResult represents the result of a worker's computation
type WorkerResult struct {
	WorkerID int
	JobID    int
	Result   int
}

// Worker represents a worker in the pool
type Worker struct {
	ID     int
	Jobs   <-chan int
	Output chan<- WorkerResult
	Done   <-chan struct{}
}

// NewWorker creates a new worker
func NewWorker(id int, jobs <-chan int, output chan<- WorkerResult, done <-chan struct{}) *Worker {
	return &Worker{
		ID:     id,
		Jobs:   jobs,
		Output: output,
		Done:   done,
	}
}

// Start starts the worker processing jobs
func (w *Worker) Start() {
	go func() {
		for {
			select {
			case <-w.Done:
				fmt.Printf("Worker %d shutting down\n", w.ID)
				return
			case job, ok := <-w.Jobs:
				if !ok {
					fmt.Printf("Worker %d: job channel closed\n", w.ID)
					return
				}
				// Process job
				time.Sleep(100 * time.Millisecond) // Simulate work
				result := WorkerResult{
					WorkerID: w.ID,
					JobID:    job,
					Result:   job * 2,
				}
				w.Output <- result
			}
		}
	}()
}

// WorkerPool manages a pool of workers
type WorkerPool struct {
	Workers    []*Worker
	Jobs       chan int
	Results    chan WorkerResult
	Done       chan struct{}
	NumWorkers int
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		Workers:    make([]*Worker, numWorkers),
		Jobs:       make(chan int),
		Results:    make(chan WorkerResult),
		Done:       make(chan struct{}),
		NumWorkers: numWorkers,
	}
}

// Start starts all workers in the pool
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.NumWorkers; i++ {
		worker := NewWorker(i+1, wp.Jobs, wp.Results, wp.Done)
		wp.Workers[i] = worker
		worker.Start()
	}
}

// Submit submits a job to the pool
func (wp *WorkerPool) Submit(job int) {
	wp.Jobs <- job
}

// Stop stops all workers in the pool
func (wp *WorkerPool) Stop() {
	close(wp.Done)
	close(wp.Jobs)
}

// CollectResults collects results from workers with timeout
func (wp *WorkerPool) CollectResults(ctx context.Context, expectedResults int) []WorkerResult {
	results := make([]WorkerResult, 0, expectedResults)
	resultCount := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context deadline exceeded")
			return results
		case result := <-wp.Results:
			results = append(results, result)
			resultCount++
			if resultCount >= expectedResults {
				return results
			}
		}
	}
}

func main() {
	// Create a worker pool with 3 workers
	pool := NewWorkerPool(3)
	pool.Start()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Submit 10 jobs
	numJobs := 10
	go func() {
		for i := 1; i <= numJobs; i++ {
			pool.Submit(i)
		}
	}()

	// Collect and print results
	results := pool.CollectResults(ctx, numJobs)
	for _, result := range results {
		fmt.Printf("Worker %d processed job %d with result %d\n",
			result.WorkerID, result.JobID, result.Result)
	}

	// Stop the worker pool
	pool.Stop()
	fmt.Println("Worker pool stopped")
}
