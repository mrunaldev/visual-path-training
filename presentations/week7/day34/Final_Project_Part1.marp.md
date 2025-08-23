## Day 34: Final Project (Part 1)

Today we'll start building our final project: a RESTful Task Management API with advanced features.

### Project Overview

1. Features:
   - User authentication
   - Task CRUD operations
   - Task categories
   - Priority levels
   - Due dates
   - Task assignments
   - Task status tracking

2. Technical Stack:
   - Go (latest version)
   - GORM for ORM
   - PostgreSQL for database
   - JWT for authentication
   - Docker for containerization

### Database Design

1. Users Table:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

2. Categories Table:
```sql
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

3. Tasks Table:
```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    priority INT CHECK (priority BETWEEN 1 AND 5),
    status VARCHAR(20) CHECK (status IN ('pending', 'in_progress', 'completed', 'cancelled')),
    due_date TIMESTAMP,
    category_id INT REFERENCES categories(id),
    created_by INT REFERENCES users(id),
    assigned_to INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Implementation Steps

1. Project Setup:
   - Initialize Go module
   - Set up project structure
   - Add dependencies
   - Configure environment

2. Database Layer:
   - Create database models
   - Implement GORM setup
   - Create database migrations
   - Write repository layer

3. Basic Operations:
   - CRUD operations for tasks
   - Category management
   - User operations
   - Error handling

### Code Structure

```
task-manager/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── database.go
│   │   └── migrations.go
│   ├── models/
│   │   ├── user.go
│   │   ├── task.go
│   │   └── category.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── task_repository.go
│   │   └── category_repository.go
│   └── service/
│       ├── user_service.go
│       ├── task_service.go
│       └── category_service.go
├── pkg/
│   ├── auth/
│   │   └── jwt.go
│   └── validator/
│       └── validator.go
└── go.mod
```

### Best Practices

1. Database:
   - Use transactions for complex operations
   - Implement soft deletes
   - Handle database errors gracefully
   - Use indexes for better performance

2. Code Organization:
   - Follow clean architecture principles
   - Separate business logic from infrastructure
   - Use dependency injection
   - Write maintainable code

3. Error Handling:
   - Create custom error types
   - Implement proper logging
   - Use meaningful error messages
   - Handle edge cases

4. Testing:
   - Write unit tests
   - Use mocks for external dependencies
   - Implement integration tests
   - Test error scenarios
