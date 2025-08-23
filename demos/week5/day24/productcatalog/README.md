# Product Catalog with Redis Caching

This demo shows how to implement caching patterns in Go using Redis, demonstrated through a product catalog system.

## Features

1. **Cache-Aside Pattern**
   - Check cache before database
   - Update cache on misses
   - Automatic cache expiration

2. **Write-Through Cache**
   - Update cache when writing to database
   - Maintain consistency

3. **Category-Based Caching**
   - Using Redis Sets for categories
   - Pipeline operations for bulk retrieval
   - Cache invalidation strategies

## Prerequisites

1. Redis server running locally on default port (6379)
2. Go installed on your system

## Installation

1. Install Redis client:
   ```bash
   go get github.com/redis/go-redis/v9
   ```

2. Run Redis server:
   ```bash
   redis-server
   ```

3. Run the demo:
   ```bash
   go run .
   ```

## Architecture

1. **CatalogService**
   - Main service coordinating cache and database operations
   - Implements caching patterns
   - Handles Redis connection

2. **ProductDB Interface**
   - Database abstraction
   - Allows different implementations
   - Mock implementation provided

3. **Redis Cache**
   - Product cache using string values
   - Category cache using sets
   - Pipeline operations for performance

## Caching Patterns Demonstrated

1. **Cache-Aside (Lazy Loading)**
   - Check cache first
   - Load from DB on miss
   - Update cache with new data

2. **Write-Through**
   - Write to DB
   - Update cache immediately
   - Maintain consistency

3. **Bulk Operations**
   - Pipeline commands
   - Batch updates
   - Efficient retrieval

## Best Practices Shown

1. **Error Handling**
   - Redis connection errors
   - Cache miss vs. error
   - Database fallback

2. **Performance**
   - Use of pipelines
   - Batch operations
   - Connection pooling

3. **Cache Management**
   - Automatic expiration
   - Consistent updates
   - Category-based invalidation

## Testing

```bash
# Start Redis
redis-server

# Run demo
go run .

# Monitor Redis
redis-cli monitor
```

## Note

This is a demonstration project and not intended for production use. In a real application, you would want to:

- Add proper error handling
- Implement retry mechanisms
- Add monitoring and metrics
- Handle cache consistency edge cases
- Add proper logging
- Add tests
- Implement cache warming
- Add cache eviction policies
