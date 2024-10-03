package repository

import (
	"context"
	"database/sql"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
	"library-management-api/books-service/init/database"
	"library-management-api/util/errorhandler"
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
func (b *BookRepository) AddBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	var addedBook domain.Book
	query := "INSERT INTO books (title, author, category, subject, genre, published_year) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Category, book.Subject, book.Genre, book.PublishedYear)
	err := row.Scan(&addedBook.ID, &addedBook.Title, &addedBook.Author, &addedBook.Category, &addedBook.Subject, &addedBook.Genre, &addedBook.PublishedYear, &addedBook.Available, &addedBook.BorrowerID, &addedBook.CreatedAt)
	if err != nil {
		return &domain.Book{}, err
	}
	return &addedBook, nil
}

// GetBooks implements ports.BookRepository.
func (b *BookRepository) GetBooks(ctx context.Context) ([]*domain.Book, error) {
	var books []*domain.Book

	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := new(domain.Book)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// GetBook implements ports.BookRepository.
func (b *BookRepository) GetBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	var foundBook domain.Book

	query := "SELECT * FROM books WHERE id=$1"
	row := b.db.QueryRow(query, book.ID)
	err := row.Scan(&foundBook.ID, &foundBook.Title, &foundBook.Author, &foundBook.Category, &foundBook.Subject, &foundBook.Genre, &foundBook.PublishedYear, &foundBook.Available, &foundBook.BorrowerID, &foundBook.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &domain.Book{}, errorhandler.ErrBookNotFound
		}
		return &domain.Book{}, err
	}
	return &foundBook, nil
}

// UpdateBook implements ports.BookRepository.
func (b *BookRepository) UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	var updatedBook domain.Book

	query := "UPDATE books SET title=$1, author=$2, category=$3, subject=$4, genre=$5, published_year=$6, available=$7, borrower_id=$8 WHERE id=$9 RETURNING *"
	row := b.db.QueryRow(query, book.Title, book.Author, book.Category, book.Subject, book.Genre, book.PublishedYear, book.Available, book.BorrowerID, book.ID)
	err := row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Category, &updatedBook.Subject, &updatedBook.Genre, &updatedBook.PublishedYear, &updatedBook.Available, &updatedBook.BorrowerID, &updatedBook.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &domain.Book{}, errorhandler.ErrBookNotFound
		}
		return &domain.Book{}, err
	}
	return &updatedBook, nil
}

// DeleteBook implements ports.BookRepository.
func (b *BookRepository) DeleteBook(ctx context.Context, book *domain.Book) error {
	query := "DELETE FROM books WHERE id=$1"
	_, err := b.db.Exec(query, book.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errorhandler.ErrBookNotFound
		}
		return err
	}
	return nil
}

// SearchBooks implements ports.BookRepository.
func (b *BookRepository) SearchBooks(ctx context.Context, book *domain.Book) ([]*domain.Book, error) {
	var books []*domain.Book

	// Build the SQL query dynamically based on which parameters are provided
	query := "SELECT id, title, author, category, subject, genre, published_year, available, borrower_id, created_at FROM books WHERE 1=1"
	var args []interface{}
	argCounter := 1

	// Add conditions based on provided search parameters
	if book.Title != "" {
		query += " AND title ILIKE $" + strconv.Itoa(argCounter)
		args = append(args, "%"+book.Title+"%") // Wildcard for partial match
		argCounter++
	}
	if book.Author != "" {
		query += " AND author ILIKE $" + strconv.Itoa(argCounter)
		args = append(args, "%"+book.Author+"%")
		argCounter++
	}
	if book.Category != "" {
		query += " AND category ILIKE $" + strconv.Itoa(argCounter)
		args = append(args, "%"+book.Category+"%")
		argCounter++
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and scan data into a new instance of book for each row
	for rows.Next() {
		book := new(domain.Book)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// CategoryBooks implements ports.BookRepository.
func (b *BookRepository) CategoryBooks(ctx context.Context, book *domain.Book) ([]*domain.Book, error) {
	var books []*domain.Book
	var categoryType, categoryValue string

	// Dynamically create the query based on the category type
	if book.Subject != "" {
		categoryType = "subject"
		categoryValue = book.Subject
	} else if book.Genre != "" {
		categoryType = "genre"
		categoryValue = book.Genre
	}

	query := "SELECT id, title, author, category, subject, genre, published_year, available, borrower_id, created_at FROM books WHERE $1 = $2"
	rows, err := b.db.QueryContext(ctx, query, categoryType, categoryValue)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Fetch and scan rows into the books slice
	for rows.Next() {
		book := new(domain.Book)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// AvailableBooks implements ports.BookRepository.
func (b *BookRepository) AvailableBooks(ctx context.Context) ([]*domain.Book, error) {
	var books []*domain.Book

	rows, err := b.db.Query("SELECT * FROM books WHERE available=true")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := new(domain.Book)
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
