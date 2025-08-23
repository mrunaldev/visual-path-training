package generics

import (
	"golang.org/x/exp/constraints"
)

// Number is a constraint that permits any number type
type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Min returns the minimum of two comparable values
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two comparable values
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Sum returns the sum of a slice of numbers
func Sum[T Number](numbers []T) T {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// Filter returns a new slice containing only the elements for which the predicate returns true
func Filter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map applies a transformation function to each element in a slice
func Map[T, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = transform(item)
	}
	return result
}

// Reduce reduces a slice to a single value using a reducer function
func Reduce[T, U any](items []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, item := range items {
		result = reducer(result, item)
	}
	return result
}

// Set implements a generic set data structure
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet creates a new empty set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// Add adds an item to the set
func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Remove removes an item from the set
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Contains checks if an item exists in the set
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// Size returns the number of items in the set
func (s *Set[T]) Size() int {
	return len(s.items)
}

// Items returns all items in the set as a slice
func (s *Set[T]) Items() []T {
	result := make([]T, 0, len(s.items))
	for item := range s.items {
		result = append(result, item)
	}
	return result
}
