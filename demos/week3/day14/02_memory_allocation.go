// Package main demonstrates memory allocation and management
package main

import (
	"fmt"
	"runtime"
)

// BigStruct represents a structure that uses significant memory
type BigStruct struct {
	data [1000000]int // ~4MB of data
}

// newBigStruct creates a new BigStruct on the heap
func newBigStruct() *BigStruct {
	return &BigStruct{}
}

// processData simulates processing with memory allocation
func processData() {
	// Create some big objects
	var objects []*BigStruct

	// Allocate memory
	for i := 0; i < 5; i++ {
		objects = append(objects, newBigStruct())
		printMemStats("After allocation")
	}

	// Clear references to trigger GC
	objects = nil
	runtime.GC()
	printMemStats("After cleanup")
}

// printMemStats prints current memory statistics
func printMemStats(msg string) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	fmt.Printf("\n=== %s ===\n", msg)
	fmt.Printf("Heap Alloc: %d MB\n", stats.HeapAlloc/1024/1024)
	fmt.Printf("Total Alloc: %d MB\n", stats.TotalAlloc/1024/1024)
	fmt.Printf("System Memory: %d MB\n", stats.Sys/1024/1024)
	fmt.Printf("GC Cycles: %d\n", stats.NumGC)
}

func main() {
	fmt.Println("Starting memory management demo...")
	printMemStats("Initial state")

	processData()

	fmt.Println("\nRunning garbage collection...")
	runtime.GC()
	printMemStats("Final state")
}
