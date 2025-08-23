// Package main demonstrates advanced data structure implementations
package main

import (
	"fmt"
	"sync"
)

// Stack represents a thread-safe stack implementation
type Stack struct {
	lock  sync.Mutex
	items []interface{}
}

// Push adds an item to the stack
func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (interface{}, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return nil, false
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

// Queue represents a thread-safe queue implementation
type Queue struct {
	lock  sync.Mutex
	items []interface{}
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item from the queue
func (q *Queue) Dequeue() (interface{}, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.items) == 0 {
		return nil, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func main() {
	// Stack example
	stack := &Stack{}

	// Push items
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Pop items
	for i := 0; i < 4; i++ {
		if item, ok := stack.Pop(); ok {
			fmt.Printf("Popped from stack: %v\n", item)
		} else {
			fmt.Println("Stack is empty")
		}
	}

	// Queue example
	queue := &Queue{}

	// Enqueue items
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	// Dequeue items
	for i := 0; i < 4; i++ {
		if item, ok := queue.Dequeue(); ok {
			fmt.Printf("Dequeued from queue: %v\n", item)
		} else {
			fmt.Println("Queue is empty")
		}
	}
}
