// if_statements.go
package main

import "fmt"

func main() {
    // Basic if statement
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult")
    }

    // If-else statement
    score := 75
    if score >= 60 {
        fmt.Println("You passed!")
    } else {
        fmt.Println("You failed!")
    }

    // If with initialization
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

    // Nested if statements
    hasID := true
    if age >= 18 {
        if hasID {
            fmt.Println("You can enter the venue")
        } else {
            fmt.Println("You need to show ID")
        }
    } else {
        fmt.Println("You must be 18 or older")
    }

    // Complex conditions
    temperature := 25
    isRaining := false
    if temperature > 20 && !isRaining {
        fmt.Println("Perfect weather for outdoor activities")
    } else if temperature > 20 && isRaining {
        fmt.Println("Warm but rainy")
    } else {
        fmt.Println("Not ideal weather")
    }

    // Using functions in conditions
    if isEven(42) {
        fmt.Println("42 is even")
    }
}

func isEven(n int) bool {
    return n%2 == 0
}
