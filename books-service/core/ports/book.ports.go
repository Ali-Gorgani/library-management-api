package ports

import (
	"context"
	"library-management-api/books-service/core/domain"
)

type BookRepository interface {
	AddBook(ctx context.Context, book domain.AddBookParam) (domain.Book, error)
	GetBooks(ctx context.Context) ([]domain.Book, error)
	UpdateBook(ctx context.Context, id int, book domain.UpdateBookParam) (domain.Book, error)
	DeleteBook(ctx context.Context, id int) error
	BorrowBook(ctx context.Context, id int) (domain.Book, error)
	ReturnBook(ctx context.Context, id int) (domain.Book, error)
	SearchBooks(ctx context.Context, query string) ([]domain.Book, error)
	CategoryBooks(ctx context.Context, category string) ([]domain.Book, error)
	AvailableBooks(ctx context.Context) ([]domain.Book, error)
}
