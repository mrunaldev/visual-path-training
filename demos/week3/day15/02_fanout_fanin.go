// Package fanout demonstrates fan-out/fan-in concurrency patterns
package fanout

import (
	"sync"
)

// Worker processes items from input channel
type Worker func(int) int

// FanOut distributes work across multiple goroutines
func FanOut(input <-chan int, workers int, fn Worker) <-chan int {
	channels := make([]<-chan int, workers)

	for i := 0; i < workers; i++ {
		channels[i] = worker(input, fn)
	}

	return FanIn(channels...)
}

// worker processes work from input channel
func worker(input <-chan int, fn Worker) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range input {
			out <- fn(n)
		}
	}()
	return out
}

// FanIn combines multiple channels into one
func FanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexed := make(chan int)

	// Function to forward values from a channel to multiplexed
	forward := func(ch <-chan int) {
		defer wg.Done()
		for value := range ch {
			multiplexed <- value
		}
	}

	// Start goroutine for each input channel
	wg.Add(len(channels))
	for _, ch := range channels {
		go forward(ch)
	}

	// Start goroutine to close multiplexed channel when all inputs are done
	go func() {
		wg.Wait()
		close(multiplexed)
	}()

	return multiplexed
}

// Example usage in main package:
/*
func main() {
    // Create input channel
    input := make(chan int)

    // Start sending numbers
    go func() {
        for i := 0; i < 100; i++ {
            input <- i
        }
        close(input)
    }()

    // Define worker function
    worker := func(x int) int {
        time.Sleep(100 * time.Millisecond) // Simulate work
        return x * 2
    }

    // Fan out processing to 5 workers
    results := FanOut(input, 5, worker)

    // Collect results
    for result := range results {
        fmt.Println(result)
    }
}
*/
