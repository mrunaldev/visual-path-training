package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Standard for loop
	fmt.Println("1. Standard for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// 2. Multiple variables in loop
	fmt.Println("2. Multiple variables in loop:")
	for i, j := 0, 10; i < 5; i, j = i+1, j+2 {
		fmt.Printf("i: %d, j: %d\n", i, j)
	}
	fmt.Println()

	// 3. Condition-only loop (while loop equivalent)
	fmt.Println("3. Condition-only loop:")
	sum := 1
	for sum < 100 {
		sum *= 2
		fmt.Printf("%d ", sum)
	}
	fmt.Println("\n")

	// 4. Break statement
	fmt.Println("4. Break statement:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// 5. Continue statement
	fmt.Println("5. Continue statement:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// 6. Infinite loop with break
	fmt.Println("6. Infinite loop with break:")
	count := 0
	for {
		fmt.Printf("%d ", count)
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println("\n")

	// 7. Nested loops
	fmt.Println("7. Nested loops:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()

	// 8. Loop with labels
	fmt.Println("8. Loop with labels:")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println("\n")

	// 9. Loop with timer
	fmt.Println("9. Loop with timer:")
	timeout := time.After(2 * time.Second)
	count = 0
loop:
	for {
		select {
		case <-timeout:
			fmt.Println("\nTimeout reached!")
			break loop
		default:
			fmt.Printf("%d ", count)
			count++
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println()

	// 10. Loop with defer
	fmt.Println("\n10. Loop with defer:")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Println("(Deferred prints follow)")
}
