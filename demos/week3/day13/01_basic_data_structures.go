// Package main demonstrates basic data structures in Go
package main

import (
	"fmt"
)

func main() {
	// Arrays
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Slices
	slice := []string{"apple", "banana", "orange"}
	slice = append(slice, "grape")
	fmt.Println("Slice:", slice)

	// Make with capacity
	numbers2 := make([]int, 0, 5)
	numbers2 = append(numbers2, 1, 2, 3)
	fmt.Printf("Slice length: %d, capacity: %d\n", len(numbers2), cap(numbers2))

	// Slicing
	fmt.Println("Slice[1:3]:", slice[1:3])

	// Maps
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}

	// Map operations
	ages["David"] = 40
	delete(ages, "Bob")

	// Check existence
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice's age: %d\n", age)
	}

	// Range over map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
