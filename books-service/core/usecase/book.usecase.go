package usecase

import (
	"context"
	"errors"
	"library-management-api/auth-service/pkg/token"
	"library-management-api/books-service/adapter/repository"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
	"library-management-api/util/errorhandler"
)

type BookUsecase struct {
	bookRepository ports.BookRepository
}

func NewBookUseCase() *BookUsecase {
	return &BookUsecase{
		bookRepository: repository.NewBookRepository(),
	}
}

func (b *BookUsecase) AddBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	_, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return domain.Book{}, errorhandler.ErrForbidden
	}

	addedBook, err := b.bookRepository.AddBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}
	return addedBook, nil
}

func (b *BookUsecase) GetBooks(ctx context.Context) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	_, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return []domain.Book{}, errorhandler.ErrForbidden
	}

	books, err := b.bookRepository.GetBooks(ctx)
	if err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}

func (b *BookUsecase) GetBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	_, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return domain.Book{}, errorhandler.ErrForbidden
	}

	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}
	return foundBook, nil
}

func (b *BookUsecase) UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	claims, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return domain.Book{}, errorhandler.ErrForbidden
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

func (b *BookUsecase) DeleteBook(ctx context.Context, book domain.Book) error {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	claims, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return errorhandler.ErrForbidden
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

func (b *BookUsecase) BorrowBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	claims, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return domain.Book{}, errorhandler.ErrForbidden
	}
	book.BorrowerID = claims.ID

	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}

	if !foundBook.Available {
		return domain.Book{}, errors.New("book is not available")
	}
	foundBook.Available = false
	foundBook.BorrowerID = book.BorrowerID

	borrowedBook, err := b.bookRepository.UpdateBook(ctx, foundBook)
	if err != nil {
		return domain.Book{}, err
	}
	return borrowedBook, nil
}

func (b *BookUsecase) ReturnBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	claims, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return domain.Book{}, errorhandler.ErrForbidden
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

func (b *BookUsecase) SearchBooks(ctx context.Context, book domain.Book) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	_, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return []domain.Book{}, errorhandler.ErrForbidden
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

func (b *BookUsecase) CategoryBooks(ctx context.Context, book domain.Book) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	_, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return []domain.Book{}, errorhandler.ErrForbidden
	}

	books, err := b.bookRepository.CategoryBooks(ctx, book)
	if err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}

func (b *BookUsecase) AvailableBooks(ctx context.Context) ([]domain.Book, error) {
	contextToken := ctx.Value("token").(string)
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="

	_, err := token.VerifyToken(contextToken, secretKey)
	if err != nil {
		return []domain.Book{}, errorhandler.ErrForbidden
	}

	books, err := b.bookRepository.AvailableBooks(ctx)
	if err != nil {
		return []domain.Book{}, err
	}
	return books, nil
}
