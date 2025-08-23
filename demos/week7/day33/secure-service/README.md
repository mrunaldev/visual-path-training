# Secure Service Demo

This demo implements a secure service showcasing various security patterns and### Security Features

1. Authentication
   - JWT tokens with EdDSA (Ed25519) signatures
   - Password hashing with Argon2id
   - Role-based access with fine-grained permissions
   - MFA support (TOTP)

2. Encryption
   - ChaCha20-Poly1305 for symmetric encryption (better for resource-constrained devices)
   - Ed25519 for asymmetric operations
   - Quantum-resistant key exchange (optional)
   - Secure key generation with hardware entropyices in Go.

## Features

1. Cryptographic Operations
   - Symmetric encryption (AES-GCM)
   - Asymmetric encryption (RSA)
   - Key management
   - Base64 encoding/decoding

2. Authentication & Authorization
   - JWT-based authentication
   - Role-based access control
   - Password hashing with bcrypt
   - Permission management

3. Security Middleware
   - Rate limiting
   - Security headers
   - Request logging
   - Token validation

4. Best Practices
   - TLS configuration
   - Input validation
   - Error handling
   - Secure defaults

## Implementation Details

### Cryptography Module
- AES-GCM for symmetric encryption
- RSA for asymmetric encryption
- Secure random number generation
- Key management utilities

### Authentication Module
- JWT token generation and validation
- Password hashing with bcrypt
- User management
- Role-based permissions

### Security Headers
- X-Frame-Options
- X-Content-Type-Options
- X-XSS-Protection
- Content-Security-Policy
- HSTS

## Usage

### Building
```bash
go build -o secure-service ./cmd/server
```

### Running
```bash
./secure-service
```

### API Endpoints

1. Public Endpoints:
   ```
   POST /register
   POST /login
   ```

2. Protected Endpoints:
   ```
   POST /api/encrypt
   POST /api/decrypt
   ```

### Example Requests

1. Register User:
   ```bash
   curl -X POST http://localhost:8443/register \
     -H "Content-Type: application/json" \
     -d '{"username":"admin","password":"secret","role":"admin"}'
   ```

2. Login:
   ```bash
   curl -X POST http://localhost:8443/login \
     -H "Content-Type: application/json" \
     -d '{"username":"admin","password":"secret"}'
   ```

3. Encrypt Data:
   ```bash
   curl -X POST http://localhost:8443/api/encrypt \
     -H "Authorization: <token>" \
     -H "Content-Type: application/json" \
     -d '{"data":"secret message"}'
   ```

## Security Features

1. Authentication
   - JWT tokens
   - Password hashing
   - Role-based access

2. Encryption
   - AES-GCM (symmetric)
   - RSA (asymmetric)
   - Secure key generation

3. Protection
   - Rate limiting
   - Security headers
   - TLS
   - Input validation

## Best Practices Demonstrated

1. Security
   - Use of strong cryptography
   - Secure defaults
   - Proper error handling
   - Input validation

2. Authentication
   - JWT best practices
   - Password hashing
   - Token validation

3. Headers
   - Security headers
   - HSTS
   - CSP

4. Error Handling
   - No sensitive data in errors
   - Proper status codes
   - Logging security events

## Requirements

- Go 1.21+
- TLS certificates (cert.pem, key.pem)
- Required Go packages:
  - github.com/golang-jwt/jwt
  - github.com/gorilla/mux
  - golang.org/x/crypto
  - golang.org/x/time
