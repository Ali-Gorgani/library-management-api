package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
	"library-management-api/books-service/init/database"
	"strconv"
)

var (
	ErrBookNotFound         = errors.New("book not found")
	ErrBookAlreadyBorrowed  = errors.New("book is already borrowed")
	ErrBookAlreadyAvailable = errors.New("book is already available")
	ErrBorrowerIDMismatch   = errors.New("borrower ID does not match")
	ErrInvalidCategoryType  = errors.New("invalid category type")
	ErrEmptyCategoryValue   = errors.New("category value cannot be empty")
	ErrInvalidSearchQuery   = errors.New("invalid search query")
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository() ports.BookRepository {
	return &BookRepository{
		db: database.P().DB,
	}
}

// Create implements ports.BookRepository.
func (b *BookRepository) AddBook(ctx context.Context, book *domain.AddBookReq) (*domain.BookRes, error) {
	var addedBook domain.BookRes

	query := "INSERT INTO books (title, author, category, subject, genre, published_year) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Category, book.Subject, book.Genre, book.PublishedYear)
	err := row.Scan(&addedBook.ID, &addedBook.Title, &addedBook.Author, &addedBook.Category, &addedBook.Subject, &addedBook.Genre, &addedBook.PublishedYear, &addedBook.Available, &addedBook.BorrowerID, &addedBook.CreatedAt)
	if err != nil {
		return &domain.BookRes{}, err
	}
	return &addedBook, nil
}

// GetBooks implements ports.BookRepository.
func (b *BookRepository) GetBooks(ctx context.Context) ([]*domain.BookRes, error) {
	var books []*domain.BookRes

	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := new(domain.BookRes)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// UpdateBook implements ports.BookRepository.
func (b *BookRepository) UpdateBook(ctx context.Context, id int, book *domain.UpdateBookReq) (*domain.BookRes, error) {
	var updatedBook domain.BookRes

	query := "UPDATE books SET title=$1, author=$2, category=$3, subject=$4, genre=$5, published_year=$6, available=$7, borrower_id=$8 WHERE id=$9 RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Category, book.Subject, book.Genre, book.PublishedYear, book.Available, book.BorrowerID, id)
	err := row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Category, &updatedBook.Subject, &updatedBook.Genre, &updatedBook.PublishedYear, &updatedBook.Available, &updatedBook.BorrowerID, &updatedBook.CreatedAt)
	if err != nil {
		return &domain.BookRes{}, err
	}
	return &updatedBook, nil
}

// DeleteBook implements ports.BookRepository.
func (b *BookRepository) DeleteBook(ctx context.Context, id int) error {
	query := "DELETE FROM books WHERE id=$1"
	_, err := b.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// BorrowBook implements ports.BookRepository.
func (b *BookRepository) BorrowBook(ctx context.Context, borrowBook *domain.BorrowBookReq) (*domain.BookRes, error) {
	var borrowedBook domain.BookRes
	// First, retrieve the book to check its availability
	query := "SELECT available FROM books WHERE id=$1"
	row := b.db.QueryRow(query, borrowBook.BookID)
	err := row.Scan(&borrowedBook.Available)
	if err != nil {
		if err == sql.ErrNoRows {
			return &domain.BookRes{}, errors.New("book not found")
		}
		return &domain.BookRes{}, err
	}

	// Check if the book is already borrowed
	if !borrowedBook.Available {
		return &domain.BookRes{}, errors.New("book is already borrowed")
	}

	// Update the book to mark it as borrowed and set the borrower ID
	query = "UPDATE books SET available=false, borrower_id=$1 WHERE id=$2 RETURNING *"
	row = b.db.QueryRow(query, borrowBook.BorrowerID, borrowBook.BookID)
	err = row.Scan(&borrowedBook.ID, &borrowedBook.Title, &borrowedBook.Author, &borrowedBook.Category, &borrowedBook.Subject, &borrowedBook.Genre, &borrowedBook.PublishedYear, &borrowedBook.Available, &borrowedBook.BorrowerID, &borrowedBook.CreatedAt)
	if err != nil {
		return &domain.BookRes{}, err
	}
	return &borrowedBook, nil
}

// ReturnBook implements ports.BookRepository.
func (b *BookRepository) ReturnBook(ctx context.Context, borrowBook *domain.ReturnBookReq) (*domain.BookRes, error) {
	var returnedBook domain.BookRes

	// First, retrieve the book to check its availability and borrower ID
	query := "SELECT available, borrower_id FROM books WHERE id=$1"
	row := b.db.QueryRow(query, borrowBook.BookID)
	err := row.Scan(&returnedBook.Available, &returnedBook.BorrowerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &domain.BookRes{}, errors.New("book not found")
		}
		return &domain.BookRes{}, err
	}

	// Check if the book is already available
	if returnedBook.Available {
		return &domain.BookRes{}, errors.New("book is already available")
	}

	// Check if the borrower ID matches
	if *returnedBook.BorrowerID != *borrowBook.BorrowerID {
		return &domain.BookRes{}, errors.New("borrower ID does not match")
	}

	// Update the book to mark it as available and clear the borrower ID
	query = "UPDATE books SET available=true, borrower_id=NULL WHERE id=$1 RETURNING *"
	row = b.db.QueryRow(query, borrowBook.BookID)
	err = row.Scan(&returnedBook.ID, &returnedBook.Title, &returnedBook.Author, &returnedBook.Category, &returnedBook.Subject, &returnedBook.Genre, &returnedBook.PublishedYear, &returnedBook.Available, &returnedBook.BorrowerID, &returnedBook.CreatedAt)
	if err != nil {
		return &domain.BookRes{}, err
	}
	return &returnedBook, nil
}

// SearchBooks implements ports.BookRepository.
func (b *BookRepository) SearchBooks(ctx context.Context, title string, author string, category string) ([]*domain.BookRes, error) {
	var books []*domain.BookRes

	// Build the SQL query dynamically based on which parameters are provided
	query := "SELECT id, title, author, category, subject, genre, published_year, available, borrower_id, created_at FROM books WHERE 1=1"
	var args []interface{}

	if title != "" {
		query += " AND title ILIKE $1"     // ILIKE for case-insensitive search
		args = append(args, "%"+title+"%") // Use wildcards for searching
	}
	if author != "" {
		query += " AND author ILIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+author+"%")
	}
	if category != "" {
		query += " AND (category ILIKE $" + strconv.Itoa(len(args)+1) + " OR subject ILIKE $" + strconv.Itoa(len(args)+2) + " OR genre ILIKE $" + strconv.Itoa(len(args)+3) + ")"
		args = append(args, "%"+category+"%", "%"+category+"%", "%"+category+"%")
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan data into a new instance of book for each row
	for rows.Next() {
		book := new(domain.BookRes)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// CategoryBooks implements ports.BookRepository.
func (b *BookRepository) CategoryBooks(ctx context.Context, categoryType string, categoryValue string) ([]*domain.BookRes, error) {
	var books []*domain.BookRes

	// Dynamically create the query based on the category type
	var query string
	if categoryType == "subject" {
		query = "SELECT id, title, author, category, subject, genre, published_year, available, borrower_id, created_at FROM books WHERE subject = $1"
	} else if categoryType == "genre" {
		query = "SELECT id, title, author, category, subject, genre, published_year, available, borrower_id, created_at FROM books WHERE genre = $1"
	}

	rows, err := b.db.QueryContext(ctx, query, categoryValue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan data into a new instance of book for each row
	for rows.Next() {
		book := new(domain.BookRes)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// AvailableBooks implements ports.BookRepository.
func (b *BookRepository) AvailableBooks(ctx context.Context) ([]*domain.BookRes, error) {
	var books []*domain.BookRes

	rows, err := b.db.Query("SELECT * FROM books WHERE available=true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := new(domain.BookRes)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
