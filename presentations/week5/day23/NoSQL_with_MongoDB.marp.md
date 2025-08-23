---
marp: true
theme: default
paginate: true
---

# NoSQL with MongoDB
## Week 5 - Day 23

---

# Today's Topics

1. NoSQL Concepts
2. MongoDB Basics
3. MongoDB Driver
4. CRUD Operations
5. Document Relations
6. Aggregation Pipeline

---

# NoSQL vs SQL

SQL (Relational):
- Fixed schema
- Structured data
- ACID transactions
- Vertical scaling

NoSQL (Document):
- Flexible schema
- Semi-structured data
- Eventually consistent
- Horizontal scaling

---

# MongoDB Concepts

```javascript
// Database
use mydb

// Collection (like SQL table)
db.users

// Document (like SQL row)
{
    "_id": ObjectId("..."),
    "name": "John",
    "age": 30,
    "addresses": [
        { "city": "New York", "type": "home" }
    ]
}
```

---

# MongoDB Driver Setup

```go
import "go.mongodb.org/mongo-driver/mongo"

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
    "mongodb://localhost:27017",
))

// Disconnect
defer client.Disconnect(context.TODO())

// Get database and collection
db := client.Database("mydb")
collection := db.Collection("users")
```

---

# Basic CRUD Operations

```go
// Insert One
result, err := collection.InsertOne(ctx, bson.D{
    {"name", "John"},
    {"age", 30},
})

// Insert Many
docs := []interface{}{
    bson.D{{"name", "Alice"}, {"age", 25}},
    bson.D{{"name", "Bob"}, {"age", 35}},
}
results, err := collection.InsertMany(ctx, docs)
```

---

# Finding Documents

```go
// Find One
var user bson.M
err := collection.FindOne(ctx, bson.D{
    {"name", "John"},
}).Decode(&user)

// Find Many
cursor, err := collection.Find(ctx, bson.D{
    {"age", bson.D{{"$gt", 25}}},
})
defer cursor.Close(ctx)

for cursor.Next(ctx) {
    var user bson.M
    cursor.Decode(&user)
    // Process user...
}
```

---

# Updating Documents

```go
// Update One
update := bson.D{{"$set", bson.D{
    {"age", 31},
}}}
result, err := collection.UpdateOne(ctx,
    bson.D{{"name", "John"}},
    update,
)

// Update Many
result, err = collection.UpdateMany(ctx,
    bson.D{{"age", bson.D{{"$lt", 30}}}},
    bson.D{{"$inc", bson.D{{"age", 1}}}},
)
```

---

# Deleting Documents

```go
// Delete One
result, err := collection.DeleteOne(ctx,
    bson.D{{"name", "John"}},
)

// Delete Many
result, err = collection.DeleteMany(ctx,
    bson.D{{"age", bson.D{{"$lt", 25}}}},
)

// Drop Collection
err = collection.Drop(ctx)
```

---

# Document Relations

```go
// Embedded Documents
type Address struct {
    Street string `bson:"street"`
    City   string `bson:"city"`
}

type User struct {
    Name      string    `bson:"name"`
    Addresses []Address `bson:"addresses"`
}

// References
type Order struct {
    UserID primitive.ObjectID `bson:"userId"`
    Items  []string          `bson:"items"`
}
```

---

# Using Indexes

```go
// Create Index
model := mongo.IndexModel{
    Keys: bson.D{{"name", 1}},
    Options: options.Index().
        SetUnique(true),
}
name, err := collection.Indexes().
    CreateOne(ctx, model)

// List Indexes
cursor, err := collection.Indexes().
    List(ctx)
```

---

# Aggregation Pipeline

```go
pipeline := mongo.Pipeline{
    {{{"$match", bson.D{
        {"age", bson.D{{"$gt", 25}}},
    }}}},
    {{{"$group", bson.D{
        {"_id", "$city"},
        {"count", bson.D{{"$sum", 1}}},
        {"avgAge", bson.D{{"$avg", "$age"}}},
    }}}},
    {{{"$sort", bson.D{
        {"count", -1},
    }}}},
}

cursor, err := collection.Aggregate(ctx, pipeline)
```

---

# Transactions

```go
session, err := client.StartSession()
defer session.EndSession(ctx)

err = session.StartTransaction()

err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
    // Perform operations...
    if err != nil {
        session.AbortTransaction(sc)
        return err
    }
    return session.CommitTransaction(sc)
})
```

---

# Best Practices

1. Use proper data types
2. Index frequently queried fields
3. Limit embedded document size
4. Use bulk operations
5. Handle errors properly
6. Monitor performance
7. Implement retry logic

---

# Exercise Time!

1. Create a data model
2. Implement CRUD operations
3. Build complex queries
4. Use aggregation pipeline
5. Add indexes

---

# Questions?

Let's practice with hands-on examples!
