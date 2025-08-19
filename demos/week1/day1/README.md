# Day 1: Introduction to Go Programming
## Demo Code Guide

This guide walks through the demo programs for Day 1 of our Go Programming course.

## 1. Hello World Program (01_hello.go)

The simplest Go program that prints "Hello, Go!" to the console.

### Key Concepts:
- Package declaration
- Import statements
- Main function
- Print function

### Running the Program:
```bash
go run 01_hello.go
```

## 2. Basic Types Demo (02_basic_types.go)

Demonstrates basic data types in Go and simple arithmetic operations.

### Key Concepts:
- Variable declaration
- Basic types (string, int, float, bool)
- Short declaration operator (:=)
- Basic arithmetic operations
- Printing variables

### Running the Program:
```bash
go run 02_basic_types.go
```

### Expected Output:
```
Name: Gopher
Age: 5
Height: 1.5
Is Active: true

Basic Arithmetic:
Addition: 15
Subtraction: 5
Multiplication: 50
Division: 2
```

## 3. Interactive Program (03_interactive.go)

Shows how to get user input and different ways to print output.

### Key Concepts:
- User input with fmt.Scan
- Different print methods:
  - fmt.Print
  - fmt.Println
  - fmt.Printf
- String formatting

### Running the Program:
```bash
go run 03_interactive.go
```

### Example Interaction:
```
What's your name? Alice
Hello, Alice! Welcome to Go programming!
This is a regular print line
This is formatted: Alice
This prints without a newline - see?
```

## Practice Exercises

1. **Modify Hello World**
   - Add more print statements
   - Print your name
   - Print today's date

2. **Extend Basic Types**
   - Add more arithmetic operations
   - Try different number types
   - Create more variables

3. **Enhance Interactive Program**
   - Ask for user's age
   - Ask for user's favorite number
   - Print a personalized message

## Common Issues and Solutions

1. **Program doesn't compile**
   - Check package declaration
   - Verify all imports are used
   - Look for syntax errors

2. **fmt.Scan doesn't work**
   - Make sure to use & before variable
   - Check variable type matches input
   - Press Enter after input

3. **Printing issues**
   - Check string formatting verbs
   - Verify number of arguments
   - Check for missing newlines

## Best Practices

1. **Code Organization**
   - One concept per file
   - Clear variable names
   - Consistent formatting

2. **Documentation**
   - Comment your code
   - Explain complex parts
   - Use meaningful names

3. **Testing**
   - Try different inputs
   - Handle errors gracefully
   - Test edge cases

## Next Steps

1. Review the code samples
2. Complete practice exercises
3. Experiment with modifications
4. Read Go documentation
5. Prepare questions for next session
