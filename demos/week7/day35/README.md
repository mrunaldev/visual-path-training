# Day 35: Task Manager API Implementation

This is the second part of our final project, focusing on implementing the REST API endpoints and deploying the application using Docker.

## API Documentation

### Authentication Endpoints

#### Register User
```http
POST /api/register
Content-Type: application/json

{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secure_password"
}
```

#### Login
```http
POST /api/login
Content-Type: application/json

{
    "username": "john_doe",
    "password": "secure_password"
}
```

### Task Endpoints

#### Create Task
```http
POST /api/tasks
Authorization: Bearer <token>
Content-Type: application/json

{
    "title": "Implement API",
    "description": "Implement REST API endpoints",
    "priority": 1,
    "status": "pending",
    "category_id": 1,
    "due_date": "2024-01-01T00:00:00Z"
}
```

#### Get Task
```http
GET /api/tasks/:id
Authorization: Bearer <token>
```

#### Update Task
```http
PUT /api/tasks/:id
Authorization: Bearer <token>
Content-Type: application/json

{
    "title": "Updated Task",
    "description": "Updated description",
    "priority": 2,
    "status": "in_progress"
}
```

#### Delete Task
```http
DELETE /api/tasks/:id
Authorization: Bearer <token>
```

### Category Endpoints

#### Create Category
```http
POST /api/categories
Authorization: Bearer <token>
Content-Type: application/json

{
    "name": "Backend Development",
    "description": "Backend development tasks"
}
```

#### Get Category
```http
GET /api/categories/:id
Authorization: Bearer <token>
```

#### Update Category
```http
PUT /api/categories/:id
Authorization: Bearer <token>
Content-Type: application/json

{
    "name": "Updated Category",
    "description": "Updated description"
}
```

#### Delete Category
```http
DELETE /api/categories/:id
Authorization: Bearer <token>
```

## Deployment

### Prerequisites
- Docker
- Docker Compose
- Go 1.21 or later

### Local Development
1. Clone the repository
2. Copy `.env.example` to `.env`
3. Start PostgreSQL:
   ```bash
   docker-compose up db
   ```
4. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

### Docker Deployment
1. Build and run containers:
   ```bash
   docker-compose up --build
   ```
2. The API will be available at http://localhost:8080

### Production Deployment
1. Update environment variables
2. Build optimized image:
   ```bash
   docker build -t task-manager:prod .
   ```
3. Run with production settings:
   ```bash
   docker-compose -f docker-compose.prod.yml up -d
   ```

## Testing

### Running Tests
```bash
go test ./...
```

### Load Testing
Using k6:
```bash
k6 run load_tests/api_test.js
```

## Monitoring

### Health Check
```http
GET /health
```

### Metrics
```http
GET /metrics
```

## Security Considerations
1. Use HTTPS in production
2. Implement rate limiting
3. Use secure headers
4. Regular security updates
5. Proper error handling
6. Input validation
7. Audit logging

## Future Enhancements
1. WebSocket support
2. File attachments
3. Email notifications
4. Advanced search
5. Activity tracking
6. User roles/permissions
7. API versioning
8. Cache layer
