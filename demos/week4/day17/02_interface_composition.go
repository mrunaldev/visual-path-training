// Package storage demonstrates interface composition and polymorphism
package storage

import (
	"fmt"
	"io"
)

// Reader interface for reading data
type Reader interface {
	Read(data []byte) (int, error)
}

// Writer interface for writing data
type Writer interface {
	Write(data []byte) (int, error)
}

// Storage combines Reader and Writer interfaces
type Storage interface {
	Reader
	Writer
	Close() error
}

// FileStorage implements Storage interface for file-based storage
type FileStorage struct {
	data []byte
}

func (fs *FileStorage) Read(data []byte) (int, error) {
	if len(fs.data) == 0 {
		return 0, io.EOF
	}
	n := copy(data, fs.data)
	fs.data = fs.data[n:]
	return n, nil
}

func (fs *FileStorage) Write(data []byte) (int, error) {
	fs.data = append(fs.data, data...)
	return len(data), nil
}

func (fs *FileStorage) Close() error {
	fs.data = nil
	return nil
}

// MemoryStorage implements Storage interface for memory-based storage
type MemoryStorage struct {
	data []byte
}

func (ms *MemoryStorage) Read(data []byte) (int, error) {
	if len(ms.data) == 0 {
		return 0, io.EOF
	}
	n := copy(data, ms.data)
	ms.data = ms.data[n:]
	return n, nil
}

func (ms *MemoryStorage) Write(data []byte) (int, error) {
	ms.data = append(ms.data, data...)
	return len(data), nil
}

func (ms *MemoryStorage) Close() error {
	ms.data = nil
	return nil
}

// DataProcessor demonstrates polymorphic behavior
type DataProcessor struct {
	storage Storage
}

func NewDataProcessor(s Storage) *DataProcessor {
	return &DataProcessor{storage: s}
}

func (dp *DataProcessor) ProcessData(input []byte) error {
	// Write data
	_, err := dp.storage.Write(input)
	if err != nil {
		return fmt.Errorf("write error: %v", err)
	}

	// Read and process data
	buffer := make([]byte, len(input))
	_, err = dp.storage.Read(buffer)
	if err != nil && err != io.EOF {
		return fmt.Errorf("read error: %v", err)
	}

	return dp.storage.Close()
}

// Example usage in main package:
/*
func main() {
    // Use FileStorage
    fileStorage := &FileStorage{}
    fileProcessor := NewDataProcessor(fileStorage)
    err := fileProcessor.ProcessData([]byte("Hello, File Storage!"))
    if err != nil {
        fmt.Printf("File storage error: %v\n", err)
    }

    // Use MemoryStorage
    memStorage := &MemoryStorage{}
    memProcessor := NewDataProcessor(memStorage)
    err = memProcessor.ProcessData([]byte("Hello, Memory Storage!"))
    if err != nil {
        fmt.Printf("Memory storage error: %v\n", err)
    }
}
*/
