---
marp: true
theme: default
paginate: true
---

# Caching with Redis in Go
## Advanced Go Programming - Day 24

---

# Overview

1. Introduction to Redis
2. Redis Data Types
3. Go-Redis Library
4. Common Caching Patterns
5. Best Practices

---

# Introduction to Redis

- In-memory data structure store
- Key-value database
- Cache and message broker
- Supports rich data structures

Key Features:
- Fast (in-memory)
- Persistent
- Supports replication
- Cluster mode available

---

# Redis Data Types

1. Strings
   ```redis
   SET key value
   GET key
   ```

2. Lists
   ```redis
   LPUSH mylist value
   RPOP mylist
   ```

3. Sets
   ```redis
   SADD myset value
   SMEMBERS myset
   ```

4. Hashes
   ```redis
   HSET user:1 name "John" age "30"
   HGET user:1 name
   ```

---

# Go-Redis Library

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

// String operations
err := rdb.Set(ctx, "key", "value", 0).Err()
val, err := rdb.Get(ctx, "key").Result()
```

---

# String Operations

```go
// Set with expiration
err := rdb.Set(ctx, "session:123", "data", time.Hour).Err()

// Get with default value
val := rdb.Get(ctx, "missing_key").Val()

// Increment
newVal, err := rdb.Incr(ctx, "counter").Result()

// Multi-get
vals, err := rdb.MGet(ctx, "key1", "key2").Result()
```

---

# List Operations

```go
// Push to list
err := rdb.LPush(ctx, "queue", "task1").Err()
err = rdb.RPush(ctx, "queue", "task2").Err()

// Pop from list
val, err := rdb.LPop(ctx, "queue").Result()

// Range of list
vals, err := rdb.LRange(ctx, "queue", 0, -1).Result()

// Block until item available
val, err := rdb.BLPop(ctx, 0, "queue").Result()
```

---

# Hash Operations

```go
// Set hash fields
err := rdb.HSet(ctx, "user:1", map[string]interface{}{
    "name": "John",
    "age":  30,
}).Err()

// Get single field
name, err := rdb.HGet(ctx, "user:1", "name").Result()

// Get all fields
fields, err := rdb.HGetAll(ctx, "user:1").Result()
```

---

# Set Operations

```go
// Add to set
err := rdb.SAdd(ctx, "online_users", "user1").Err()

// Check membership
exists, err := rdb.SIsMember(ctx, "online_users", "user1").Result()

// Get all members
members, err := rdb.SMembers(ctx, "online_users").Result()

// Set operations
common, err := rdb.SInter(ctx, "set1", "set2").Result()
```

---

# Common Caching Patterns

1. Cache-Aside (Lazy Loading)
```go
func GetUser(id string) (*User, error) {
    // Try cache first
    user, err := getUserFromCache(id)
    if err == nil {
        return user, nil
    }

    // Cache miss - get from DB
    user, err = getUserFromDB(id)
    if err != nil {
        return nil, err
    }

    // Store in cache
    cacheUser(id, user)
    return user, nil
}
```

---

# Write-Through Cache

```go
func SaveUser(user *User) error {
    // Save to DB first
    if err := saveUserToDB(user); err != nil {
        return err
    }

    // Then update cache
    return cacheUser(user.ID, user)
}

func cacheUser(id string, user *User) error {
    return rdb.HSet(ctx, "user:"+id, map[string]interface{}{
        "name":  user.Name,
        "email": user.Email,
    }).Err()
}
```

---

# Pipeline Operations

```go
pipe := rdb.Pipeline()

incr := pipe.Incr(ctx, "counter")
pipe.Expire(ctx, "counter", time.Hour)

// Execute pipeline
_, err := pipe.Exec(ctx)
if err != nil {
    panic(err)
}

// Get result
fmt.Println(incr.Val())
```

---

# Transactions

```go
txf := func(tx *redis.Tx) error {
    // Get current value
    n, err := tx.Get(ctx, "counter").Int()
    if err != nil && err != redis.Nil {
        return err
    }

    // Increment
    n++

    // Operation is committed only if the watched keys remain unchanged
    _, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
        pipe.Set(ctx, "counter", n, 0)
        return nil
    })
    return err
}

// Retry if key changed
for {
    err := rdb.Watch(ctx, txf, "counter")
    if err != redis.TxFailedErr {
        break
    }
}
```

---

# Error Handling

```go
val, err := rdb.Get(ctx, "key").Result()
switch {
case err == redis.Nil:
    fmt.Println("key does not exist")
case err != nil:
    fmt.Println("error:", err)
case val == "":
    fmt.Println("value is empty")
default:
    fmt.Println("value:", val)
}
```

---

# Best Practices

1. Connection Management
   - Use connection pools
   - Monitor pool stats
   - Handle reconnection

2. Error Handling
   - Handle redis.Nil
   - Use timeouts
   - Implement retries

3. Data Structure Choice
   - Use appropriate types
   - Consider memory usage
   - Plan key expiration

4. Performance
   - Use pipelining
   - Batch operations
   - Monitor metrics

---

# Common Use Cases

1. Session Storage
2. Rate Limiting
3. Caching
4. Leaderboards
5. Real-time Analytics
6. Job Queues
7. Pub/Sub

---

# Questions?

1. When to use Redis?
2. How to handle cache invalidation?
3. Best practices for key naming?
4. How to implement rate limiting?
