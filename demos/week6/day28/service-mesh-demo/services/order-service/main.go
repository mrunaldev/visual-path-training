package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Order represents an order in the system
type Order struct {
	ID         string    `json:"id"`
	ProductID  string    `json:"product_id"`
	Quantity   int       `json:"quantity"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"create_time"`
}

var (
	productServiceURL = os.Getenv("PRODUCT_SERVICE_URL")
	version           = os.Getenv("SERVICE_VERSION")
	tracer            = otel.Tracer("order-service")
)

func main() {
	if productServiceURL == "" {
		productServiceURL = "http://product-service:8080"
	}

	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	// API endpoints
	api := r.PathPrefix("/api").Subrouter()
	api.Handle("/orders", otelhttp.NewHandler(http.HandlerFunc(createOrder), "create-order")).Methods("POST")
	api.Handle("/orders/{id}", otelhttp.NewHandler(http.HandlerFunc(getOrder), "get-order")).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Order service %s starting on port %s", version, port)
	log.Fatal(srv.ListenAndServe())
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := trace.SpanFromContext(ctx)
	defer span.End()

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate product exists
	productURL := fmt.Sprintf("%s/api/products/%s", productServiceURL, order.ProductID)
	resp, err := http.Get(productURL)
	if err != nil {
		http.Error(w, "Failed to validate product", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	// Process order
	order.ID = fmt.Sprintf("ord-%d", time.Now().UnixNano())
	order.Status = "created"
	order.CreateTime = time.Now()

	// Simulate processing time
	time.Sleep(200 * time.Millisecond)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	span := trace.SpanFromContext(ctx)
	defer span.End()

	vars := mux.Vars(r)
	orderID := vars["id"]

	// Simulate database lookup
	time.Sleep(100 * time.Millisecond)

	// For demo purposes, return a mock order
	order := Order{
		ID:         orderID,
		ProductID:  "1",
		Quantity:   1,
		Status:     "processed",
		CreateTime: time.Now().Add(-time.Hour),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
