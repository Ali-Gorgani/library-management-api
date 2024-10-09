package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

// AddBook implements ports.BookRepository.
func (b *BookRepository) AddBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	var addedBook Book
	mappedBook := MapBookDomainToBookEntity(book)

	query := "INSERT INTO books (title, author, category, subject, genre, published_year, available, borrower_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"
	row := b.db.QueryRow(query, mappedBook.Title, mappedBook.Author, mappedBook.Category, mappedBook.Subject, mappedBook.Genre, mappedBook.PublishedYear, mappedBook.Available, mappedBook.BorrowerID)
	err := row.Scan(&addedBook.ID, &addedBook.Title, &addedBook.Author, &addedBook.Category, &addedBook.Subject, &addedBook.Genre, &addedBook.PublishedYear, &addedBook.Available, &addedBook.BorrowerID, &addedBook.CreatedAt)
	if err != nil {
		return domain.Book{}, err
	}
	res := MapBookEntityToBookDomain(addedBook)
	return res, nil
}

// GetBooks implements ports.BookRepository.
func (b *BookRepository) GetBooks(ctx context.Context) ([]domain.Book, error) {
	var books []Book

	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		return []domain.Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return []domain.Book{}, err
		}
		books = append(books, book)
	}
	res := MapBooksEntityToBooksDomain(books)
	return res, nil
}

// GetBook implements ports.BookRepository.
func (b *BookRepository) GetBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	var foundBook Book

	query := "SELECT * FROM books WHERE id=$1"
	row := b.db.QueryRow(query, book.ID)
	err := row.Scan(&foundBook.ID, &foundBook.Title, &foundBook.Author, &foundBook.Category, &foundBook.Subject, &foundBook.Genre, &foundBook.PublishedYear, &foundBook.Available, &foundBook.BorrowerID, &foundBook.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Book{}, errorhandler.ErrBookNotFound
		}
		return domain.Book{}, err
	}
	res := MapBookEntityToBookDomain(foundBook)
	return res, nil
}

// UpdateBook implements ports.BookRepository.
func (b *BookRepository) UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	var updatedBook Book

	mappedBook := MapBookDomainToBookEntity(book)
	query := "UPDATE books SET title=$1, author=$2, category=$3, subject=$4, genre=$5, published_year=$6, available=$7, borrower_id=$8 WHERE id=$9 RETURNING *"
	row := b.db.QueryRow(query, mappedBook.Title, mappedBook.Author, mappedBook.Category, mappedBook.Subject, mappedBook.Genre, mappedBook.PublishedYear, mappedBook.Available, mappedBook.BorrowerID, mappedBook.ID)
	err := row.Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Author, &updatedBook.Category, &updatedBook.Subject, &updatedBook.Genre, &updatedBook.PublishedYear, &updatedBook.Available, &updatedBook.BorrowerID, &updatedBook.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Book{}, errorhandler.ErrBookNotFound
		}
		return domain.Book{}, err
	}
	res := MapBookEntityToBookDomain(updatedBook)
	return res, nil
}

// DeleteBook implements ports.BookRepository.
func (b *BookRepository) DeleteBook(ctx context.Context, book domain.Book) error {
	query := "DELETE FROM books WHERE id=$1"
	_, err := b.db.Exec(query, book.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorhandler.ErrBookNotFound
		}
		return err
	}
	return nil
}

// SearchBooks implements ports.BookRepository.
func (b *BookRepository) SearchBooks(ctx context.Context, book domain.Book) ([]domain.Book, error) {
	var books []Book

	mappedBook := MapBookDomainToBookEntity(book)

	// Build the SQL query dynamically based on which parameters are provided
	query := "SELECT * FROM books WHERE TRUE"
	var args []interface{}
	argCounter := 1

	// Add conditions based on provided search parameters
	if mappedBook.Title.Valid {
		query += " AND title ILIKE $" + strconv.Itoa(argCounter)
		args = append(args, mappedBook.Title.String)
		argCounter++
	}
	if mappedBook.Author.Valid {
		query += " AND author ILIKE $" + strconv.Itoa(argCounter)
		args = append(args, mappedBook.Author.String)
		argCounter++
	}
	if mappedBook.Category.Valid {
		query += " AND category ILIKE $" + strconv.Itoa(argCounter)
		args = append(args, mappedBook.Category.String)
		argCounter++
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return []domain.Book{}, err
	}
	defer rows.Close()

	// Iterate over the rows and scan data into a new instance of book for each row
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return []domain.Book{}, err
		}
		books = append(books, book)
	}
	res := MapBooksEntityToBooksDomain(books)
	return res, nil
}

// CategoryBooks implements ports.BookRepository.
func (b *BookRepository) CategoryBooks(ctx context.Context, book domain.Book) ([]domain.Book, error) {
	var books []Book
	var categoryType, categoryValue string

	mappedBook := MapBookDomainToBookEntity(book)

	// Dynamically create the query based on the category type
	if mappedBook.Subject.Valid {
		categoryType = "subject"
		categoryValue = mappedBook.Subject.String
	} else if mappedBook.Genre.Valid {
		categoryType = "genre"
		categoryValue = mappedBook.Genre.String
	}

	query := fmt.Sprintf("SELECT * FROM books WHERE %s=$1", categoryType)
	rows, err := b.db.Query(query, categoryValue)
	if err != nil {
		return []domain.Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return []domain.Book{}, err
		}
		books = append(books, book)
	}
	res := MapBooksEntityToBooksDomain(books)
	return res, nil
}

// AvailableBooks implements ports.BookRepository.
func (b *BookRepository) AvailableBooks(ctx context.Context) ([]domain.Book, error) {
	var books []Book

	query := "SELECT * FROM books WHERE available=true"
	rows, err := b.db.Query(query)
	if err != nil {
		return []domain.Book{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category, &book.Subject, &book.Genre, &book.PublishedYear, &book.Available, &book.BorrowerID, &book.CreatedAt)
		if err != nil {
			return []domain.Book{}, err
		}
		books = append(books, book)
	}
	res := MapBooksEntityToBooksDomain(books)
	return res, nil
}
