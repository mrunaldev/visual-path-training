package calculator

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
		{"zeros", 0, 0, 0},
		{"decimals", 1.5, 2.5, 4},
	}

	c := &Calculator{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := c.Add(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Add(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive division", 6, 2, 3},
		{"negative division", -6, 2, -3},
		{"division by zero", 6, 0, 0}, // returns 0 instead of panicking
		{"decimal division", 5.5, 2, 2.75},
	}

	c := &Calculator{}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := c.Divide(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Divide(%v, %v) = %v; want %v", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestCalculator_Memory(t *testing.T) {
	c := &Calculator{}

	// Test storing and recalling
	c.Store(42)
	if got := c.Recall(); got != 42 {
		t.Errorf("After Store(42), Recall() = %v; want 42", got)
	}

	// Test clearing memory
	c.ClearMemory()
	if got := c.Recall(); got != 0 {
		t.Errorf("After ClearMemory(), Recall() = %v; want 0", got)
	}
}

func BenchmarkCalculator_Add(b *testing.B) {
	c := &Calculator{}
	for i := 0; i < b.N; i++ {
		c.Add(2, 3)
	}
}

func BenchmarkCalculator_Power(b *testing.B) {
	c := &Calculator{}
	for i := 0; i < b.N; i++ {
		c.Power(2, 8)
	}
}

func ExampleCalculator_Add() {
	c := &Calculator{}
	result := c.Add(2, 3)
	println(result)
	// Output: 5
}

// Helper function example
func setupCalculator(t *testing.T) (*Calculator, func()) {
	t.Helper() // Marks this as a helper function
	c := &Calculator{}

	// Return the calculator and a cleanup function
	return c, func() {
		c.ClearMemory()
	}
}
