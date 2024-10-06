package usecase

import (
	"context"
	"library-management-api/books-service/adapter/repository"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
	"library-management-api/util/errorhandler"
)

type BookUseCase struct {
	bookRepository ports.BookRepository
}

func NewBookUseCase() *BookUseCase {
	return &BookUseCase{
		bookRepository: repository.NewBookRepository(),
	}
}

func (b *BookUseCase) AddBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return domain.Book{}, errorhandler.ErrInvalidSession
	}

	addedBook, err := b.bookRepository.AddBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}
	return addedBook, nil
}

func (b *BookUseCase) GetBooks(ctx context.Context) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return []domain.Book{}, errorhandler.ErrInvalidSession
	}

	books, err := b.bookRepository.GetBooks(ctx)
	if err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}

func (b *BookUseCase) GetBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return domain.Book{}, errorhandler.ErrInvalidSession
	}

	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}
	return foundBook, nil
}

func (b *BookUseCase) UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return domain.Book{}, errorhandler.ErrInvalidSession
	}

	if !claims.IsAdmin {
		return domain.Book{}, errorhandler.ErrForbidden
	}

	updatedBook, err := b.bookRepository.UpdateBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}
	return updatedBook, nil
}

func (b *BookUseCase) DeleteBook(ctx context.Context, book domain.Book) error {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return errorhandler.ErrInvalidSession
	}

	if !claims.IsAdmin {
		return errorhandler.ErrForbidden
	}

	err = b.bookRepository.DeleteBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookUseCase) BorrowBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return domain.Book{}, errorhandler.ErrInvalidSession
	}
	book.BorrowerID = claims.ID

	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}

	if !foundBook.Available {
		return domain.Book{}, errorhandler.ErrBookAlreadyBorrowed
	}
	foundBook.Available = false
	foundBook.BorrowerID = book.BorrowerID

	borrowedBook, err := b.bookRepository.UpdateBook(ctx, foundBook)
	if err != nil {
		return domain.Book{}, err
	}
	return borrowedBook, nil
}

func (b *BookUseCase) ReturnBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return domain.Book{}, errorhandler.ErrInvalidSession
	}
	book.BorrowerID = claims.ID

	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}

	if foundBook.Available {
		return domain.Book{}, errorhandler.ErrBookAlreadyAvailable
	}
	foundBook.Available = true

	if foundBook.BorrowerID != book.BorrowerID {
		return domain.Book{}, errorhandler.ErrBorrowerIDMismatch
	}
	foundBook.BorrowerID = 0

	returnedBook, err := b.bookRepository.UpdateBook(ctx, foundBook)
	if err != nil {
		return domain.Book{}, err
	}
	return returnedBook, nil
}

func (b *BookUseCase) SearchBooks(ctx context.Context, book domain.Book) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return []domain.Book{}, errorhandler.ErrInvalidSession
	}

	// Check if at least one of the fields is provided
	if book.Title == "" && book.Author == "" && book.Category == "" {
		return []domain.Book{}, errorhandler.ErrInvalidSearchQuery
	}

	books, err := b.bookRepository.SearchBooks(ctx, book)
	if err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}

func (b *BookUseCase) CategoryBooks(ctx context.Context, book domain.Book) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return []domain.Book{}, errorhandler.ErrInvalidSession
	}

	books, err := b.bookRepository.CategoryBooks(ctx, book)
	if err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}

func (b *BookUseCase) AvailableBooks(ctx context.Context) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: verify contextToken with gRPC from auth-service and get claims
	if err != nil {
		return []domain.Book{}, errorhandler.ErrInvalidSession
	}

	books, err := b.bookRepository.AvailableBooks(ctx)
	if err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}
