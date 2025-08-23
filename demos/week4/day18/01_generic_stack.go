package main

import "fmt"

// Stack is a generic stack implementation
type Stack[T any] struct {
	items []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}

	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

// Peek returns the item at the top of the stack without removing it
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

func main() {
	// Stack of integers
	intStack := NewStack[int]()
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println("Integer Stack:")
	for !intStack.IsEmpty() {
		if val, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\n", val)
		}
	}

	// Stack of strings
	stringStack := NewStack[string]()
	stringStack.Push("Hello")
	stringStack.Push("World")
	stringStack.Push("!")

	fmt.Println("\nString Stack:")
	for !stringStack.IsEmpty() {
		if val, ok := stringStack.Pop(); ok {
			fmt.Printf("Popped: %s\n", val)
		}
	}

	// Stack of custom type
	type Person struct {
		Name string
		Age  int
	}

	personStack := NewStack[Person]()
	personStack.Push(Person{Name: "Alice", Age: 30})
	personStack.Push(Person{Name: "Bob", Age: 25})
	personStack.Push(Person{Name: "Charlie", Age: 35})

	fmt.Println("\nPerson Stack:")
	for !personStack.IsEmpty() {
		if val, ok := personStack.Pop(); ok {
			fmt.Printf("Popped: %s, %d years old\n", val.Name, val.Age)
		}
	}
}
