package repository

import (
	"context"
	"database/sql"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
	"library-management-api/books-service/init/database"
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
	query := "INSERT INTO books (title, author, subject, category, subject, genre, published_year) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Subject, book.Genre, book.PublishedYear)
	err := row.Scan(&addedBook.ID, &addedBook.Title, &addedBook.Author, &addedBook.Subject, &addedBook.Genre, &addedBook.PublishedYear, &addedBook.Available, &addedBook.CreatedAt)
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
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.CreatedAt)
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
	query := "UPDATE books SET title=$1, author=$2, category=$3, subject=$4, genre=$5, published_year=$6, available=$7 WHERE id=$8 RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Category, book.Subject, book.Genre, book.PublishedYear, book.Available, id)
	err := row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Category, &updatedBook.Subject, &updatedBook.Genre, &updatedBook.PublishedYear, &updatedBook.Available, &updatedBook.CreatedAt)
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
func (b *BookRepository) BorrowBook(ctx context.Context, id int) (domain.Book, error) {
	var borrowedBook domain.Book
	query := "UPDATE books SET available=false WHERE id=$1 RETURNING *"
	row := b.db.QueryRow(query, id)
	err := row.Scan(&borrowedBook.ID, &borrowedBook.Title, &borrowedBook.Author, &borrowedBook.Category, &borrowedBook.Subject, &borrowedBook.Genre, &borrowedBook.PublishedYear, &borrowedBook.Available, &borrowedBook.CreatedAt)
	if err != nil {
		return domain.Book{}, err
	}
	return borrowedBook, nil
}

// ReturnBook implements ports.BookRepository.
func (b *BookRepository) ReturnBook(ctx context.Context, id int) (domain.Book, error) {
	var returnedBook domain.Book
	query := "UPDATE books SET available=true WHERE id=$1 RETURNING *"
	row := b.db.QueryRow(query, id)
	err := row.Scan(&returnedBook.ID, &returnedBook.Title, &returnedBook.Author, &returnedBook.Category, &returnedBook.Subject, &returnedBook.Genre, &returnedBook.PublishedYear, &returnedBook.Available, &returnedBook.CreatedAt)
	if err != nil {
		return domain.Book{}, err
	}
	return returnedBook, nil
}

// SearchBooks implements ports.BookRepository.
func (b *BookRepository) SearchBooks(ctx context.Context, query string) ([]domain.Book, error) {
	var books []domain.Book

	// Use a prepared statement with safe parameter substitution
	searchQuery := "SELECT id, title, author, category, subject, genre, published_year, available, created_at FROM books WHERE title ILIKE $1 OR author ILIKE $1 OR category ILIKE $1"
	rows, err := b.db.QueryContext(ctx, searchQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan each result into a new book instance
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.CreatedAt)
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
	query := "SELECT id, title, author, category, subject, genre, published_year, available, created_at FROM books WHERE subject=$1 OR genre=$1"
	rows, err := b.db.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan data into a new instance of book for each row
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.CreatedAt)
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
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
