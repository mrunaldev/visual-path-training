package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"secure-service/security"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

// Server represents the secure service
type Server struct {
	router      *mux.Router
	auth        *security.AuthManager
	crypto      *security.Crypto
	rateLimiter *rate.Limiter
}

// NewServer creates a new server instance
func NewServer(jwtSecret []byte) (*Server, error) {
	crypto, err := security.NewCrypto()
	if err != nil {
		return nil, err
	}

	server := &Server{
		router:      mux.NewRouter(),
		auth:        security.NewAuthManager(jwtSecret),
		crypto:      crypto,
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 10),
	}

	server.setupRoutes()
	return server, nil
}

// setupRoutes configures server routes
func (s *Server) setupRoutes() {
	// Middleware
	s.router.Use(s.loggingMiddleware)
	s.router.Use(s.securityHeadersMiddleware)
	s.router.Use(s.rateLimitMiddleware)

	// Public routes
	s.router.HandleFunc("/register", s.handleRegister).Methods("POST")
	s.router.HandleFunc("/login", s.handleLogin).Methods("POST")

	// Protected routes
	protected := s.router.PathPrefix("/api").Subrouter()
	protected.Use(s.authMiddleware)

	protected.HandleFunc("/encrypt", s.handleEncrypt).Methods("POST")
	protected.HandleFunc("/decrypt", s.handleDecrypt).Methods("POST")
}

// Run starts the server
func (s *Server) Run(addr string) error {
	srv := &http.Server{
		Handler:      s.router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server starting on %s", addr)
	return srv.ListenAndServeTLS("cert.pem", "key.pem")
}

// Middleware functions

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
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

func (s *Server) securityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.rateLimiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, err := s.auth.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims to request context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler functions

func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.auth.RegisterUser(req.Username, req.Password, req.Role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := s.auth.Authenticate(req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func (s *Server) handleEncrypt(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Data string `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate random key for symmetric encryption
	key, err := security.GenerateRandomKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encrypt data
	ciphertext, err := s.crypto.EncryptSymmetric(key, []byte(req.Data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"ciphertext": security.EncodeBase64(ciphertext),
		"key":        security.EncodeBase64(key),
	})
}

func (s *Server) handleDecrypt(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Ciphertext string `json:"ciphertext"`
		Key        string `json:"key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decode base64
	ciphertext, err := security.DecodeBase64(req.Ciphertext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key, err := security.DecodeBase64(req.Key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decrypt data
	plaintext, err := s.crypto.DecryptSymmetric(key, ciphertext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"plaintext": string(plaintext),
	})
}

func main() {
	// Generate random JWT secret
	jwtSecret, err := security.GenerateRandomKey()
	if err != nil {
		log.Fatal(err)
	}

	server, err := NewServer(jwtSecret)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Run(":8443"))
}
