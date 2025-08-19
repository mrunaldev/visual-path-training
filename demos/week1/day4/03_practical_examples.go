// practical_examples.go
package main

import (
    "fmt"
    "strings"
)

// User represents a basic user profile
type User struct {
    name     string
    age      int
    role     string
    active   bool
    settings map[string]string
}

func main() {
    // User validation example
    user := User{
        name:     "John Doe",
        age:      25,
        role:     "user",
        active:   true,
        settings: map[string]string{"theme": "dark"},
    }

    // Validate user with if-else
    if user.name == "" {
        fmt.Println("Error: Name is required")
    } else if user.age < 18 {
        fmt.Println("Error: Must be 18 or older")
    } else if !user.active {
        fmt.Println("Error: Account is not active")
    } else {
        fmt.Println("User is valid!")
    }

    // Role-based access control with switch
    switch strings.ToLower(user.role) {
    case "admin":
        fmt.Println("Full access granted")
    case "moderator":
        fmt.Println("Moderate access granted")
    case "user":
        fmt.Println("Basic access granted")
    default:
        fmt.Println("No access")
    }

    // Settings checker
    checkUserSettings(user)

    // Temperature classifier
    temperatures := []int{-5, 0, 15, 25, 35}
    for _, temp := range temperatures {
        classifyTemperature(temp)
    }

    // Password strength checker
    checkPasswordStrength("abc123")
    checkPasswordStrength("Abc123!@#")
}

func checkUserSettings(user User) {
    // Check theme setting
    if theme, exists := user.settings["theme"]; exists {
        switch theme {
        case "dark":
            fmt.Println("Dark theme enabled")
        case "light":
            fmt.Println("Light theme enabled")
        default:
            fmt.Println("Using default theme")
        }
    } else {
        fmt.Println("No theme setting found")
    }
}

func classifyTemperature(temp int) {
    switch {
    case temp < 0:
        fmt.Printf("%d°C: Freezing!\n", temp)
    case temp < 10:
        fmt.Printf("%d°C: Cold\n", temp)
    case temp < 20:
        fmt.Printf("%d°C: Cool\n", temp)
    case temp < 30:
        fmt.Printf("%d°C: Warm\n", temp)
    default:
        fmt.Printf("%d°C: Hot!\n", temp)
    }
}

func checkPasswordStrength(password string) {
    // Initialize score
    score := 0
    
    // Check length
    if len(password) >= 8 {
        score++
    }
    
    // Check for numbers
    if strings.ContainsAny(password, "0123456789") {
        score++
    }
    
    // Check for uppercase
    if strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
        score++
    }
    
    // Check for special characters
    if strings.ContainsAny(password, "!@#$%^&*()") {
        score++
    }

    // Evaluate password strength
    switch score {
    case 0:
        fmt.Printf("Password '%s': Very Weak\n", password)
    case 1:
        fmt.Printf("Password '%s': Weak\n", password)
    case 2:
        fmt.Printf("Password '%s': Moderate\n", password)
    case 3:
        fmt.Printf("Password '%s': Strong\n", password)
    case 4:
        fmt.Printf("Password '%s': Very Strong\n", password)
    }
}
