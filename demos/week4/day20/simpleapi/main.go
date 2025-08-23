package main

import (
	"log"
	"net/http"
	"sync"
)

// User represents a user in our system
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UserStore manages user data with thread-safe operations
type UserStore struct {
	sync.RWMutex
	users  map[int]User
	nextID int
}

// NewUserStore creates a new UserStore
func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

// GetAllUsers returns all users
func (s *UserStore) GetAllUsers() []User {
	s.RLock()
	defer s.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// GetUser returns a user by ID
func (s *UserStore) GetUser(id int) (User, bool) {
	s.RLock()
	defer s.RUnlock()

	user, exists := s.users[id]
	return user, exists
}

// CreateUser adds a new user
func (s *UserStore) CreateUser(user User) User {
	s.Lock()
	defer s.Unlock()

	user.ID = s.nextID
	s.users[user.ID] = user
	s.nextID++
	return user
}

// UpdateUser updates an existing user
func (s *UserStore) UpdateUser(id int, user User) (User, bool) {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.users[id]; !exists {
		return User{}, false
	}

	user.ID = id
	s.users[id] = user
	return user, true
}

// DeleteUser removes a user
func (s *UserStore) DeleteUser(id int) bool {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.users[id]; !exists {
		return false
	}

	delete(s.users, id)
	return true
}

func main() {
	store := NewUserStore()
	server := NewServer(store)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", server.router))
}
