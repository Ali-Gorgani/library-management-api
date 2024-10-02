package ports

import (
	"context"
	"library-management-api/books-service/core/domain"
)

type BookRepository interface {
	AddBook(ctx context.Context, book *domain.AddBookReq) (*domain.BookRes, error)
	GetBooks(ctx context.Context) ([]*domain.BookRes, error)
	UpdateBook(ctx context.Context, id int, book *domain.UpdateBookReq) (*domain.BookRes, error)
	DeleteBook(ctx context.Context, id int) error
	BorrowBook(ctx context.Context, borrowBook *domain.BorrowBookReq) (*domain.BookRes, error)
	ReturnBook(ctx context.Context, borrowBook *domain.ReturnBookReq) (*domain.BookRes, error)
	SearchBooks(ctx context.Context, query string) ([]*domain.BookRes, error)
	CategoryBooks(ctx context.Context, category string) ([]*domain.BookRes, error)
	AvailableBooks(ctx context.Context) ([]*domain.BookRes, error)
}
