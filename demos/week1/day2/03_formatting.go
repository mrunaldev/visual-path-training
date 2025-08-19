// formatting.go
package main

import "fmt"

func main() {
    // Basic variables
    name := "Alice"
    age := 25
    height := 1.75
    isStudent := true

    // Basic formatting
    fmt.Println("=== Basic Formatting ===")
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Height: %.2f\n", height)
    fmt.Printf("Is Student: %t\n", isStudent)

    // Width and padding
    fmt.Println("\n=== Width and Padding ===")
    fmt.Printf("Name: %10s\n", name)      // Right-aligned, width 10
    fmt.Printf("Name: %-10s\n", name)     // Left-aligned, width 10
    fmt.Printf("Age: %05d\n", age)        // Zero-padded

    // Different number formats
    fmt.Println("\n=== Number Formats ===")
    number := 42
    fmt.Printf("Decimal: %d\n", number)
    fmt.Printf("Binary: %b\n", number)
    fmt.Printf("Octal: %o\n", number)
    fmt.Printf("Hexadecimal: %x\n", number)

    // Type information
    fmt.Println("\n=== Type Information ===")
    fmt.Printf("Variable: %v, Type: %T\n", name, name)
    fmt.Printf("Variable: %v, Type: %T\n", age, age)
    fmt.Printf("Variable: %v, Type: %T\n", height, height)
    fmt.Printf("Variable: %v, Type: %T\n", isStudent, isStudent)

    // String operations
    fmt.Println("\n=== String Operations ===")
    fmt.Printf("Quoted string: %q\n", name)
    fmt.Printf("Unicode point: %U\n", 'A')

    // Multiple values
    fmt.Println("\n=== Multiple Values ===")
    fmt.Printf("Profile: %s is %d years old and %.2f meters tall\n", 
        name, age, height)
}
