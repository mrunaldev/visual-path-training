// Package shapes demonstrates basic interface usage
package shapes

import "math"

// Shape defines the behavior of a geometric shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements the Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle implements the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle implements the Shape interface
type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
	SideC  float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// ShapeProcessor demonstrates processing different shapes
type ShapeProcessor struct {
	shapes []Shape
}

// AddShape adds a shape to the processor
func (sp *ShapeProcessor) AddShape(s Shape) {
	sp.shapes = append(sp.shapes, s)
}

// TotalArea calculates the total area of all shapes
func (sp *ShapeProcessor) TotalArea() float64 {
	total := 0.0
	for _, shape := range sp.shapes {
		total += shape.Area()
	}
	return total
}

// TotalPerimeter calculates the total perimeter of all shapes
func (sp *ShapeProcessor) TotalPerimeter() float64 {
	total := 0.0
	for _, shape := range sp.shapes {
		total += shape.Perimeter()
	}
	return total
}

// Example usage in main package:
/*
func main() {
    processor := &ShapeProcessor{}

    // Add different shapes
    processor.AddShape(Circle{Radius: 5})
    processor.AddShape(Rectangle{Width: 4, Height: 6})
    processor.AddShape(Triangle{
        Base: 3,
        Height: 4,
        SideA: 3,
        SideB: 4,
        SideC: 5,
    })

    fmt.Printf("Total Area: %.2f\n", processor.TotalArea())
    fmt.Printf("Total Perimeter: %.2f\n", processor.TotalPerimeter())
}
*/
