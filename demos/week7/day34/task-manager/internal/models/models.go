package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"size:50;unique;not null" json:"username"`
	Email        string         `gorm:"size:100;unique;not null" json:"email"`
	PasswordHash string         `gorm:"size:255;not null" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:100;not null" json:"title"`
	Description string         `json:"description"`
	Priority    int            `gorm:"check:priority BETWEEN 1 AND 5" json:"priority"`
	Status      string         `gorm:"check:status IN ('pending', 'in_progress', 'completed', 'cancelled')" json:"status"`
	DueDate     *time.Time     `json:"due_date"`
	CategoryID  *uint          `json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedBy   uint           `json:"created_by"`
	Creator     User           `gorm:"foreignKey:CreatedBy" json:"creator"`
	AssignedTo  *uint          `json:"assigned_to"`
	Assignee    User           `gorm:"foreignKey:AssignedTo" json:"assignee"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
