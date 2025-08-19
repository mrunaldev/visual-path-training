// interactive.go
package main

import "fmt"

func main() {
    var name string
    
    fmt.Print("What's your name? ")
    fmt.Scan(&name)
    
    fmt.Printf("Hello, %s! Welcome to Go programming!\n", name)
    
    // Show different ways to print
    fmt.Println("This is a regular print line")
    fmt.Printf("This is formatted: %s\n", name)
    fmt.Print("This prints without a newline")
    fmt.Print(" - see?\n")
}
