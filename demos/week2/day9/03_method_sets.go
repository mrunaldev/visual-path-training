// Package main demonstrates method sets in Go
package main

import (
	"fmt"
	"strings"
)

// StringManipulator provides string manipulation methods
type StringManipulator string

// ToUpper converts string to uppercase
func (s StringManipulator) ToUpper() string {
	return strings.ToUpper(string(s))
}

// ToLower converts string to lowercase
func (s StringManipulator) ToLower() string {
	return strings.ToLower(string(s))
}

// Reverse reverses the string
func (s StringManipulator) Reverse() string {
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Chain demonstrates method chaining
func (s StringManipulator) Chain() StringManipulator {
	return StringManipulator(s.ToUpper().Reverse())
}

func main() {
	text := StringManipulator("Hello, World!")

	fmt.Println("Original:", text)
	fmt.Println("Uppercase:", text.ToUpper())
	fmt.Println("Lowercase:", text.ToLower())
	fmt.Println("Reversed:", text.Reverse())
	fmt.Println("Chained:", text.Chain())
}
