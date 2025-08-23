---
marp: true
theme: default
paginate: true
---

# Security and Cryptography in Go
## Day 33: Advanced Security Patterns

---

# Overview

1. Cryptographic Fundamentals
2. Secure Communication
3. Authentication Patterns
4. Authorization Systems
5. Secret Management
6. Secure Coding Practices

---

# Cryptographic Operations

1. Symmetric Encryption
```go
func encrypt(key, plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    gcm, err := cipher.NewGCM(block)
    // ... encryption logic
}
```

2. Asymmetric Encryption
```go
func generateKeyPair() (*rsa.PrivateKey, error) {
    return rsa.GenerateKey(rand.Reader, 2048)
}
```

---

# Secure Communication

1. TLS Configuration
```go
config := &tls.Config{
    MinVersion: tls.VersionTLS12,
    CurvePreferences: []tls.CurveID{
        tls.X25519,
        tls.CurveP256,
    },
    CipherSuites: []uint16{
        tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
        tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
    },
}
```

2. Certificate Management
```go
func loadCertificate() (*tls.Certificate, error) {
    return tls.LoadX509KeyPair(
        "server.crt", "server.key")
}
```

---

# Authentication Patterns

1. JWT Implementation
```go
type Claims struct {
    UserID string `json:"uid"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

func generateToken(user User) (string, error) {
    claims := &Claims{
        UserID: user.ID,
        Role:   user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(24*time.Hour).Unix(),
        },
    }
    return jwt.NewWithClaims(
        jwt.SigningMethodHS256, claims).SignedString(secretKey)
}
```

---

# Password Security

1. Password Hashing
```go
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword(
        []byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
```

2. Password Validation
```go
func validatePassword(hash, password string) bool {
    err := bcrypt.CompareHashAndPassword(
        []byte(hash), []byte(password))
    return err == nil
}
```

---

# Authorization Systems

1. Role-Based Access Control (RBAC)
```go
type Permission struct {
    Resource string
    Action   string
}

type Role struct {
    Name        string
    Permissions []Permission
}

func checkPermission(user User, 
    resource, action string) bool {
    // Permission checking logic
}
```

---

# Secret Management

1. Vault Integration
```go
func getSecret(path string) (string, error) {
    client, err := vault.NewClient(nil)
    if err != nil {
        return "", err
    }
    
    secret, err := client.Logical().Read(path)
    // ... secret retrieval logic
}
```

2. Environment Security
```go
func loadSecrets() error {
    return godotenv.Load(".env.encrypted")
}
```

---

# Secure Headers

1. Security Headers
```go
func securityHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        next.ServeHTTP(w, r)
    })
}
```

---

# Input Validation

1. Request Validation
```go
func validateInput(input string) error {
    if len(input) > MaxInputLength {
        return ErrInputTooLong
    }
    return nil
}
```

2. SQL Injection Prevention
```go
func getUserSafely(db *sql.DB, id string) (*User, error) {
    var user User
    err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Name)
    return &user, err
}
```

---

# Rate Limiting

1. Token Bucket
```go
limiter := rate.NewLimiter(rate.Every(time.Second), 10)

func rateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Rate limit exceeded", 429)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---

# Secure File Operations

1. Safe File Handling
```go
func safeFileWrite(path string, data []byte) error {
    return ioutil.WriteFile(path, data, 0600)
}
```

2. Path Traversal Prevention
```go
func isValidPath(path string) bool {
    cleaned := filepath.Clean(path)
    return !strings.Contains(cleaned, "..")
}
```

---

# Best Practices

1. Always validate input
2. Use secure defaults
3. Implement defense in depth
4. Regular security audits
5. Keep dependencies updated
6. Use security headers

---

# Hands-on Exercise

Build a secure service with:
1. JWT authentication
2. RBAC authorization
3. Secure communication
4. Input validation
5. Rate limiting

---

# Key Takeaways

1. Security is multifaceted
2. Use proven libraries
3. Validate all input
4. Implement proper logging
5. Regular security reviews

---

# Next Steps

1. Advanced Encryption
2. Custom Auth Systems
3. Security Testing
4. Compliance Requirements
5. Threat Modeling
