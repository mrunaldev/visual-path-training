package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Server encapsulates the HTTP server and its dependencies
type Server struct {
	store  *UserStore
	router *http.ServeMux
}

// NewServer creates a new server instance with routes configured
func NewServer(store *UserStore) *Server {
	s := &Server{
		store:  store,
		router: http.NewServeMux(),
	}
	s.routes()
	return s
}

// routes sets up all the routes for our server
func (s *Server) routes() {
	// Add middleware to all routes
	s.router.Handle("/", s.logRequest(s.handleIndex()))
	s.router.Handle("/users", s.logRequest(s.handleUsers()))
	s.router.Handle("/users/", s.logRequest(s.handleUser()))

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	s.router.Handle("/static/", http.StripPrefix("/static/", fs))
}

// Middleware function to log requests
func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"%s %s %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

// Handle the index route
func (s *Server) handleIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Welcome to the User API!")
	})
}

// Handle /users route for listing all users and creating new ones
func (s *Server) handleUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			users := s.store.GetAllUsers()
			respondJSON(w, users, http.StatusOK)

		case http.MethodPost:
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			user = s.store.CreateUser(user)
			respondJSON(w, user, http.StatusCreated)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// Handle /users/{id} route for individual user operations
func (s *Server) handleUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := extractID(r.URL.Path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			if user, exists := s.store.GetUser(id); exists {
				respondJSON(w, user, http.StatusOK)
			} else {
				http.NotFound(w, r)
			}

		case http.MethodPut:
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if updated, exists := s.store.UpdateUser(id, user); exists {
				respondJSON(w, updated, http.StatusOK)
			} else {
				http.NotFound(w, r)
			}

		case http.MethodDelete:
			if deleted := s.store.DeleteUser(id); deleted {
				w.WriteHeader(http.StatusNoContent)
			} else {
				http.NotFound(w, r)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

// Helper function to respond with JSON
func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	}
}

// Helper function to extract ID from URL path
func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid path")
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid id")
	}

	return id, nil
}
