// Package calculator_test demonstrates testing for the calculator package
package calculator_test

import (
	"testing"

	"calculator"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(2, 3)
	expected := 5.0
	if result != expected {
		t.Errorf("Add(2, 3) = %f; expected %f", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := calculator.Divide(6, 2)
	expected := 3.0
	if result != expected {
		t.Errorf("Divide(6, 2) = %f; expected %f", result, expected)
	}

	// Test division by zero
	result = calculator.Divide(6, 0)
	expected = 0.0
	if result != expected {
		t.Errorf("Divide(6, 0) = %f; expected %f", result, expected)
	}
}
