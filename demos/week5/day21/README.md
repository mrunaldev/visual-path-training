# Library Management System

This is a demonstration of using Go's `database/sql` package to build a simple library management system with PostgreSQL.

## Features

- Book management (add, update, delete, search)
- Member management (add, update, delete)
- Borrowing system with due dates
- Overdue book tracking
- Concurrent access support
- Transaction handling

## Project Structure

```
.
├── main.go      # Database connection and demo
├── books.go     # Book-related operations
└── members.go   # Member-related operations
```

## Database Schema

```sql
-- Books table
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(13) UNIQUE NOT NULL,
    available BOOLEAN DEFAULT true
);

-- Members table
CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    join_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    active BOOLEAN DEFAULT true
);

-- Borrowings table
CREATE TABLE borrowings (
    id SERIAL PRIMARY KEY,
    book_id INTEGER REFERENCES books(id),
    member_id INTEGER REFERENCES members(id),
    borrow_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP
);
```

## Setup

1. Install PostgreSQL if not already installed
2. Create a new database and user:
   ```sql
   CREATE DATABASE library;
   CREATE USER library_user WITH PASSWORD 'library_pass';
   GRANT ALL PRIVILEGES ON DATABASE library TO library_user;
   ```

3. Install dependencies:
   ```bash
   go get github.com/lib/pq
   ```

4. Run the program:
   ```bash
   go run .
   ```

## Key Concepts Demonstrated

1. Database connection management
2. Connection pooling
3. CRUD operations
4. Prepared statements
5. Transaction handling
6. Error handling
7. NULL handling with sql.NullTime
8. Query parameter safety

## Example Usage

The `main.go` file includes a demonstration that shows:
- Adding books and members
- Borrowing and returning books
- Listing all books
- Searching for books
- Checking member borrowings
- Finding overdue books

## Best Practices Implemented

1. Connection pooling with appropriate limits
2. Proper resource cleanup
3. Error handling and propagation
4. Use of prepared statements
5. Transaction management
6. Concurrent access safety
7. Parameterized queries for safety

## Learning Objectives

1. Working with `database/sql`
2. Understanding database connections
3. Managing transactions
4. Handling database errors
5. Building safe queries
6. Managing database resources
