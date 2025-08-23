package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

// Gateway represents the API Gateway
type Gateway struct {
	router      *mux.Router
	rateLimiter *rate.Limiter
	services    map[string]string
	mu          sync.RWMutex
}

// ServiceResponse represents an aggregated response
type ServiceResponse struct {
	Service string      `json:"service"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

// NewGateway creates a new API Gateway
func NewGateway() *Gateway {
	g := &Gateway{
		router:      mux.NewRouter(),
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 100), // 100 requests per second
		services:    make(map[string]string),
	}

	g.setupRoutes()
	return g
}

// setupRoutes configures the gateway routes
func (g *Gateway) setupRoutes() {
	g.router.Use(g.authMiddleware)
	g.router.Use(g.rateLimitMiddleware)
	g.router.Use(g.loggingMiddleware)

	// API routes
	api := g.router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/aggregate", g.handleAggregate).Methods("GET")
	api.HandleFunc("/services", g.handleListServices).Methods("GET")
	api.HandleFunc("/services/{name}", g.handleRegisterService).Methods("POST")
}

// Middleware functions

func (g *Gateway) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// TODO: Implement proper token validation
		next.ServeHTTP(w, r)
	})
}

func (g *Gateway) rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !g.rateLimiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (g *Gateway) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s %v",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}

// Route handlers

func (g *Gateway) handleAggregate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	services := r.URL.Query()["service"]
	if len(services) == 0 {
		http.Error(w, "No services specified", http.StatusBadRequest)
		return
	}

	responses := g.aggregateResponses(ctx, services)
	json.NewEncoder(w).Encode(responses)
}

func (g *Gateway) handleListServices(w http.ResponseWriter, r *http.Request) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	json.NewEncoder(w).Encode(g.services)
}

func (g *Gateway) handleRegisterService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceName := vars["name"]

	var serviceURL string
	if err := json.NewDecoder(r.Body).Decode(&serviceURL); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	g.mu.Lock()
	g.services[serviceName] = serviceURL
	g.mu.Unlock()

	w.WriteHeader(http.StatusCreated)
}

// Helper functions

func (g *Gateway) aggregateResponses(ctx context.Context, services []string) []ServiceResponse {
	var wg sync.WaitGroup
	responses := make([]ServiceResponse, len(services))

	for i, service := range services {
		wg.Add(1)
		go func(i int, service string) {
			defer wg.Done()
			responses[i] = g.callService(ctx, service)
		}(i, service)
	}

	wg.Wait()
	return responses
}

func (g *Gateway) callService(ctx context.Context, service string) ServiceResponse {
	g.mu.RLock()
	serviceURL, ok := g.services[service]
	g.mu.RUnlock()

	if !ok {
		return ServiceResponse{
			Service: service,
			Error:   "Service not found",
		}
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", serviceURL, nil)
	if err != nil {
		return ServiceResponse{
			Service: service,
			Error:   err.Error(),
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return ServiceResponse{
			Service: service,
			Error:   err.Error(),
		}
	}
	defer resp.Body.Close()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ServiceResponse{
			Service: service,
			Error:   err.Error(),
		}
	}

	return ServiceResponse{
		Service: service,
		Data:    data,
	}
}

func main() {
	gateway := NewGateway()
	log.Printf("Starting API Gateway on :8080")
	log.Fatal(http.ListenAndServe(":8080", gateway.router))
}
