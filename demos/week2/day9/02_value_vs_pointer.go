// Package main demonstrates value vs pointer receivers in Go
package main

import (
	"fmt"
)

// Counter demonstrates value receiver
type Counter struct {
	count int
}

// Value receiver - doesn't modify the original Counter
func (c Counter) IncrementValue() {
	c.count++
}

// Pointer receiver - modifies the original Counter
func (c *Counter) IncrementPointer() {
	c.count++
}

// Person demonstrates when to use pointer receivers
type Person struct {
	Name string
	Age  int
}

// UpdateAge uses pointer receiver for modification
func (p *Person) UpdateAge(newAge int) {
	p.Age = newAge
}

// String uses value receiver for reading
func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func main() {
	// Counter example
	counter := Counter{count: 0}

	counter.IncrementValue()
	fmt.Printf("After value increment: %d\n", counter.count) // Still 0

	counter.IncrementPointer()
	fmt.Printf("After pointer increment: %d\n", counter.count) // Now 1

	// Person example
	person := Person{Name: "Alice", Age: 30}
	fmt.Println(person.String())

	person.UpdateAge(31)
	fmt.Println(person.String())
}
