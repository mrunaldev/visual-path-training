package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Base model for common fields
type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// User represents a customer
type User struct {
	Base
	Name    string     `gorm:"size:255;not null"`
	Email   string     `gorm:"size:255;not null;uniqueIndex"`
	Address string     `gorm:"type:text"`
	Orders  []Order    `gorm:"foreignKey:UserID"`
	Reviews []Review   `gorm:"foreignKey:UserID"`
	Cart    []CartItem `gorm:"foreignKey:UserID"`
}

// Product represents items available for sale
type Product struct {
	Base
	Name        string     `gorm:"size:255;not null"`
	Description string     `gorm:"type:text"`
	Price       float64    `gorm:"type:decimal(10,2);not null"`
	Stock       int        `gorm:"not null"`
	Categories  []Category `gorm:"many2many:product_categories;"`
	Reviews     []Review   `gorm:"foreignKey:ProductID"`
}

// Category represents product categories
type Category struct {
	Base
	Name        string    `gorm:"size:255;not null;uniqueIndex"`
	Description string    `gorm:"type:text"`
	Products    []Product `gorm:"many2many:product_categories;"`
}

// Order represents a customer order
type Order struct {
	Base
	UserID     uint        `gorm:"not null"`
	User       User        `gorm:"foreignKey:UserID"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
	Status     string      `gorm:"size:50;not null;default:'pending'"`
	TotalPrice float64     `gorm:"type:decimal(10,2);not null"`
}

// OrderItem represents items in an order
type OrderItem struct {
	Base
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"type:decimal(10,2);not null"`
}

// CartItem represents items in a user's shopping cart
type CartItem struct {
	Base
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
}

// Review represents a product review
type Review struct {
	Base
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Rating    int     `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment   string  `gorm:"type:text"`
}

// Hooks example
func (u *User) BeforeCreate(tx *gorm.DB) error {
	log.Printf("Creating new user: %s\n", u.Name)
	return nil
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	// Calculate total price
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	o.TotalPrice = total
	return nil
}

func (o *Order) AfterCreate(tx *gorm.DB) error {
	log.Printf("New order created with total price: $%.2f\n", o.TotalPrice)
	return nil
}

func main() {
	// Database connection
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	err = db.AutoMigrate(
		&User{},
		&Product{},
		&Category{},
		&Order{},
		&OrderItem{},
		&CartItem{},
		&Review{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Run demo
	if err := runDemo(db); err != nil {
		log.Fatal("Demo failed:", err)
	}

	fmt.Println("Demo completed successfully!")
}
