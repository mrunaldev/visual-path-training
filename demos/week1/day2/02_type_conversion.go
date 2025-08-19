// type_conversion.go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Integer to Float conversion
    age := 25
    ageFloat := float64(age)
    fmt.Printf("Age as float: %.2f\n", ageFloat)

    // Float to Integer conversion
    height := 1.75
    heightInt := int(height)
    fmt.Printf("Height as integer: %d\n", heightInt)

    // String to Number conversions
    ageString := "25"
    ageFromString, err := strconv.Atoi(ageString)
    if err != nil {
        fmt.Println("Error converting string to integer:", err)
    } else {
        fmt.Printf("Age from string: %d\n", ageFromString)
    }

    // Number to String conversion
    score := 95
    scoreString := strconv.Itoa(score)
    fmt.Printf("Score as string: %s\n", scoreString)

    // String to Float conversion
    weightString := "68.5"
    weightFloat, err := strconv.ParseFloat(weightString, 64)
    if err != nil {
        fmt.Println("Error converting string to float:", err)
    } else {
        fmt.Printf("Weight from string: %.2f\n", weightFloat)
    }

    // Boolean conversions
    boolString := "true"
    boolValue, err := strconv.ParseBool(boolString)
    if err != nil {
        fmt.Println("Error converting string to boolean:", err)
    } else {
        fmt.Printf("Boolean from string: %t\n", boolValue)
    }
}
