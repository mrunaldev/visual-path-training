// Package main demonstrates multi-file organization
package main

import (
	"fmt"
	"time"
)

// User represents a system user
type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

// NewUser creates a new user with the given name
func NewUser(name string) *User {
	return &User{
		Name:      name,
		CreatedAt: time.Now(),
	}
}

// String implements the Stringer interface
func (u *User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s, Created: %s}",
		u.ID, u.Name, u.CreatedAt.Format(time.RFC3339))
}

func main() {
	// Create and use a new user
	user := NewUser("Alice")
	fmt.Println(user)

	// Demonstrate time handling
	fmt.Printf("User created %v ago\n",
		time.Since(user.CreatedAt).Round(time.Second))
}
