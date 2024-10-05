package repository

import (
	"database/sql"
	"library-management-api/books-service/core/domain"
)

type Book struct {
	ID            uint
	Title         sql.NullString
	Author        sql.NullString
	Category      sql.NullString
	Subject       sql.NullString
	Genre         sql.NullString
	PublishedYear uint
	Available     sql.NullBool
	BorrowerID    uint
	CreatedAt     sql.NullTime
}

func MapBookEntityToBookDomain(book Book) domain.Book {
	return domain.Book{
		ID:            book.ID,
		Title:         book.Title.String,
		Author:        book.Author.String,
		Category:      book.Category.String,
		Subject:       book.Subject.String,
		Genre:         book.Genre.String,
		PublishedYear: book.PublishedYear,
		Available:     book.Available.Bool,
		BorrowerID:    book.BorrowerID,
		CreatedAt:     book.CreatedAt.Time,
	}
}

func MapBooksEntityToBooksDomain(books []Book) []domain.Book {
	var res []domain.Book
	for _, book := range books {
		res = append(res, MapBookEntityToBookDomain(book))
	}
	return res
}

func MapBookDomainToBookEntity(book domain.Book) Book {
	return Book{
		ID:            book.ID,
		Title:         sql.NullString{String: book.Title, Valid: book.Title != ""},
		Author:        sql.NullString{String: book.Author, Valid: book.Author != ""},
		Category:      sql.NullString{String: book.Category, Valid: book.Category != ""},
		Subject:       sql.NullString{String: book.Subject, Valid: book.Subject != ""},
		Genre:         sql.NullString{String: book.Genre, Valid: book.Genre != ""},
		PublishedYear: book.PublishedYear,
		Available:     sql.NullBool{Bool: book.Available, Valid: true},
		BorrowerID:    book.BorrowerID,
		CreatedAt:     sql.NullTime{Time: book.CreatedAt, Valid: true},
	}
}
