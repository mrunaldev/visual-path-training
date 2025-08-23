// Package main demonstrates basic channel concepts
package main

import (
	"fmt"
	"time"
)

// producer generates numbers and sends them to the channel
func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i // Send value to channel
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close channel when done
}

// consumer receives numbers from the channel
func consumer(ch <-chan int) {
	// Range over channel until it's closed
	for num := range ch {
		fmt.Printf("Consuming: %d\n", num)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	// Create an unbuffered channel
	ch := make(chan int)

	// Start producer and consumer goroutines
	go producer(ch)
	go consumer(ch)

	// Wait for some time to let them work
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}
