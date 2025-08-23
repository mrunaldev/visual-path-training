---
marp: true
theme: default
paginate: true
---

# ORM with GORM
## Week 5 - Day 22

---

# Today's Topics

1. What is ORM?
2. GORM Basics
3. Model Definition
4. CRUD Operations
5. Relationships
6. Hooks & Callbacks

---

# What is ORM?

Object-Relational Mapping:
- Maps database tables to structs
- Handles data conversion
- Manages relationships
- Simplifies database operations

Benefits:
- Type safety
- Less boilerplate
- Consistent API
- Easier migrations

---

# GORM Setup

```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

Configuration options:
```go
config := &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
    NowFunc: func() time.Time { return time.Now().UTC() },
}
```

---

# Model Definition

```go
type User struct {
    gorm.Model               // Adds ID, CreatedAt, UpdatedAt, DeletedAt
    Name      string
    Age       uint
    Email     string        `gorm:"uniqueIndex"`
    Birthday  *time.Time
    Profile   Profile       // Has One relationship
    Orders    []Order       // Has Many relationship
}

type Profile struct {
    ID     uint
    UserID uint           // Foreign key
    Bio    string
}
```

---

# Basic CRUD

```go
// Create
user := User{Name: "John", Age: 30}
db.Create(&user)

// Read
var user User
db.First(&user, 1)                // Find by primary key
db.First(&user, "name = ?", "John") // Find by condition

// Update
db.Model(&user).Update("name", "Jane")
db.Model(&user).Updates(User{Name: "Jane", Age: 31})

// Delete
db.Delete(&user)
```

---

# Query Building

```go
// Complex queries
var users []User
db.Where("age > ?", 20).
   Where("name LIKE ?", "%Jin%").
   Order("age desc, name").
   Limit(10).
   Find(&users)

// Joins
db.Joins("Profile").
   Joins("Orders").
   Find(&users)

// Count
var count int64
db.Model(&User{}).Where("age > ?", 20).Count(&count)
```

---

# Relationships

```go
// Has One
type User struct {
    ID      uint
    Profile Profile
}

// Has Many
type User struct {
    ID     uint
    Orders []Order
}

// Belongs To
type Order struct {
    ID     uint
    UserID uint
    User   User
}

// Many To Many
type User struct {
    ID    uint
    Roles []Role `gorm:"many2many:user_roles;"`
}
```

---

# Loading Relationships

```go
// Eager loading
db.Preload("Orders").Find(&users)

// Nested preloading
db.Preload("Orders.Items").Find(&users)

// Conditional preloading
db.Preload("Orders", "state = ?", "paid").Find(&users)

// Lazy loading
var user User
db.First(&user)
db.Model(&user).Association("Orders").Find(&orders)
```

---

# Hooks & Callbacks

```go
func (u *User) BeforeCreate(tx *gorm.DB) error {
    u.UUID = uuid.New()
    return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
    return tx.Model(&u).Update("processed", true).Error
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
    return tx.Model(&u).Update("deleted", true).Error
}
```

---

# Migrations

```go
// Auto migration
db.AutoMigrate(&User{}, &Profile{}, &Order{})

// Manual migrations
type Migration struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string
    Timestamp time.Time
}

func Migrate(db *gorm.DB) error {
    return db.Transaction(func(tx *gorm.DB) error {
        // Your migration logic here
        return nil
    })
}
```

---

# Transactions

```go
err := db.Transaction(func(tx *gorm.DB) error {
    // Create user
    if err := tx.Create(&user).Error; err != nil {
        return err
    }

    // Create profile
    if err := tx.Create(&profile).Error; err != nil {
        return err
    }

    return nil
})
```

---

# Best Practices

1. Use appropriate indexes
2. Implement soft deletes
3. Handle transactions properly
4. Use preloading wisely
5. Monitor query performance
6. Implement proper error handling
7. Use migrations

---

# Exercise Time!

1. Create models with relationships
2. Implement CRUD operations
3. Use transactions
4. Add hooks and callbacks
5. Write migrations

---

# Questions?

Let's practice with hands-on examples!
