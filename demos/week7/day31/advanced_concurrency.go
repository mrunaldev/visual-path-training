package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker Pool Pattern
func workerPool(numWorkers int, jobs <-chan int, results chan<- int) {
	// Create multiple workers
	for i := 1; i <= numWorkers; i++ {
		worker := i
		go func() {
			for j := range jobs {
				fmt.Printf("Worker %d processing job %d\n", worker, j)
				time.Sleep(100 * time.Millisecond) // Simulate work
				results <- j * 2
			}
		}()
	}
}

// Pipeline Pattern
func generator(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-done:
				return
			case out <- n * n:
			}
		}
	}()
	return out
}

func double(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-done:
				return
			case out <- n * 2:
			}
		}
	}()
	return out
}

// Fan-out Fan-in Pattern
func fanOut(done <-chan struct{}, ch <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	for i := 0; i < workers; i++ {
		outputs[i] = square(done, ch)
	}
	return outputs
}

func fanIn(done <-chan struct{}, channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch <-chan int) {
			defer wg.Done()
			for val := range ch {
				select {
				case <-done:
					return
				case merged <- val:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	// Worker Pool Example
	fmt.Println("Worker Pool Pattern:")
	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	workerPool(3, jobs, results)

	// Send jobs
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	for i := 1; i <= numJobs; i++ {
		<-results
	}

	fmt.Println("\nPipeline Pattern:")
	done := make(chan struct{})
	defer close(done)

	// Create pipeline
	numbers := generator(done)
	squares := square(done, numbers)
	doubles := double(done, squares)

	// Process pipeline results
	for result := range doubles {
		fmt.Printf("Pipeline result: %d\n", result)
	}

	fmt.Println("\nFan-out Fan-in Pattern:")
	input := generator(done)
	workers := fanOut(done, input, 3)
	merged := fanIn(done, workers...)

	// Process fan-in results
	for result := range merged {
		fmt.Printf("Fan-in result: %d\n", result)
	}
}
