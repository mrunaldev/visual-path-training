# Task Manager API
## Day 34: Final Project (Part 1)

### Project Overview
- Task Management System
- RESTful API
- Full CRUD operations
- User authentication
- Role-based access

### Technical Stack
- Go 1.24+ (as of August 2025)
- GORM v2.0+
- PostgreSQL 16+
- JWT Authentication with up-to-date crypto standards
- Docker with multi-stage builds
- DevContainer support for VS Code

### Architecture
```
task-manager/
├── cmd/
│   └── server/          # Main application entry point
├── internal/            # Private application code
│   ├── config/         # Configuration
│   ├── database/       # Database setup
│   ├── models/         # Database models
│   ├── repository/     # Data access layer
│   └── service/        # Business logic
└── pkg/                # Public packages
    └── auth/           # Authentication utilities
```

### Key Features

1. Authentication
   - JWT-based
   - Secure password hashing
   - Token validation

2. Authorization
   - Role-based access
   - Permission checks
   - Secure routes

3. Task Management
   - Create/Read/Update/Delete tasks
   - Assign tasks to users
   - Set priorities and due dates
   - Track status changes

4. Category Management
   - Organize tasks by categories
   - Flexible categorization
   - Category-based filtering

### Code Architecture

1. Models Layer
   - Database schema representation
   - Model relationships
   - Field validation

2. Repository Layer
   - Data access abstraction
   - CRUD operations
   - Database queries

3. Service Layer
   - Business logic
   - Data validation
   - Cross-cutting concerns

4. API Layer (Coming in Part 2)
   - HTTP handlers
   - Request/Response handling
   - Middleware
