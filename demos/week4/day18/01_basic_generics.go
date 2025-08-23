// Package generics demonstrates basic generic functions and types
package generics

import (
	"golang.org/x/exp/constraints"
)

// Stack is a generic stack implementation
type Stack[T any] struct {
	items []T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// FindMax finds the maximum value in a slice of ordered types
func FindMax[T constraints.Ordered](items []T) (T, bool) {
	if len(items) == 0 {
		var zero T
		return zero, false
	}

	max := items[0]
	for _, item := range items[1:] {
		if item > max {
			max = item
		}
	}
	return max, true
}

// Map applies a function to each element in a slice
func Map[T, U any](items []T, f func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = f(item)
	}
	return result
}

// Example usage in main package:
/*
func main() {
	// Stack of integers
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	for !intStack.IsEmpty() {
		if value, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\n", value)
		}
	}

	// Stack of strings
	stringStack := &Stack[string]{}
	stringStack.Push("hello")
	stringStack.Push("world")

	for !stringStack.IsEmpty() {
		if value, ok := stringStack.Pop(); ok {
			fmt.Printf("Popped: %s\n", value)
		}
	}

	// Finding maximum
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	if max, ok := FindMax(numbers); ok {
		fmt.Printf("Maximum number: %d\n", max)
	}

	// Mapping
	squares := Map(numbers, func(x int) int {
		return x * x
	})
	fmt.Printf("Squares: %v\n", squares)
}
*/
