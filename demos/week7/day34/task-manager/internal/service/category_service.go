package service

import (
	"errors"
	"task-manager/internal/models"
	"task-manager/internal/repository"
)

type CategoryService struct {
	categoryRepo *repository.CategoryRepository
	taskRepo     *repository.TaskRepository
}

func NewCategoryService(
	categoryRepo *repository.CategoryRepository,
	taskRepo *repository.TaskRepository,
) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
		taskRepo:     taskRepo,
	}
}

// CreateCategory creates a new category
func (s *CategoryService) CreateCategory(category *models.Category) error {
	// Check if category with same name exists
	existing, err := s.categoryRepo.GetByName(category.Name)
	if err == nil && existing != nil {
		return errors.New("category with this name already exists")
	}

	return s.categoryRepo.Create(category)
}

// GetCategory retrieves a category by ID
func (s *CategoryService) GetCategory(id uint) (*models.Category, error) {
	return s.categoryRepo.GetByID(id)
}

// UpdateCategory updates an existing category
func (s *CategoryService) UpdateCategory(category *models.Category) error {
	// Check if category exists
	existing, err := s.categoryRepo.GetByID(category.ID)
	if err != nil {
		return err
	}

	// Check if new name conflicts with another category
	if category.Name != existing.Name {
		if other, err := s.categoryRepo.GetByName(category.Name); err == nil && other != nil {
			return errors.New("category with this name already exists")
		}
	}

	return s.categoryRepo.Update(category)
}

// DeleteCategory deletes a category and handles associated tasks
func (s *CategoryService) DeleteCategory(id uint) error {
	// Get tasks in this category
	tasks, err := s.taskRepo.GetByCategory(id)
	if err != nil {
		return err
	}

	// If there are tasks in this category, prevent deletion
	if len(tasks) > 0 {
		return errors.New("cannot delete category with existing tasks")
	}

	return s.categoryRepo.Delete(id)
}

// ListCategories lists all categories
func (s *CategoryService) ListCategories() ([]models.Category, error) {
	return s.categoryRepo.List()
}

// GetCategoryByName gets a category by name
func (s *CategoryService) GetCategoryByName(name string) (*models.Category, error) {
	return s.categoryRepo.GetByName(name)
}
