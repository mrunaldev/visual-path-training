package main

import "fmt"

func main() {
	// 1. Range over slice
	fmt.Println("1. Range over slice:")
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}
	fmt.Println()

	// 2. Range over array
	fmt.Println("2. Range over array:")
	colors := [3]string{"red", "green", "blue"}
	for i, color := range colors {
		fmt.Printf("Position %d: %s\n", i, color)
	}
	fmt.Println()

	// 3. Range over map
	fmt.Println("3. Range over map:")
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
	fmt.Println()

	// 4. Range over string (by rune)
	fmt.Println("4. Range over string:")
	str := "Hello, 世界"
	for i, char := range str {
		fmt.Printf("Position %d: %c (Unicode: %U)\n", i, char, char)
	}
	fmt.Println()

	// 5. Range with blank identifier (ignoring index)
	fmt.Println("5. Range with blank identifier:")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println("\n")

	// 6. Range over map keys only
	fmt.Println("6. Range over map keys only:")
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}
	fmt.Println()

	// 7. Range over channel
	fmt.Println("7. Range over channel:")
	ch := make(chan int, 5)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for num := range ch {
		fmt.Printf("%d ", num)
	}
	fmt.Println("\n")

	// 8. Range over nested slices
	fmt.Println("8. Range over nested slices:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i, row := range matrix {
		for j, val := range row {
			fmt.Printf("(%d,%d)=%d ", i, j, val)
		}
		fmt.Println()
	}
	fmt.Println()

	// 9. Range over map with sorted keys
	fmt.Println("9. Range over map with sorted keys:")
	scores := map[string]int{
		"Math":    90,
		"English": 85,
		"Science": 95,
	}
	// Get keys
	subjects := make([]string, 0, len(scores))
	for subject := range scores {
		subjects = append(subjects, subject)
	}
	// Sort keys (we'll learn about sort package later)
	for i := 0; i < len(subjects)-1; i++ {
		for j := i + 1; j < len(subjects); j++ {
			if subjects[i] > subjects[j] {
				subjects[i], subjects[j] = subjects[j], subjects[i]
			}
		}
	}
	// Print in sorted order
	for _, subject := range subjects {
		fmt.Printf("%s: %d\n", subject, scores[subject])
	}
	fmt.Println()

	// 10. Range with type inference
	fmt.Println("10. Range with type inference:")
	mixedMap := map[interface{}]interface{}{
		"name":   "John",
		42:       "answer",
		"scores": []int{85, 90, 95},
	}
	for key, value := range mixedMap {
		fmt.Printf("Type of key: %-14T Value: %v\n", key, value)
	}
}
