// Package main demonstrates practical examples of function usage
package main

import (
	"fmt"
	"time"
)

// calculateFibonacci demonstrates a closure for memoization
func calculateFibonacci() func(int) int {
	cache := make(map[int]int)

	return func(n int) int {
		if val, exists := cache[n]; exists {
			return val
		}

		if n <= 1 {
			cache[n] = n
			return n
		}

		cache[n] = calculateFibonacci()(n-1) + calculateFibonacci()(n-2)
		return cache[n]
	}
}

// timeFuncExecution is a higher-order function that measures execution time
func timeFuncExecution(fn func()) {
	start := time.Now()
	defer func() {
		fmt.Printf("Execution time: %v\n", time.Since(start))
	}()
	fn()
}

func main() {
	// Using fibonacci closure
	fib := calculateFibonacci()
	fmt.Println("Fibonacci(10):", fib(10))
	fmt.Println("Fibonacci(20):", fib(20))

	// Using timer function
	timeFuncExecution(func() {
		sum := 0
		for i := 0; i < 1000000; i++ {
			sum += i
		}
		fmt.Println("Sum calculated")
	})
}
