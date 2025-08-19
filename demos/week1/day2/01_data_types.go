// data_types.go
package main

import "fmt"

func main() {
    // Integer types
    var age int = 25
    var temperature int8 = 35
    var population uint64 = 7800000000

    // Floating point
    var height float64 = 1.75
    var weight float32 = 68.5

    // Boolean
    var isStudent bool = true
    var isEmployed = false

    // String
    var name string = "Gopher"
    var message = "Welcome to Go!"

    // Printing all variables with their types
    fmt.Printf("Age: %d (Type: %T)\n", age, age)
    fmt.Printf("Temperature: %d (Type: %T)\n", temperature, temperature)
    fmt.Printf("Population: %d (Type: %T)\n", population, population)
    fmt.Printf("Height: %.2f (Type: %T)\n", height, height)
    fmt.Printf("Weight: %.1f (Type: %T)\n", weight, weight)
    fmt.Printf("Is Student: %t (Type: %T)\n", isStudent, isStudent)
    fmt.Printf("Is Employed: %t (Type: %T)\n", isEmployed, isEmployed)
    fmt.Printf("Name: %s (Type: %T)\n", name, name)
    fmt.Printf("Message: %s (Type: %T)\n", message, message)
}
