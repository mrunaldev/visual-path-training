# Simple User API with Go

This is a demonstration of building a RESTful API with Go's `net/http` package, including static file serving and middleware.

## Features

- RESTful API endpoints for user management
- Static file serving
- Request logging middleware
- Concurrent-safe in-memory data store
- Simple web interface

## Project Structure

```
.
├── main.go         # Entry point and data store
├── server.go       # HTTP server implementation
└── static/         # Static files
    ├── index.html  # Web interface
    ├── styles.css  # CSS styles
    └── script.js   # Frontend JavaScript
```

## API Endpoints

- `GET /users` - List all users
- `POST /users` - Create a new user
- `GET /users/{id}` - Get a specific user
- `PUT /users/{id}` - Update a user
- `DELETE /users/{id}` - Delete a user

## Running the Server

```bash
go run *.go
```

Then visit http://localhost:8080 in your browser.

## Example API Usage

### Create a User
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","age":30}'
```

### Get All Users
```bash
curl http://localhost:8080/users
```

### Get a Specific User
```bash
curl http://localhost:8080/users/1
```

### Update a User
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Smith","age":31}'
```

### Delete a User
```bash
curl -X DELETE http://localhost:8080/users/1
```

## Key Concepts Demonstrated

1. HTTP server setup
2. Route handling
3. Middleware implementation
4. JSON request/response handling
5. Static file serving
6. Concurrent data access
7. RESTful API design
8. Error handling

## Learning Objectives

1. Understanding HTTP server basics in Go
2. Working with `net/http` package
3. Implementing REST APIs
4. Using middleware for cross-cutting concerns
5. Managing concurrent data access
6. Serving static files
7. Frontend integration
