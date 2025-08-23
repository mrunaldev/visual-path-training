package repository

import (
	"task-manager/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.Preload("Category").Preload("Creator").Preload("Assignee").First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

func (r *TaskRepository) List(filters map[string]interface{}) ([]models.Task, error) {
	var tasks []models.Task
	query := r.db.Preload("Category").Preload("Creator").Preload("Assignee")

	for key, value := range filters {
		query = query.Where(key, value)
	}

	err := query.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Preload("Category").
		Preload("Creator").
		Preload("Assignee").
		Where("created_by = ? OR assigned_to = ?", userID, userID).
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetByCategory(categoryID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Preload("Category").
		Preload("Creator").
		Preload("Assignee").
		Where("category_id = ?", categoryID).
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
