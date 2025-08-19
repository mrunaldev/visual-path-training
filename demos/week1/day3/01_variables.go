// variables.go
package main

import "fmt"

// Global variable declaration
var globalCounter int = 0

func main() {
    // Different ways to declare variables
    var age int = 25                  // Explicit type and value
    var name = "Gopher"              // Type inference
    score := 95                      // Short declaration

    // Multiple variable declaration
    var x, y int = 10, 20
    width, height := 100, 50

    // Zero values
    var (
        defaultInt    int
        defaultFloat  float64
        defaultBool   bool
        defaultString string
    )

    // Printing variables
    fmt.Println("=== Variable Values ===")
    fmt.Printf("age: %d\n", age)
    fmt.Printf("name: %s\n", name)
    fmt.Printf("score: %d\n", score)
    fmt.Printf("x: %d, y: %d\n", x, y)
    fmt.Printf("width: %d, height: %d\n", width, height)

    fmt.Println("\n=== Zero Values ===")
    fmt.Printf("defaultInt: %d\n", defaultInt)
    fmt.Printf("defaultFloat: %f\n", defaultFloat)
    fmt.Printf("defaultBool: %t\n", defaultBool)
    fmt.Printf("defaultString: %q\n", defaultString)

    // Scope demonstration
    {
        // Block scope
        innerVar := "I'm inside the block"
        fmt.Println("\n=== Block Scope ===")
        fmt.Println(innerVar)
    }
    // fmt.Println(innerVar)  // This would cause an error

    // Variable shadowing
    count := 5
    {
        count := 10  // This creates a new variable
        fmt.Printf("\nInner count: %d\n", count)
    }
    fmt.Printf("Outer count: %d\n", count)
}
