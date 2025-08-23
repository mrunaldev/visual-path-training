package service

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repository"
	"task-manager/pkg/auth"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
	taskRepo *repository.TaskRepository
}

func NewUserService(
	userRepo *repository.UserRepository,
	taskRepo *repository.TaskRepository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		taskRepo: taskRepo,
	}
}

// RegisterUser registers a new user
func (s *UserService) RegisterUser(user *models.User, password string) error {
	// Check if username exists
	if existing, err := s.userRepo.GetByUsername(user.Username); err == nil && existing != nil {
		return errors.New("username already exists")
	}

	// Check if email exists
	if existing, err := s.userRepo.GetByEmail(user.Email); err == nil && existing != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	return s.userRepo.Create(user)
}

// AuthenticateUser authenticates a user and returns a JWT token
func (s *UserService) AuthenticateUser(username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *models.User) error {
	existing, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}

	// Check if new username conflicts with another user
	if user.Username != existing.Username {
		if other, err := s.userRepo.GetByUsername(user.Username); err == nil && other != nil {
			return errors.New("username already exists")
		}
	}

	// Check if new email conflicts with another user
	if user.Email != existing.Email {
		if other, err := s.userRepo.GetByEmail(user.Email); err == nil && other != nil {
			return errors.New("email already exists")
		}
	}

	return s.userRepo.Update(user)
}

// DeleteUser deletes a user and handles their tasks
func (s *UserService) DeleteUser(id uint) error {
	// Get tasks created by or assigned to this user
	tasks, err := s.taskRepo.GetByUserID(id)
	if err != nil {
		return err
	}

	// If there are tasks associated with this user, prevent deletion
	if len(tasks) > 0 {
		return errors.New("cannot delete user with existing tasks")
	}

	return s.userRepo.Delete(id)
}

// ListUsers lists all users
func (s *UserService) ListUsers() ([]models.User, error) {
	return s.userRepo.List()
}

// ChangePassword changes a user's password
func (s *UserService) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// Verify old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	return s.userRepo.Update(user)
}
