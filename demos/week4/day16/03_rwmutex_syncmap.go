// Package sync_examples demonstrates RWMutex and sync.Map
package sync_examples

import (
	"sync"
	"time"
)

// Cache represents a thread-safe cache using RWMutex
type Cache struct {
	sync.RWMutex
	data map[string]string
}

// NewCache creates a new cache
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (string, bool) {
	c.RLock()
	defer c.RUnlock()
	val, exists := c.data[key]
	return val, exists
}

// Set adds or updates a value in the cache
func (c *Cache) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

// UserSession demonstrates sync.Map usage
type UserSession struct {
	sessions sync.Map
}

// AddSession adds a new user session
func (us *UserSession) AddSession(userID string, lastSeen time.Time) {
	us.sessions.Store(userID, lastSeen)
}

// GetSession retrieves a user's session
func (us *UserSession) GetSession(userID string) (time.Time, bool) {
	value, exists := us.sessions.Load(userID)
	if !exists {
		return time.Time{}, false
	}
	return value.(time.Time), true
}

// RemoveSession removes a user's session
func (us *UserSession) RemoveSession(userID string) {
	us.sessions.Delete(userID)
}

// Example usage in main package:
/*
func main() {
	// Cache example
	cache := NewCache()
	var wg sync.WaitGroup

	// Multiple readers, one writer
	wg.Add(3)
	go func() {
		defer wg.Done()
		cache.Set("key1", "value1")
	}()

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			if val, exists := cache.Get("key1"); exists {
				fmt.Printf("Got value: %s\n", val)
			}
		}()
	}

	wg.Wait()

	// UserSession example
	sessions := &UserSession{}

	// Add sessions
	sessions.AddSession("user1", time.Now())
	sessions.AddSession("user2", time.Now().Add(-1*time.Hour))

	// Range over sessions
	sessions.sessions.Range(func(key, value interface{}) bool {
		userID := key.(string)
		lastSeen := value.(time.Time)
		fmt.Printf("User %s last seen: %v\n", userID, lastSeen)
		return true
	})
}
*/
