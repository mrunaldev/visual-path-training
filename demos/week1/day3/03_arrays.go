// arrays.go
package main

import "fmt"

func main() {
    // Basic array declaration
    var numbers [5]int
    fmt.Println("Empty array:", numbers)

    // Array initialization
    colors := [3]string{"red", "green", "blue"}
    fmt.Println("Colors:", colors)

    // Array with implicit size
    scores := [...]int{95, 89, 78, 92, 85}
    fmt.Println("Scores:", scores)

    // Accessing and modifying elements
    fmt.Println("\n=== Array Operations ===")
    fmt.Println("First color:", colors[0])
    colors[1] = "yellow"
    fmt.Println("Modified colors:", colors)

    // Array length
    fmt.Printf("Number of scores: %d\n", len(scores))

    // Iterating over array
    fmt.Println("\n=== Using traditional for loop ===")
    for i := 0; i < len(scores); i++ {
        fmt.Printf("Score %d: %d\n", i+1, scores[i])
    }

    fmt.Println("\n=== Using range ===")
    for index, value := range scores {
        fmt.Printf("Score %d: %d\n", index+1, value)
    }

    // Multi-dimensional array
    fmt.Println("\n=== 2D Array ===")
    matrix := [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    // Printing 2D array
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("%d ", matrix[i][j])
        }
        fmt.Println()
    }

    // Array calculations
    fmt.Println("\n=== Array Calculations ===")
    sum := 0
    for _, score := range scores {
        sum += score
    }
    average := float64(sum) / float64(len(scores))
    fmt.Printf("Average score: %.2f\n", average)

    // Finding max value
    max := scores[0]
    for _, score := range scores {
        if score > max {
            max = score
        }
    }
    fmt.Printf("Highest score: %d\n", max)
}
