package security

import (
	"errors"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid token")
)

// User represents a system user
type User struct {
	ID           string
	Username     string
	PasswordHash string
	Role         string
	Permissions  []string
}

// AuthManager handles authentication and authorization
type AuthManager struct {
	users     map[string]User
	mu        sync.RWMutex
	jwtSecret []byte
}

// NewAuthManager creates a new authentication manager
func NewAuthManager(jwtSecret []byte) *AuthManager {
	return &AuthManager{
		users:     make(map[string]User),
		jwtSecret: jwtSecret,
	}
}

// RegisterUser registers a new user
func (am *AuthManager) RegisterUser(username, password, role string) error {
	am.mu.Lock()
	defer am.mu.Unlock()

	// Check if user exists
	if _, exists := am.users[username]; exists {
		return ErrUserExists
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create user
	user := User{
		ID:           username, // In a real system, use UUID
		Username:     username,
		PasswordHash: string(hash),
		Role:         role,
		Permissions:  getPermissionsForRole(role),
	}

	am.users[username] = user
	return nil
}

// Authenticate verifies user credentials
func (am *AuthManager) Authenticate(username, password string) (string, error) {
	am.mu.RLock()
	user, exists := am.users[username]
	am.mu.RUnlock()

	if !exists {
		return "", ErrUserNotFound
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		return "", ErrInvalidCredentials
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(am.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken validates a JWT token
func (am *AuthManager) VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return am.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, ErrInvalidToken
}

// CheckPermission checks if a user has a specific permission
func (am *AuthManager) CheckPermission(userID, permission string) bool {
	am.mu.RLock()
	defer am.mu.RUnlock()

	user, exists := am.users[userID]
	if !exists {
		return false
	}

	for _, p := range user.Permissions {
		if p == permission {
			return true
		}
	}

	return false
}

// Helper functions

func getPermissionsForRole(role string) []string {
	switch role {
	case "admin":
		return []string{"read", "write", "delete", "admin"}
	case "user":
		return []string{"read", "write"}
	case "guest":
		return []string{"read"}
	default:
		return []string{}
	}
}
