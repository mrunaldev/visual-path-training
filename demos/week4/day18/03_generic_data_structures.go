// Package generics demonstrates practical generic data structures
package generics

// Result represents a computation result with potential error
type Result[T any] struct {
	Value T
	Error error
}

// Queue implements a generic FIFO queue
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an item to the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// LinkedList represents a generic linked list node
type LinkedList[T any] struct {
	Value T
	Next  *LinkedList[T]
}

// NewLinkedList creates a new linked list with a single node
func NewLinkedList[T any](value T) *LinkedList[T] {
	return &LinkedList[T]{
		Value: value,
		Next:  nil,
	}
}

// Add appends a value to the end of the list
func (l *LinkedList[T]) Add(value T) {
	current := l
	for current.Next != nil {
		current = current.Next
	}
	current.Next = NewLinkedList(value)
}

// ToSlice converts the linked list to a slice
func (l *LinkedList[T]) ToSlice() []T {
	var result []T
	current := l
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// Example usage in main package:
/*
func main() {
	// Queue example
	queue := NewQueue[string]()
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	for !queue.IsEmpty() {
		if value, ok := queue.Dequeue(); ok {
			fmt.Printf("Dequeued: %s\n", value)
		}
	}

	// Linked list example
	list := NewLinkedList(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	fmt.Println("List as slice:", list.ToSlice())

	// Result example
	results := []Result[int]{
		{Value: 42, Error: nil},
		{Value: 0, Error: errors.New("computation failed")},
	}

	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		} else {
			fmt.Printf("Success: %v\n", result.Value)
		}
	}
}
*/
