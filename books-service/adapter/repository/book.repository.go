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

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository() ports.BookRepository {
	return &BookRepository{
		db: database.P().DB,
	}
}

// Create implements ports.BookRepository.
func (b *BookRepository) AddBook(ctx context.Context, book domain.AddBookParam) (domain.Book, error) {
	var addedBook domain.Book
	publishedYear, err := strconv.Atoi(book.PublishedYear)
	if err != nil {
		return domain.Book{}, err
	}
	query := "INSERT INTO books (title, author, category, subject, genre, published_year) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Category, book.Subject, book.Genre, publishedYear)
	err = row.Scan(&addedBook.ID, &addedBook.Title, &addedBook.Author, &addedBook.Category, &addedBook.Subject, &addedBook.Genre, &addedBook.PublishedYear, &addedBook.Available, &addedBook.BorrowerID, &addedBook.CreatedAt)
	if err != nil {
		return domain.Book{}, err
	}
	return addedBook, nil
}

// GetBooks implements ports.BookRepository.
func (b *BookRepository) GetBooks(ctx context.Context) ([]domain.Book, error) {
	var books []domain.Book
	var book domain.Book

	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// UpdateBook implements ports.BookRepository.
func (b *BookRepository) UpdateBook(ctx context.Context, id int, book domain.UpdateBookParam) (domain.Book, error) {
	var updatedBook domain.Book
	publishedYear, err := strconv.Atoi(book.PublishedYear)
	if err != nil {
		return domain.Book{}, err
	}
	query := "UPDATE books SET title=$1, author=$2, category=$3, subject=$4, genre=$5, published_year=$6, available=$7, borrower_id=$8 WHERE id=$9 RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Category, book.Subject, book.Genre, publishedYear, book.Available, book.BorrowerID, id)
	err = row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Category, &updatedBook.Subject, &updatedBook.Genre, &updatedBook.PublishedYear, &updatedBook.Available, &updatedBook.BorrowerID, &updatedBook.CreatedAt)
	if err != nil {
		return domain.Book{}, err
	}
	return updatedBook, nil
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
func (b *BookRepository) BorrowBook(ctx context.Context, borrowBook domain.BorrowBookRequest) (domain.Book, error) {
	var borrowedBook domain.Book
	// First, retrieve the book to check its availability
	query := "SELECT available FROM books WHERE id=$1"
	row := b.db.QueryRow(query, borrowBook.BookID)
	err := row.Scan(&borrowedBook.Available)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Book{}, errors.New("book not found")
		}
		return domain.Book{}, err
	}

	// Check if the book is already borrowed
	if !borrowedBook.Available {
		return domain.Book{}, errors.New("book is already borrowed")
	}

	// Update the book to mark it as borrowed and set the borrower ID
	query = "UPDATE books SET available=false, borrower_id=$1 WHERE id=$2 RETURNING *"
	row = b.db.QueryRow(query, borrowBook.BorrowerID, borrowBook.BookID)
	err = row.Scan(&borrowedBook.ID, &borrowedBook.Title, &borrowedBook.Author, &borrowedBook.Category, &borrowedBook.Subject, &borrowedBook.Genre, &borrowedBook.PublishedYear, &borrowedBook.Available, &borrowedBook.BorrowerID, &borrowedBook.CreatedAt)
	if err != nil {
		return domain.Book{}, err
	}
	return borrowedBook, nil
}

// ReturnBook implements ports.BookRepository.
func (b *BookRepository) ReturnBook(ctx context.Context, borrowBook domain.BorrowBookRequest) (domain.Book, error) {
	var returnedBook domain.Book

	// First, retrieve the book to check its availability and borrower ID
	query := "SELECT available, borrower_id FROM books WHERE id=$1"
	row := b.db.QueryRow(query, borrowBook.BookID)
	err := row.Scan(&returnedBook.Available, &returnedBook.BorrowerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Book{}, errors.New("book not found")
		}
		return domain.Book{}, err
	}

	// Check if the book is already available
	if returnedBook.Available {
		return domain.Book{}, errors.New("book is already available")
	}

	// Check if the borrower ID matches
	if *returnedBook.BorrowerID != *borrowBook.BorrowerID {
		return domain.Book{}, errors.New("borrower ID does not match")
	}

	// Update the book to mark it as available and clear the borrower ID
	query = "UPDATE books SET available=true, borrower_id=NULL WHERE id=$1 RETURNING *"
	row = b.db.QueryRow(query, borrowBook.BookID)
	err = row.Scan(&returnedBook.ID, &returnedBook.Title, &returnedBook.Author, &returnedBook.Category, &returnedBook.Subject, &returnedBook.Genre, &returnedBook.PublishedYear, &returnedBook.Available, &returnedBook.BorrowerID, &returnedBook.CreatedAt)
	if err != nil {
		return domain.Book{}, err
	}
	return returnedBook, nil
}

// SearchBooks implements ports.BookRepository.
func (b *BookRepository) SearchBooks(ctx context.Context, query string) ([]domain.Book, error) {
	var books []domain.Book

	// Use a prepared statement with safe parameter substitution
	searchQuery := "SELECT id, title, author, category, subject, genre, published_year, available, borrower_id, created_at FROM books WHERE title ILIKE $1 OR author ILIKE $1 OR category ILIKE $1"
	rows, err := b.db.QueryContext(ctx, searchQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan each result into a new book instance
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	// Check if there were any errors during row iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// CategoryBooks implements ports.BookRepository.
func (b *BookRepository) CategoryBooks(ctx context.Context, category string) ([]domain.Book, error) {
	var books []domain.Book

	// Use a prepared statement with explicit column selection
	query := "SELECT id, title, author, category, subject, genre, published_year, available, borrower_id, created_at FROM books WHERE subject=$1 OR genre=$1"
	rows, err := b.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan data into a new instance of book for each row
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	// Check if there were any errors during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

// AvailableBooks implements ports.BookRepository.
func (b *BookRepository) AvailableBooks(ctx context.Context) ([]domain.Book, error) {
	var books []domain.Book
	var book domain.Book

	rows, err := b.db.Query("SELECT * FROM books WHERE available=true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
