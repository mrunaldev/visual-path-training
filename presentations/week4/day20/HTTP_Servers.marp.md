---
marp: true
theme: default
paginate: true
---

# HTTP Servers in Go
## Week 4 - Day 20

---

# Today's Topics

1. HTTP Basics
2. `net/http` Package
3. Routing & Handlers
4. Middleware
5. REST APIs

---

# HTTP Basics

```go
http.StatusOK           // 200
http.StatusCreated      // 201
http.StatusNotFound     // 404
http.StatusBadRequest   // 400
http.StatusInternalServerError // 500
```

- Request/Response cycle
- HTTP methods
- Status codes
- Headers and body

---

# Basic HTTP Server

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

---

# Request Handling

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Access request details
    method := r.Method
    path := r.URL.Path
    query := r.URL.Query()
    headers := r.Header
    
    // Read body
    body, err := io.ReadAll(r.Body)
    defer r.Body.Close()
}
```

---

# Response Writing

```go
func handler(w http.ResponseWriter, r *http.Request) {
    // Set headers
    w.Header().Set("Content-Type", "application/json")
    
    // Set status
    w.WriteHeader(http.StatusOK)
    
    // Write response
    json.NewEncoder(w).Encode(response)
}
```

---

# Routing Patterns

```go
mux := http.NewServeMux()

mux.HandleFunc("/", homeHandler)
mux.HandleFunc("/api/", apiHandler)
mux.HandleFunc("/api/users/", usersHandler)

http.ListenAndServe(":8080", mux)
```

---

# Middleware Example

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

mux.Handle("/", loggingMiddleware(handler))
```

---

# Static File Server

```go
// Serve files from ./static directory
fs := http.FileServer(http.Dir("static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))

// Serve single file
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
})
```

---

# JSON Handling

```go
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func handleUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Process user...
}
```

---

# REST API Example

```go
func main() {
    router := http.NewServeMux()
    
    router.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:    getAllUsers(w, r)
        case http.MethodPost:   createUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
}
```

---

# Best Practices

1. Use HTTPS in production
2. Implement proper error handling
3. Add request validation
4. Include logging/monitoring
5. Use middleware for common tasks
6. Follow REST conventions
7. Document your API

---

# Exercise Time!

1. Create a REST API
2. Add middleware
3. Handle JSON data
4. Serve static files
5. Add error handling

---

# Questions?

Let's practice with hands-on examples!
