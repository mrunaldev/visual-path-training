# E-commerce System with GORM

This is a demonstration of using GORM (Go Object Relational Mapper) to build a simple e-commerce system with PostgreSQL.

## Features

- User management
- Product catalog with categories
- Shopping cart functionality
- Order processing
- Product reviews
- Relationships (One-to-One, One-to-Many, Many-to-Many)
- Transaction handling
- Hooks and callbacks

## Project Structure

```
.
├── main.go    # Models and database setup
└── demo.go    # Demo functionality
```

## Models

1. `User`
   - Basic user information
   - Has many Orders
   - Has many Reviews
   - Has many CartItems

2. `Product`
   - Product information
   - Belongs to many Categories
   - Has many Reviews

3. `Category`
   - Category information
   - Has many Products (many-to-many)

4. `Order`
   - Order information
   - Belongs to User
   - Has many OrderItems

5. `OrderItem`
   - Individual items in an order
   - Belongs to Order
   - Belongs to Product

6. `CartItem`
   - Shopping cart items
   - Belongs to User
   - Belongs to Product

7. `Review`
   - Product reviews
   - Belongs to User
   - Belongs to Product

## Setup

1. Install PostgreSQL if not already installed

2. Create a new database and user:
   ```sql
   CREATE DATABASE gorm;
   CREATE USER gorm WITH PASSWORD 'gorm';
   GRANT ALL PRIVILEGES ON DATABASE gorm TO gorm;
   ```

3. Install dependencies:
   ```bash
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/postgres
   ```

4. Run the program:
   ```bash
   go run .
   ```

## Key Concepts Demonstrated

1. Model definition with GORM tags
2. Relationships
   - Has One
   - Has Many
   - Belongs To
   - Many to Many
3. Auto-migration
4. CRUD operations
5. Preloading relationships
6. Transaction handling
7. Hooks and callbacks

## Example Operations

The demo shows:
1. Creating categories and products
2. Adding users
3. Shopping cart operations
4. Order processing with transactions
5. Adding product reviews
6. Complex queries with preloading

## Best Practices Implemented

1. Using base model for common fields
2. Proper relationship definitions
3. Transaction handling for order processing
4. Input validation through struct tags
5. Soft deletes
6. Logging and error handling
7. Database connection management

## Learning Objectives

1. Understanding GORM basics
2. Working with relationships
3. Managing transactions
4. Using hooks and callbacks
5. Performing complex queries
6. Handling database migrations
7. Following ORM best practices
