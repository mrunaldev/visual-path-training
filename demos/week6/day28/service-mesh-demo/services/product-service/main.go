package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Product represents a product in the catalog
type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Version   string  `json:"version"`
	Available bool    `json:"available"`
}

var (
	products = []Product{
		{ID: "1", Name: "Product 1", Price: 19.99, Version: "v1", Available: true},
		{ID: "2", Name: "Product 2", Price: 29.99, Version: "v1", Available: true},
	}
	version = os.Getenv("SERVICE_VERSION")
	tracer  = otel.Tracer("product-service")
)

func main() {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	// API endpoints
	api := r.PathPrefix("/api").Subrouter()
	api.Handle("/products", otelhttp.NewHandler(http.HandlerFunc(getProducts), "get-products")).Methods("GET")
	api.Handle("/products/{id}", otelhttp.NewHandler(http.HandlerFunc(getProduct), "get-product")).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Product service %s starting on port %s", version, port)
	log.Fatal(srv.ListenAndServe())
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := trace.SpanFromContext(ctx)
	defer span.End()

	// Simulate some processing time
	time.Sleep(100 * time.Millisecond)

	// Add version to response
	response := map[string]interface{}{
		"version":  version,
		"products": products,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := trace.SpanFromContext(ctx)
	defer span.End()

	vars := mux.Vars(r)
	id := vars["id"]

	// Simulate some processing time
	time.Sleep(50 * time.Millisecond)

	for _, product := range products {
		if product.ID == id {
			product.Version = version
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Product not found"})
}
