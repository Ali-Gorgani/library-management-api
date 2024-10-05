package ports

import (
	"context"
	"library-management-api/books-service/core/domain"
)

type BookRepository interface {
	AddBook(ctx context.Context, book domain.Book) (domain.Book, error)
	GetBooks(ctx context.Context) ([]domain.Book, error)
	GetBook(ctx context.Context, book domain.Book) (domain.Book, error)
	UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error)
	DeleteBook(ctx context.Context, book domain.Book) error
	SearchBooks(ctx context.Context, book domain.Book) ([]domain.Book, error)
	CategoryBooks(ctx context.Context, book domain.Book) ([]domain.Book, error)
	AvailableBooks(ctx context.Context) ([]domain.Book, error)
}
