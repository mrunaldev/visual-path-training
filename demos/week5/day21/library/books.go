package main

import (
	"database/sql"
	"fmt"
	"time"
)

// AddBook adds a new book to the library
func (db *DB) AddBook(title, author, isbn string) (*Book, error) {
	var book Book
	err := db.QueryRow(
		"INSERT INTO books (title, author, isbn) VALUES ($1, $2, $3) RETURNING id, title, author, isbn, available",
		title, author, isbn,
	).Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Available)

	if err != nil {
		return nil, fmt.Errorf("error adding book: %v", err)
	}
	return &book, nil
}

// GetBook retrieves a book by ID
func (db *DB) GetBook(id int) (*Book, error) {
	var book Book
	err := db.QueryRow(
		"SELECT id, title, author, isbn, available FROM books WHERE id = $1",
		id,
	).Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Available)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("book not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting book: %v", err)
	}
	return &book, nil
}

// UpdateBook updates a book's details
func (db *DB) UpdateBook(book *Book) error {
	result, err := db.Exec(
		"UPDATE books SET title = $1, author = $2, isbn = $3, available = $4 WHERE id = $5",
		book.Title, book.Author, book.ISBN, book.Available, book.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating book: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("book not found")
	}
	return nil
}

// DeleteBook removes a book from the library
func (db *DB) DeleteBook(id int) error {
	result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting book: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("book not found")
	}
	return nil
}

// ListBooks returns all books in the library
func (db *DB) ListBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author, isbn, available FROM books")
	if err != nil {
		return nil, fmt.Errorf("error listing books: %v", err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Available)
		if err != nil {
			return nil, fmt.Errorf("error scanning book row: %v", err)
		}
		books = append(books, book)
	}
	return books, nil
}

// SearchBooks searches for books by title or author
func (db *DB) SearchBooks(query string) ([]Book, error) {
	rows, err := db.Query(
		"SELECT id, title, author, isbn, available FROM books WHERE title ILIKE $1 OR author ILIKE $1",
		"%"+query+"%",
	)
	if err != nil {
		return nil, fmt.Errorf("error searching books: %v", err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Available)
		if err != nil {
			return nil, fmt.Errorf("error scanning book row: %v", err)
		}
		books = append(books, book)
	}
	return books, nil
}

// BorrowBook marks a book as borrowed by a member
func (db *DB) BorrowBook(bookID, memberID int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	// Check if book is available
	var available bool
	err = tx.QueryRow("SELECT available FROM books WHERE id = $1", bookID).Scan(&available)
	if err == sql.ErrNoRows {
		return fmt.Errorf("book not found")
	}
	if err != nil {
		return fmt.Errorf("error checking book availability: %v", err)
	}
	if !available {
		return fmt.Errorf("book is not available")
	}

	// Mark book as unavailable
	_, err = tx.Exec("UPDATE books SET available = false WHERE id = $1", bookID)
	if err != nil {
		return fmt.Errorf("error updating book availability: %v", err)
	}

	// Create borrowing record
	dueDate := time.Now().AddDate(0, 0, 14) // 2 weeks loan period
	_, err = tx.Exec(
		"INSERT INTO borrowings (book_id, member_id, due_date) VALUES ($1, $2, $3)",
		bookID, memberID, dueDate,
	)
	if err != nil {
		return fmt.Errorf("error creating borrowing record: %v", err)
	}

	return tx.Commit()
}

// ReturnBook marks a book as returned
func (db *DB) ReturnBook(bookID int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	// Update borrowing record
	result, err := tx.Exec(
		"UPDATE borrowings SET return_date = CURRENT_TIMESTAMP WHERE book_id = $1 AND return_date IS NULL",
		bookID,
	)
	if err != nil {
		return fmt.Errorf("error updating borrowing record: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no active borrowing found for this book")
	}

	// Mark book as available
	_, err = tx.Exec("UPDATE books SET available = true WHERE id = $1", bookID)
	if err != nil {
		return fmt.Errorf("error updating book availability: %v", err)
	}

	return tx.Commit()
}
