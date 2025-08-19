# Day 3: Variables, Constants, Operators, and Arrays
## Demo Code Guide

This guide walks through the demo programs for Day 3, focusing on variables, constants, operators, and an introduction to arrays in Go.

## 1. Variables (01_variables.go)

Demonstrates different ways to declare and use variables in Go.

### Key Concepts:
- Variable declaration methods
- Type inference
- Zero values
- Scope rules
- Multiple variable declaration
- Variable shadowing

### Running the Program:
```bash
go run 01_variables.go
```

## 2. Constants and Operators (02_constants_and_operators.go)

Shows how to work with constants and various operators in Go.

### Key Concepts:
- Constant declaration
- iota usage
- Arithmetic operators
- Assignment operators
- Comparison operators
- Logical operators
- Operator precedence

### Running the Program:
```bash
go run 02_constants_and_operators.go
```

## 3. Arrays (03_arrays.go)

Introduction to arrays in Go.

### Key Concepts:
- Array declaration
- Array initialization
- Accessing and modifying elements
- Array iteration
- Multi-dimensional arrays
- Array calculations

### Running the Program:
```bash
go run 03_arrays.go
```

## Practice Exercises

1. **Variable Practice**
   - Create variables of different types
   - Use different declaration methods
   - Experiment with scope
   - Try variable shadowing

2. **Constants and Operators**
   - Create constant groups
   - Use iota for enumeration
   - Practice operator precedence
   - Combine different operators

3. **Array Operations**
   - Create and initialize arrays
   - Find array maximum/minimum
   - Calculate averages
   - Process 2D arrays

## Challenge Exercises

1. **Temperature Converter**
```go
// Create a program that:
// - Declares temperature constants (freezing, boiling points)
// - Converts between Celsius and Fahrenheit
// - Uses operators for calculations
```

2. **Grade Calculator**
```go
// Create a program that:
// - Stores student grades in an array
// - Calculates average, highest, and lowest grades
// - Uses constants for grade boundaries
// - Determines letter grades using operators
```

3. **Array Manipulator**
```go
// Create a program that:
// - Creates a 2D array representing a game board
// - Allows moving pieces using operators
// - Validates moves using logical operators
// - Keeps score using constants
```

## Common Issues and Solutions

1. **Variable Scope**
   ```go
   // Wrong
   if x := 5; x > 0 {
       // x is valid
   }
   fmt.Println(x)  // Error: x not defined

   // Right
   x := 5
   if x > 0 {
       // x is valid
   }
   fmt.Println(x)  // Works fine
   ```

2. **Constant Values**
   ```go
   // Wrong
   const x = someFunction()  // Error: const initializer must be constant

   // Right
   const x = 100  // OK
   ```

3. **Array Bounds**
   ```go
   // Wrong
   arr := [3]int{1, 2, 3}
   fmt.Println(arr[3])  // Runtime error: index out of bounds

   // Right
   fmt.Println(arr[len(arr)-1])  // Safely access last element
   ```

## Best Practices

1. **Variable Naming**
   - Use descriptive names
   - Follow camelCase convention
   - Keep names concise but clear

2. **Constants**
   - Group related constants
   - Use iota where appropriate
   - Make values clear and meaningful

3. **Arrays**
   - Check bounds before access
   - Use appropriate size
   - Consider slices for dynamic size

## Next Steps

1. Review the code examples
2. Complete all exercises
3. Try the challenge exercises
4. Read about slices (preview for future sessions)
5. Practice operator combinations

## Additional Resources

- [Go Variables](https://tour.golang.org/basics/8)
- [Go Constants](https://tour.golang.org/basics/15)
- [Go Operators](https://golang.org/ref/spec#Operators)
- [Go Arrays](https://tour.golang.org/moretypes/6)
- [Effective Go](https://golang.org/doc/effective_go)
