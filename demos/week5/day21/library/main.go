package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Book represents a book in the library
type Book struct {
	ID        int
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// Member represents a library member
type Member struct {
	ID       int
	Name     string
	Email    string
	JoinDate time.Time
	Active   bool
}

// Borrowing represents a book borrowing record
type Borrowing struct {
	ID         int
	BookID     int
	MemberID   int
	BorrowDate time.Time
	DueDate    time.Time
	ReturnDate sql.NullTime
}

// DB wraps the database connection
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func main() {
	// Connection parameters
	connStr := "host=localhost port=5432 user=library_user password=library_pass dbname=library sslmode=disable"

	// Create database connection
	db, err := NewDB(connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Initialize schema
	if err := initSchema(db); err != nil {
		log.Fatal("Error initializing schema:", err)
	}

	// Example usage
	err = demoLibrary(db)
	if err != nil {
		log.Fatal("Error running demo:", err)
	}
}

func initSchema(db *DB) error {
	// Create tables if they don't exist
	queries := []string{
		`CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			isbn VARCHAR(13) UNIQUE NOT NULL,
			available BOOLEAN DEFAULT true
		)`,
		`CREATE TABLE IF NOT EXISTS members (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			join_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			active BOOLEAN DEFAULT true
		)`,
		`CREATE TABLE IF NOT EXISTS borrowings (
			id SERIAL PRIMARY KEY,
			book_id INTEGER REFERENCES books(id),
			member_id INTEGER REFERENCES members(id),
			borrow_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			due_date TIMESTAMP NOT NULL,
			return_date TIMESTAMP
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error creating schema: %v", err)
		}
	}

	return nil
}

func demoLibrary(db *DB) error {
	// Add some sample books
	book1, err := db.AddBook("The Go Programming Language", "Alan A. A. Donovan, Brian W. Kernighan", "9780134190440")
	if err != nil {
		return err
	}
	fmt.Printf("Added book: %s\n", book1.Title)

	book2, err := db.AddBook("Clean Code", "Robert C. Martin", "9780132350884")
	if err != nil {
		return err
	}
	fmt.Printf("Added book: %s\n", book2.Title)

	// Add some members
	member1, err := db.AddMember("John Doe", "john@example.com")
	if err != nil {
		return err
	}
	fmt.Printf("Added member: %s\n", member1.Name)

	member2, err := db.AddMember("Jane Smith", "jane@example.com")
	if err != nil {
		return err
	}
	fmt.Printf("Added member: %s\n", member2.Name)

	// Borrow books
	fmt.Printf("\nBorrowing books...\n")
	err = db.BorrowBook(book1.ID, member1.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Book '%s' borrowed by %s\n", book1.Title, member1.Name)

	err = db.BorrowBook(book2.ID, member2.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Book '%s' borrowed by %s\n", book2.Title, member2.Name)

	// List all books
	fmt.Printf("\nListing all books:\n")
	books, err := db.ListBooks()
	if err != nil {
		return err
	}
	for _, book := range books {
		fmt.Printf("- %s by %s (Available: %v)\n", book.Title, book.Author, book.Available)
	}

	// Return a book
	fmt.Printf("\nReturning books...\n")
	err = db.ReturnBook(book1.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Book '%s' returned\n", book1.Title)

	// Check member borrowings
	fmt.Printf("\nChecking member borrowings...\n")
	borrowings, err := db.GetMemberBorrowings(member1.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Member %s has %d borrowing records\n", member1.Name, len(borrowings))

	// Check overdue books
	fmt.Printf("\nChecking overdue books...\n")
	overdue, err := db.GetOverdueBorrowings()
	if err != nil {
		return err
	}
	fmt.Printf("Found %d overdue books\n", len(overdue))

	// Search for books
	fmt.Printf("\nSearching for books with 'Go'...\n")
	searchResults, err := db.SearchBooks("Go")
	if err != nil {
		return err
	}
	for _, book := range searchResults {
		fmt.Printf("- %s by %s\n", book.Title, book.Author)
	}

	return nil
}
