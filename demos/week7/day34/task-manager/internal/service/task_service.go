package service

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repository"
)

type TaskService struct {
	taskRepo     *repository.TaskRepository
	categoryRepo *repository.CategoryRepository
	userRepo     *repository.UserRepository
}

func NewTaskService(
	taskRepo *repository.TaskRepository,
	categoryRepo *repository.CategoryRepository,
	userRepo *repository.UserRepository,
) *TaskService {
	return &TaskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

// CreateTask creates a new task with validation
func (s *TaskService) CreateTask(task *models.Task) error {
	// Validate priority
	if task.Priority < 1 || task.Priority > 5 {
		return errors.New("priority must be between 1 and 5")
	}

	// Validate status
	validStatuses := map[string]bool{
		"pending":     true,
		"in_progress": true,
		"completed":   true,
		"cancelled":   true,
	}
	if !validStatuses[task.Status] {
		return errors.New("invalid task status")
	}

	// Validate category if provided
	if task.CategoryID != nil {
		if _, err := s.categoryRepo.GetByID(*task.CategoryID); err != nil {
			return errors.New("invalid category ID")
		}
	}

	// Validate assigned user if provided
	if task.AssignedTo != nil {
		if _, err := s.userRepo.GetByID(*task.AssignedTo); err != nil {
			return errors.New("invalid assigned user ID")
		}
	}

	// Create task
	return s.taskRepo.Create(task)
}

// GetTask retrieves a task by ID
func (s *TaskService) GetTask(id uint) (*models.Task, error) {
	return s.taskRepo.GetByID(id)
}

// UpdateTask updates an existing task with validation
func (s *TaskService) UpdateTask(task *models.Task) error {
	// Validate existing task
	existingTask, err := s.taskRepo.GetByID(task.ID)
	if err != nil {
		return err
	}

	// Validate priority
	if task.Priority < 1 || task.Priority > 5 {
		return errors.New("priority must be between 1 and 5")
	}

	// Validate status
	validStatuses := map[string]bool{
		"pending":     true,
		"in_progress": true,
		"completed":   true,
		"cancelled":   true,
	}
	if !validStatuses[task.Status] {
		return errors.New("invalid task status")
	}

	// Validate category if changed
	if task.CategoryID != nil && *task.CategoryID != *existingTask.CategoryID {
		if _, err := s.categoryRepo.GetByID(*task.CategoryID); err != nil {
			return errors.New("invalid category ID")
		}
	}

	// Validate assigned user if changed
	if task.AssignedTo != nil && *task.AssignedTo != *existingTask.AssignedTo {
		if _, err := s.userRepo.GetByID(*task.AssignedTo); err != nil {
			return errors.New("invalid assigned user ID")
		}
	}

	// Update task
	return s.taskRepo.Update(task)
}

// DeleteTask deletes a task by ID
func (s *TaskService) DeleteTask(id uint) error {
	return s.taskRepo.Delete(id)
}

// ListTasks lists all tasks with optional filters
func (s *TaskService) ListTasks(filters map[string]interface{}) ([]models.Task, error) {
	return s.taskRepo.List(filters)
}

// GetUserTasks gets all tasks associated with a user (either created by or assigned to)
func (s *TaskService) GetUserTasks(userID uint) ([]models.Task, error) {
	return s.taskRepo.GetByUserID(userID)
}

// GetCategoryTasks gets all tasks in a category
func (s *TaskService) GetCategoryTasks(categoryID uint) ([]models.Task, error) {
	return s.taskRepo.GetByCategory(categoryID)
}

// AssignTask assigns a task to a user
func (s *TaskService) AssignTask(taskID, userID uint) error {
	// Validate task exists
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return err
	}

	// Validate user exists
	if _, err := s.userRepo.GetByID(userID); err != nil {
		return errors.New("invalid user ID")
	}

	// Update task assignment
	task.AssignedTo = &userID
	return s.taskRepo.Update(task)
}

// UpdateTaskStatus updates the status of a task
func (s *TaskService) UpdateTaskStatus(taskID uint, status string) error {
	// Validate task exists
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return err
	}

	// Validate status
	validStatuses := map[string]bool{
		"pending":     true,
		"in_progress": true,
		"completed":   true,
		"cancelled":   true,
	}
	if !validStatuses[status] {
		return errors.New("invalid task status")
	}

	// Update task status
	task.Status = status
	return s.taskRepo.Update(task)
}
