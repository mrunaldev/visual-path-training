// Package main demonstrates method declaration and usage in Go
package main

import (
	"fmt"
	"math"
)

// Rectangle represents a geometric rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle represents a geometric circle
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 5}

	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())
	fmt.Printf("Circle Area: %.2f\n", circle.Area())
}
