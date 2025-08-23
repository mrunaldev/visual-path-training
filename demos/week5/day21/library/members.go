package main

import (
	"database/sql"
	"fmt"
	"time"
)

// AddMember adds a new member to the library
func (db *DB) AddMember(name, email string) (*Member, error) {
	var member Member
	err := db.QueryRow(
		"INSERT INTO members (name, email) VALUES ($1, $2) RETURNING id, name, email, join_date, active",
		name, email,
	).Scan(&member.ID, &member.Name, &member.Email, &member.JoinDate, &member.Active)

	if err != nil {
		return nil, fmt.Errorf("error adding member: %v", err)
	}
	return &member, nil
}

// GetMember retrieves a member by ID
func (db *DB) GetMember(id int) (*Member, error) {
	var member Member
	err := db.QueryRow(
		"SELECT id, name, email, join_date, active FROM members WHERE id = $1",
		id,
	).Scan(&member.ID, &member.Name, &member.Email, &member.JoinDate, &member.Active)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("member not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting member: %v", err)
	}
	return &member, nil
}

// UpdateMember updates a member's details
func (db *DB) UpdateMember(member *Member) error {
	result, err := db.Exec(
		"UPDATE members SET name = $1, email = $2, active = $3 WHERE id = $4",
		member.Name, member.Email, member.Active, member.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating member: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("member not found")
	}
	return nil
}

// DeleteMember removes a member from the library
func (db *DB) DeleteMember(id int) error {
	result, err := db.Exec("DELETE FROM members WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting member: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("member not found")
	}
	return nil
}

// ListMembers returns all members in the library
func (db *DB) ListMembers() ([]Member, error) {
	rows, err := db.Query("SELECT id, name, email, join_date, active FROM members")
	if err != nil {
		return nil, fmt.Errorf("error listing members: %v", err)
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.ID, &member.Name, &member.Email, &member.JoinDate, &member.Active)
		if err != nil {
			return nil, fmt.Errorf("error scanning member row: %v", err)
		}
		members = append(members, member)
	}
	return members, nil
}

// GetMemberBorrowings returns all books borrowed by a member
func (db *DB) GetMemberBorrowings(memberID int) ([]Borrowing, error) {
	rows, err := db.Query(`
		SELECT b.id, b.book_id, b.member_id, b.borrow_date, b.due_date, b.return_date 
		FROM borrowings b
		WHERE b.member_id = $1
		ORDER BY b.borrow_date DESC`,
		memberID,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting member borrowings: %v", err)
	}
	defer rows.Close()

	var borrowings []Borrowing
	for rows.Next() {
		var b Borrowing
		err := rows.Scan(&b.ID, &b.BookID, &b.MemberID, &b.BorrowDate, &b.DueDate, &b.ReturnDate)
		if err != nil {
			return nil, fmt.Errorf("error scanning borrowing row: %v", err)
		}
		borrowings = append(borrowings, b)
	}
	return borrowings, nil
}

// GetOverdueBorrowings returns all overdue books
func (db *DB) GetOverdueBorrowings() ([]Borrowing, error) {
	rows, err := db.Query(`
		SELECT b.id, b.book_id, b.member_id, b.borrow_date, b.due_date, b.return_date 
		FROM borrowings b
		WHERE b.return_date IS NULL AND b.due_date < $1
		ORDER BY b.due_date ASC`,
		time.Now(),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting overdue borrowings: %v", err)
	}
	defer rows.Close()

	var borrowings []Borrowing
	for rows.Next() {
		var b Borrowing
		err := rows.Scan(&b.ID, &b.BookID, &b.MemberID, &b.BorrowDate, &b.DueDate, &b.ReturnDate)
		if err != nil {
			return nil, fmt.Errorf("error scanning borrowing row: %v", err)
		}
		borrowings = append(borrowings, b)
	}
	return borrowings, nil
}
