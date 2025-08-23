package calculator

// Calculator represents a simple calculator
type Calculator struct {
	memory float64
}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference between two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
// If b is 0, returns 0 to avoid panic
func (c *Calculator) Divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// Memory operations

// Store stores a number in memory
func (c *Calculator) Store(n float64) {
	c.memory = n
}

// Recall returns the number from memory
func (c *Calculator) Recall() float64 {
	return c.memory
}

// ClearMemory sets memory to 0
func (c *Calculator) ClearMemory() {
	c.memory = 0
}

// Power returns a to the power of b
func (c *Calculator) Power(a, b float64) float64 {
	result := float64(1)
	if b < 0 {
		return 0 // simplified handling of negative exponents
	}
	for i := float64(0); i < b; i++ {
		result *= a
	}
	return result
}
