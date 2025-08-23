package main

import (
	"fmt"
	"log"
)

// This file contains example usage of the library management system
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

	// Run demo
	fmt.Println("Running Library Management System Demo...")
	if err := demoLibrary(db); err != nil {
		log.Fatal("Error running demo:", err)
	}
	fmt.Println("\nDemo completed successfully!")
}
