// Package main demonstrates the select statement with channels
package main

import (
	"fmt"
	"time"
)

// generateNumbers sends numbers to a channel
func generateNumbers(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

// generateLetters sends letters to a channel
func generateLetters(ch chan<- string) {
	letters := []string{"A", "B", "C", "D", "E"}
	for _, letter := range letters {
		ch <- letter
		time.Sleep(150 * time.Millisecond)
	}
	close(ch)
}

func main() {
	// Create channels
	numbers := make(chan int)
	letters := make(chan string)
	done := make(chan bool)

	// Start producers
	go generateNumbers(numbers)
	go generateLetters(letters)

	// Start consumer with select
	go func() {
		for {
			select {
			case num, ok := <-numbers:
				if !ok {
					numbers = nil // Disable this case
					if numbers == nil && letters == nil {
						done <- true
						return
					}
					continue
				}
				fmt.Printf("Received number: %d\n", num)

			case letter, ok := <-letters:
				if !ok {
					letters = nil // Disable this case
					if numbers == nil && letters == nil {
						done <- true
						return
					}
					continue
				}
				fmt.Printf("Received letter: %s\n", letter)

			case <-time.After(500 * time.Millisecond):
				fmt.Println("Timeout!")
				done <- true
				return
			}
		}
	}()

	<-done // Wait for completion
	fmt.Println("All done!")
}
