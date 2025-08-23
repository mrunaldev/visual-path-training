package calculator

import (
	"testing"
)

func FuzzCalculator_Add(f *testing.F) {
	c := &Calculator{}

	// Add some seed values
	f.Add(float64(1), float64(2))
	f.Add(float64(-1), float64(1))
	f.Add(float64(0), float64(0))

	// Fuzz test
	f.Fuzz(func(t *testing.T, a float64, b float64) {
		result := c.Add(a, b)

		// Properties that should always hold true
		if a > 0 && b > 0 && result <= 0 {
			t.Errorf("Add(%v, %v) = %v; expected positive result for positive inputs", a, b, result)
		}

		if result-a != b {
			t.Errorf("Add(%v, %v) = %v; does not satisfy result-a=b", a, b, result)
		}
	})
}
