---
marp: true
theme: default
paginate: true
---

# Database Integration in Go
## Week 5 - Day 21

---

# Today's Topics

1. SQL Database Basics
2. `database/sql` Package
3. Connection Management
4. CRUD Operations
5. Prepared Statements
6. Transactions

---

# Database Drivers

```go
import (
    "database/sql"
    _ "github.com/lib/pq"        // PostgreSQL
    _ "github.com/go-sql-driver/mysql" // MySQL
    _ "github.com/mattn/go-sqlite3"    // SQLite
)
```

The `_` means import for side effects only

---

# Connecting to Database

```go
// PostgreSQL
db, err := sql.Open("postgres", 
    "host=localhost port=5432 user=myuser password=mypass dbname=mydb sslmode=disable")

// MySQL
db, err := sql.Open("mysql", 
    "user:password@tcp(localhost:3306)/dbname")

// Check connection
if err = db.Ping(); err != nil {
    log.Fatal(err)
}
defer db.Close()
```

---

# Basic Queries

```go
// Query multiple rows
rows, err := db.Query("SELECT id, name FROM users")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string
    if err := rows.Scan(&id, &name); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("id: %d, name: %s\n", id, name)
}
```

---

# Single Row Query

```go
// QueryRow returns a single row
var name string
err := db.QueryRow("SELECT name FROM users WHERE id = $1", id).Scan(&name)

switch {
case err == sql.ErrNoRows:
    log.Println("No user found")
case err != nil:
    log.Fatal(err)
default:
    fmt.Printf("Name is %s\n", name)
}
```

---

# Modifying Data

```go
// Insert
result, err := db.Exec(
    "INSERT INTO users (name, email) VALUES ($1, $2)",
    "John Doe", "john@example.com")

// Get inserted ID
id, err := result.LastInsertId()

// Get affected rows
affected, err := result.RowsAffected()
```

---

# Prepared Statements

```go
stmt, err := db.Prepare("INSERT INTO users(name, email) VALUES($1, $2)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Execute multiple times
for _, user := range users {
    _, err := stmt.Exec(user.Name, user.Email)
    if err != nil {
        log.Printf("Error inserting user: %v", err)
    }
}
```

---

# Transactions

```go
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}

_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

err = tx.Commit()
if err != nil {
    log.Fatal(err)
}
```

---

# Connection Pool Settings

```go
db.SetMaxOpenConns(25)      // Maximum connections
db.SetMaxIdleConns(25)      // Idle connections
db.SetConnMaxLifetime(5 * time.Minute)
```

Best practices:
1. Set reasonable limits
2. Monitor pool metrics
3. Handle connection errors

---

# NULL Handling

```go
var name sql.NullString
var age sql.NullInt64

err := db.QueryRow("SELECT name, age FROM users WHERE id = $1", id).
    Scan(&name, &age)

if name.Valid {
    fmt.Printf("Name: %s\n", name.String)
} else {
    fmt.Println("Name is NULL")
}
```

---

# Error Handling

Common errors:
- `sql.ErrNoRows`
- Connection errors
- Constraint violations
- Transaction deadlocks

Always check for errors and handle them appropriately!

---

# Best Practices

1. Use connection pooling
2. Close resources (rows, statements)
3. Use prepared statements
4. Handle NULL values
5. Use transactions when needed
6. Validate input
7. Log important operations

---

# Exercise Time!

1. Create database schema
2. Implement CRUD operations
3. Use transactions
4. Handle errors properly
5. Add connection pooling

---

# Questions?

Let's practice with hands-on examples!
