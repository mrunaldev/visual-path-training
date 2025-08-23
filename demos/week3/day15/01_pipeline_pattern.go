// Package pipeline demonstrates the pipeline concurrency pattern
package pipeline

// Generator - converts a list of integers to a channel
func Generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Square - squares numbers from input channel and sends to output channel
func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Filter - filters out odd numbers
func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 == 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// Example usage in main package:
/*
func main() {
    // Create a pipeline
    numbers := Generator(1, 2, 3, 4, 5)
    squares := Square(numbers)
    evens := Filter(squares)

    // Consume the output
    for n := range evens {
        fmt.Println(n) // Will print even square numbers
    }
}
*/
