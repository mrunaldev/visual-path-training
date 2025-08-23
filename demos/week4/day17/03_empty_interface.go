// Package interfaces demonstrates empty interface and type assertions
package interfaces

import (
	"fmt"
	"strconv"
)

// ValueProcessor handles different types of values
type ValueProcessor struct {
	values []interface{}
}

// Add adds a value to the processor
func (vp *ValueProcessor) Add(value interface{}) {
	vp.values = append(vp.values, value)
}

// ProcessValues demonstrates type assertions and type switches
func (vp *ValueProcessor) ProcessValues() []string {
	var results []string

	for _, value := range vp.values {
		// Type switch for different types
		switch v := value.(type) {
		case int:
			results = append(results, fmt.Sprintf("Integer: %d", v))
		case float64:
			results = append(results, fmt.Sprintf("Float: %.2f", v))
		case string:
			// Try to convert string to number
			if num, err := strconv.Atoi(v); err == nil {
				results = append(results, fmt.Sprintf("String number: %d", num))
			} else {
				results = append(results, fmt.Sprintf("String: %s", v))
			}
		case []int:
			sum := 0
			for _, n := range v {
				sum += n
			}
			results = append(results, fmt.Sprintf("Int slice sum: %d", sum))
		case nil:
			results = append(results, "Nil value")
		default:
			results = append(results, fmt.Sprintf("Unknown type: %T", v))
		}
	}

	return results
}

// StringConverter converts values to strings
type StringConverter interface {
	ToString() string
}

// CustomType demonstrates implementing StringConverter
type CustomType struct {
	ID   int
	Name string
}

func (ct CustomType) ToString() string {
	return fmt.Sprintf("CustomType{ID: %d, Name: %s}", ct.ID, ct.Name)
}

// ProcessCustomValues handles values implementing StringConverter
func (vp *ValueProcessor) ProcessCustomValues() []string {
	var results []string

	for _, value := range vp.values {
		// Try type assertion for StringConverter
		if converter, ok := value.(StringConverter); ok {
			results = append(results, converter.ToString())
		} else {
			results = append(results, fmt.Sprintf("Not convertible: %v", value))
		}
	}

	return results
}

// Example usage in main package:
/*
func main() {
    processor := &ValueProcessor{}

    // Add various types of values
    processor.Add(42)
    processor.Add(3.14)
    processor.Add("123")
    processor.Add("hello")
    processor.Add([]int{1, 2, 3})
    processor.Add(nil)
    processor.Add(CustomType{ID: 1, Name: "Test"})

    // Process regular values
    fmt.Println("Regular processing:")
    for _, result := range processor.ProcessValues() {
        fmt.Println(result)
    }

    // Process custom values
    fmt.Println("\nCustom processing:")
    for _, result := range processor.ProcessCustomValues() {
        fmt.Println(result)
    }
}
*/
