// Package main demonstrates pointer basics and memory concepts
package main

import (
	"fmt"
)

// Value demonstrates passing by value
func modifyValue(x int) {
	x = 42 // This modification is local to the function
	fmt.Printf("Inside modifyValue: x = %d\n", x)
}

// Pointer demonstrates passing by reference
func modifyPointer(x *int) {
	*x = 42 // This modification affects the original variable
	fmt.Printf("Inside modifyPointer: x = %d\n", *x)
}

// Person demonstrates struct with pointers
type Person struct {
	Name string
	Age  *int // Pointer to allow nil age
}

func main() {
	// Basic pointer usage
	x := 10
	var ptr *int = &x

	fmt.Printf("x = %d\n", x)
	fmt.Printf("ptr points to %d\n", *ptr)

	// Pass by value vs pointer
	original := 10
	fmt.Printf("\nBefore modifyValue: original = %d\n", original)
	modifyValue(original)
	fmt.Printf("After modifyValue: original = %d\n", original)

	fmt.Printf("\nBefore modifyPointer: original = %d\n", original)
	modifyPointer(&original)
	fmt.Printf("After modifyPointer: original = %d\n", original)

	// Struct with pointer field
	age := 25
	person := Person{
		Name: "Alice",
		Age:  &age,
	}

	fmt.Printf("\nPerson: %s, Age: %d\n", person.Name, *person.Age)

	// Nil pointer example
	var nilPerson Person
	fmt.Printf("Nil age person: %s, Age: %v\n", nilPerson.Name, nilPerson.Age)
}
