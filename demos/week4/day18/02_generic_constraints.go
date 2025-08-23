// Package generics demonstrates generic constraints and types
package generics

// Number is a constraint that permits any number type
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Pair represents a generic key-value pair
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// NewPair creates a new key-value pair
func NewPair[K comparable, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{
		Key:   key,
		Value: value,
	}
}

// Dictionary implements a generic map-like structure
type Dictionary[K comparable, V any] struct {
	items map[K]V
}

// NewDictionary creates a new dictionary
func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{
		items: make(map[K]V),
	}
}

// Set adds or updates a key-value pair
func (d *Dictionary[K, V]) Set(key K, value V) {
	d.items[key] = value
}

// Get retrieves a value by key
func (d *Dictionary[K, V]) Get(key K) (V, bool) {
	value, exists := d.items[key]
	return value, exists
}

// Delete removes a key-value pair
func (d *Dictionary[K, V]) Delete(key K) {
	delete(d.items, key)
}

// Keys returns all keys in the dictionary
func (d *Dictionary[K, V]) Keys() []K {
	keys := make([]K, 0, len(d.items))
	for k := range d.items {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values in the dictionary
func (d *Dictionary[K, V]) Values() []V {
	values := make([]V, 0, len(d.items))
	for _, v := range d.items {
		values = append(values, v)
	}
	return values
}

// Example usage in main package:
/*
func main() {
	// Create a dictionary with string keys and int values
	scores := NewDictionary[string, int]()

	// Add some scores
	scores.Set("Alice", 95)
	scores.Set("Bob", 87)
	scores.Set("Charlie", 92)

	// Get a score
	if score, exists := scores.Get("Alice"); exists {
		fmt.Printf("Alice's score: %d\n", score)
	}

	// List all students and scores
	fmt.Println("All students:", scores.Keys())
	fmt.Println("All scores:", scores.Values())

	// Create pairs
	pair1 := NewPair("x", 10)
	pair2 := NewPair("y", 20.5)

	fmt.Printf("Pair 1: %v = %v\n", pair1.Key, pair1.Value)
	fmt.Printf("Pair 2: %v = %v\n", pair2.Key, pair2.Value)
}
*/
