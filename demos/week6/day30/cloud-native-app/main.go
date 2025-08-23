package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
)

// Metrics
var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

// App represents our cloud-native application
type App struct {
	server *http.Server
	done   chan bool
	wg     sync.WaitGroup
}

// NewApp creates a new application instance
func NewApp() *App {
	app := &App{
		done: make(chan bool),
	}

	// Register metrics
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)

	// Create router
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/health", app.healthHandler)
	mux.HandleFunc("/ready", app.readinessHandler)
	mux.HandleFunc("/api/data", app.dataHandler)

	// Configure server
	app.server = &http.Server{
		Addr:         ":8080",
		Handler:      app.metricsMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return app
}

// Run starts the application
func (a *App) Run() error {
	// Start server
	go func() {
		log.Printf("Server starting on %s", a.server.Addr)
		if err := a.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	// Handle shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")
	return a.Shutdown()
}

// Shutdown gracefully shuts down the server
func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	a.server.SetKeepAlivesEnabled(false)
	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}

	close(a.done)
	a.wg.Wait()
	log.Println("Server shutdown complete")
	return nil
}

// Middleware

func (a *App) metricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create metrics wrapper for response
		wrapped := newResponseWriter(w)

		// Handle request
		next.ServeHTTP(wrapped, r)

		// Record metrics
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
		requestCounter.WithLabelValues(
			r.Method,
			r.URL.Path,
			wrapped.status(),
		).Inc()
	})
}

// Handlers

func (a *App) healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func (a *App) readinessHandler(w http.ResponseWriter, r *http.Request) {
	// Add your readiness checks here
	json.NewEncoder(w).Encode(map[string]string{
		"status": "READY",
		"time":   time.Now().Format(time.RFC3339),
	})
}

func (a *App) dataHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tracer := otel.Tracer("api")

	ctx, span := tracer.Start(ctx, "data-handler")
	defer span.End()

	// Simulate some work
	time.Sleep(100 * time.Millisecond)

	data := map[string]interface{}{
		"message": "Hello from cloud-native app",
		"time":    time.Now().Format(time.RFC3339),
	}

	span.AddEvent("sending response")
	json.NewEncoder(w).Encode(data)
}

// Helper types

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) status() string {
	return http.StatusText(rw.statusCode)
}

func main() {
	app := NewApp()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
