// constants_and_operators.go
package main

import "fmt"

// Constants declaration
const Pi = 3.14159
const (
    StatusOK       = 200
    StatusNotFound = 404
    MaxConnections = 1000
)

// Using iota
const (
    Sunday = iota    // 0
    Monday          // 1
    Tuesday         // 2
    Wednesday       // 3
    Thursday        // 4
    Friday          // 5
    Saturday        // 6
)

func main() {
    // Using constants
    fmt.Println("=== Constants ===")
    fmt.Printf("Pi: %f\n", Pi)
    fmt.Printf("Status OK: %d\n", StatusOK)
    fmt.Printf("Monday: %d\n", Monday)

    // Arithmetic operators
    a, b := 15, 4
    fmt.Println("\n=== Arithmetic Operators ===")
    fmt.Printf("a = %d, b = %d\n", a, b)
    fmt.Printf("Addition: %d\n", a+b)
    fmt.Printf("Subtraction: %d\n", a-b)
    fmt.Printf("Multiplication: %d\n", a*b)
    fmt.Printf("Division: %d\n", a/b)
    fmt.Printf("Remainder: %d\n", a%b)

    // Assignment operators
    x := 10
    fmt.Println("\n=== Assignment Operators ===")
    fmt.Printf("Initial x: %d\n", x)
    x += 5
    fmt.Printf("After x += 5: %d\n", x)
    x -= 3
    fmt.Printf("After x -= 3: %d\n", x)
    x *= 2
    fmt.Printf("After x *= 2: %d\n", x)
    x /= 4
    fmt.Printf("After x /= 4: %d\n", x)

    // Comparison operators
    fmt.Println("\n=== Comparison Operators ===")
    fmt.Printf("a == b: %t\n", a == b)
    fmt.Printf("a != b: %t\n", a != b)
    fmt.Printf("a < b: %t\n", a < b)
    fmt.Printf("a > b: %t\n", a > b)
    fmt.Printf("a <= b: %t\n", a <= b)
    fmt.Printf("a >= b: %t\n", a >= b)

    // Logical operators
    p, q := true, false
    fmt.Println("\n=== Logical Operators ===")
    fmt.Printf("p: %t, q: %t\n", p, q)
    fmt.Printf("p && q: %t\n", p && q)
    fmt.Printf("p || q: %t\n", p || q)
    fmt.Printf("!p: %t\n", !p)

    // Operator precedence
    result := 5 + 3*2
    fmt.Println("\n=== Operator Precedence ===")
    fmt.Printf("5 + 3 * 2 = %d\n", result)
    result = (5 + 3) * 2
    fmt.Printf("(5 + 3) * 2 = %d\n", result)
}
