// switch_statements.go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Basic switch
    day := "Wednesday"
    switch day {
    case "Monday":
        fmt.Println("Start of work week")
    case "Wednesday":
        fmt.Println("Mid week")
    case "Friday":
        fmt.Println("TGIF!")
    default:
        fmt.Println("Regular day")
    }

    // Switch with multiple cases
    hour := time.Now().Hour()
    switch {
    case hour < 12:
        fmt.Println("Good morning!")
    case hour < 17:
        fmt.Println("Good afternoon!")
    default:
        fmt.Println("Good evening!")
    }

    // Switch with fallthrough
    num := 1
    switch num {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three")
    }

    // Type switch
    printType("Hello")
    printType(42)
    printType(true)
    printType(3.14)

    // Switch with initialization
    switch os := getOS(); os {
    case "darwin":
        fmt.Println("Mac OS")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Printf("%s\n", os)
    }

    // Switch without expression
    score := 85
    switch {
    case score >= 90:
        fmt.Println("Grade: A")
    case score >= 80:
        fmt.Println("Grade: B")
    case score >= 70:
        fmt.Println("Grade: C")
    default:
        fmt.Println("Grade: F")
    }
}

func printType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func getOS() string {
    return "linux"
}
