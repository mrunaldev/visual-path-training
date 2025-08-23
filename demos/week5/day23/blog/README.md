# MongoDB Blog Demo

This is a simple blog system demo showcasing MongoDB integration with Go. It demonstrates various MongoDB operations including CRUD operations, indexing, and aggregation.

## Prerequisites

1. MongoDB server running locally on default port (27017)
2. Go installed on your system

## Installation

1. Clone the repository
2. Install MongoDB driver:
   ```bash
   go get go.mongodb.org/mongo-driver/mongo
   ```
3. Run the demo:
   ```bash
   go run .
   ```

## Features Demonstrated

1. **Database Connection**
   - Connection to MongoDB
   - Database and collection management
   - Index creation

2. **User Management**
   - Create users with username and email
   - Username uniqueness enforced through indexes

3. **Blog Posts**
   - Create posts with title, content, and tags
   - Add comments to posts
   - Track view counts
   - Search posts by tags

4. **Aggregation**
   - Get post statistics by tags
   - Find popular posts
   - Calculate average views per tag

## Code Structure

- `main.go`: Entry point and database setup
- `models.go`: Data models and database structures
- `operations.go`: MongoDB CRUD operations
- `demo.go`: Demo runner showcasing features

## Sample Output

Running the demo will:

1. Create a test user
2. Create sample blog posts
3. Add comments
4. Simulate post views
5. Display posts filtered by tag
6. Show popular posts
7. Display post statistics

## Note

This is a demonstration project and not intended for production use. In a real application, you would want to:

- Add proper error handling
- Implement authentication and authorization
- Hash passwords
- Add input validation
- Implement proper session management
- Add proper logging
- Add tests
