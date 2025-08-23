// Package main demonstrates struct types and custom data structures
package main

import (
	"fmt"
	"time"
)

// Person represents a person with basic information
type Person struct {
	ID        int
	FirstName string
	LastName  string
	BirthDate time.Time
	Address   Address
}

// Address represents a physical address
type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

// NewPerson creates a new Person instance
func NewPerson(firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
	}
}

// FullName returns the person's full name
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// SetAddress updates the person's address
func (p *Person) SetAddress(street, city, state, zip string) {
	p.Address = Address{
		Street:  street,
		City:    city,
		State:   state,
		ZipCode: zip,
	}
}

func main() {
	// Create a new person
	person := NewPerson("John", "Doe")

	// Set address
	person.SetAddress("123 Main St", "Anytown", "ST", "12345")

	// Print person details
	fmt.Printf("Name: %s\n", person.FullName())
	fmt.Printf("Address: %s, %s, %s %s\n",
		person.Address.Street,
		person.Address.City,
		person.Address.State,
		person.Address.ZipCode)
}
