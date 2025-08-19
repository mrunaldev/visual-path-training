package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 1. Fibonacci Generator
func generateFibonacci(n int) []int {
	fib := make([]int, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			fib[i] = i
		} else {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}
	return fib
}

// 2. Pattern Printer
func printPyramid(height int) {
	for i := 0; i < height; i++ {
		// Print spaces
		for j := 0; j < height-i-1; j++ {
			fmt.Print(" ")
		}
		// Print stars
		for k := 0; k <= i*2; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 3. Data Processing
func processNumbers(numbers []int) (min, max, sum int) {
	if len(numbers) == 0 {
		return 0, 0, 0
	}
	
	min = numbers[0]
	max = numbers[0]
	sum = 0

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}
	return
}

// 4. Search Implementation
func searchElement(arr []int, target int) (index int, found bool) {
	for i, num := range arr {
		if num == target {
			return i, true
		}
	}
	return -1, false
}

// 5. Retry Mechanism
func retryOperation(maxAttempts int) error {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		// Simulate an operation that might fail
		if rand.Float64() < 0.7 { // 70% chance of failure
			fmt.Printf("Attempt %d failed\n", attempt)
			time.Sleep(time.Second) // Wait before retry
			continue
		}
		fmt.Printf("Attempt %d succeeded\n", attempt)
		return nil
	}
	return fmt.Errorf("failed after %d attempts", maxAttempts)
}

// 6. Word Counter
func countWords(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)
	
	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}
	return wordCount
}

// 7. Matrix Spiral Print
func spiralPrint(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// Print top row
		for i := left; i <= right; i++ {
			fmt.Printf("%d ", matrix[top][i])
		}
		top++

		// Print right column
		for i := top; i <= bottom; i++ {
			fmt.Printf("%d ", matrix[i][right])
		}
		right--

		if top <= bottom {
			// Print bottom row
			for i := right; i >= left; i-- {
				fmt.Printf("%d ", matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// Print left column
			for i := bottom; i >= top; i-- {
				fmt.Printf("%d ", matrix[i][left])
			}
			left++
		}
	}
}

// 8. Batch Processing
func processBatch(items []int, batchSize int, process func([]int)) {
	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}
		batch := items[i:end]
		process(batch)
	}
}

func main() {
	// 1. Fibonacci Example
	fmt.Println("1. Fibonacci Series:")
	fib := generateFibonacci(10)
	fmt.Println(fib)
	fmt.Println()

	// 2. Pattern Example
	fmt.Println("2. Pyramid Pattern:")
	printPyramid(5)
	fmt.Println()

	// 3. Data Processing Example
	fmt.Println("3. Data Processing:")
	numbers := []int{23, 45, 12, 67, 89, 34, 12, 56}
	min, max, sum := processNumbers(numbers)
	fmt.Printf("Min: %d, Max: %d, Sum: %d, Average: %.2f\n",
		min, max, sum, float64(sum)/float64(len(numbers)))
	fmt.Println()

	// 4. Search Example
	fmt.Println("4. Search Implementation:")
	target := 67
	if index, found := searchElement(numbers, target); found {
		fmt.Printf("Found %d at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found\n", target)
	}
	fmt.Println()

	// 5. Retry Example
	fmt.Println("5. Retry Mechanism:")
	rand.Seed(time.Now().UnixNano())
	err := retryOperation(5)
	if err != nil {
		fmt.Println("Operation failed:", err)
	}
	fmt.Println()

	// 6. Word Counter Example
	fmt.Println("6. Word Counter:")
	text := "The quick brown fox jumps over the lazy dog. The dog sleeps."
	wordCount := countWords(text)
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
	fmt.Println()

	// 7. Spiral Matrix Example
	fmt.Println("7. Spiral Matrix:")
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("Spiral Order:")
	spiralPrint(matrix)
	fmt.Println("\n")

	// 8. Batch Processing Example
	fmt.Println("8. Batch Processing:")
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	batchSize := 3
	processBatch(items, batchSize, func(batch []int) {
		fmt.Printf("Processing batch: %v\n", batch)
	})
}
