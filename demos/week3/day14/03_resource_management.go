// Package main demonstrates memory leaks and best practices
package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// ResourceManager demonstrates proper resource management
type ResourceManager struct {
	resources map[string]io.Closer
	mu        sync.Mutex
}

// NewResourceManager creates a new resource manager
func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		resources: make(map[string]io.Closer),
	}
}

// AddResource adds a resource to manage
func (rm *ResourceManager) AddResource(name string, resource io.Closer) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.resources[name] = resource
}

// CloseResource closes and removes a specific resource
func (rm *ResourceManager) CloseResource(name string) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if resource, exists := rm.resources[name]; exists {
		delete(rm.resources, name)
		return resource.Close()
	}
	return nil
}

// CloseAll closes all resources
func (rm *ResourceManager) CloseAll() {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	for name, resource := range rm.resources {
		if err := resource.Close(); err != nil {
			fmt.Printf("Error closing %s: %v\n", name, err)
		}
		delete(rm.resources, name)
	}
}

// TempFile represents a temporary file resource
type TempFile struct {
	*os.File
}

// simulateWork simulates some work with resources
func simulateWork(rm *ResourceManager) {
	// Create temporary files
	for i := 1; i <= 3; i++ {
		name := fmt.Sprintf("temp_%d", i)
		file, err := os.CreateTemp("", name)
		if err != nil {
			fmt.Printf("Error creating temp file: %v\n", err)
			continue
		}

		rm.AddResource(name, &TempFile{file})
		fmt.Printf("Created resource: %s\n", name)
	}

	// Simulate some work
	time.Sleep(time.Second)

	// Clean up resources properly
	rm.CloseAll()
}

func main() {
	fmt.Println("Starting resource management demo...")

	rm := NewResourceManager()

	// Run garbage collection before
	runtime.GC()
	printMemoryStats("Before work")

	simulateWork(rm)

	// Run garbage collection after
	runtime.GC()
	printMemoryStats("After work")
}

func printMemoryStats(msg string) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	fmt.Printf("\n=== %s ===\n", msg)
	fmt.Printf("Heap Objects: %d\n", stats.HeapObjects)
	fmt.Printf("Heap Alloc: %d MB\n", stats.HeapAlloc/1024/1024)
	fmt.Printf("GC Cycles: %d\n", stats.NumGC)
}
