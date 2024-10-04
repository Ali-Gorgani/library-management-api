package usecase

import (
	"context"
	"errors"
	"library-management-api/books-service/adapter/repository"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
)

type BookUsecase struct {
	bookRepository ports.BookRepository
}

func NewBookUseCase() *BookUsecase {
	return &BookUsecase{
		bookRepository: repository.NewBookRepository(),
	}
}

func (b *BookUsecase) AddBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	addedBook, err := b.bookRepository.AddBook(ctx, book)
	if err != nil {
		return &domain.Book{}, err
	}
	return addedBook, nil
}

func (b *BookUsecase) GetBooks(ctx context.Context) ([]*domain.Book, error) {
	books, err := b.bookRepository.GetBooks(ctx)
	if err != nil {
		return []*domain.Book{}, err
	}
	return books, nil
}

func (b *BookUsecase) GetBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return &domain.Book{}, err
	}
	return foundBook, nil
}

func (b *BookUsecase) UpdateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	updatedBook, err := b.bookRepository.UpdateBook(ctx, book)
	if err != nil {
		return &domain.Book{}, err
	}
	return updatedBook, nil
}

func (b *BookUsecase) DeleteBook(ctx context.Context, book *domain.Book) error {
	err := b.bookRepository.DeleteBook(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookUsecase) BorrowBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return &domain.Book{}, err
	}
	if foundBook.Available == false {
		return &domain.Book{}, errors.New("book is already borrowed")
	}
	foundBook.Available = false
	foundBook.BorrowerID = book.BorrowerID
	borrowedBook, err := b.bookRepository.UpdateBook(ctx, foundBook)
	if err != nil {
		return &domain.Book{}, err
	}
	return borrowedBook, nil
}

func (b *BookUsecase) ReturnBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	foundBook, err := b.bookRepository.GetBook(ctx, book)
	if err != nil {
		return &domain.Book{}, err
	}
	if foundBook.Available == true {
		return &domain.Book{}, errors.New("book is already returned")
	}
	if *foundBook.BorrowerID != *book.BorrowerID {
		return &domain.Book{}, errors.New("borrower ID does not match")
	}
	foundBook.Available = true
	foundBook.BorrowerID = nil
	returnedBook, err := b.bookRepository.UpdateBook(ctx, foundBook)
	if err != nil {
		return &domain.Book{}, err
	}
	return returnedBook, nil
}

func (b *BookUsecase) SearchBooks(ctx context.Context, book *domain.Book) ([]*domain.Book, error) {
	books, err := b.bookRepository.SearchBooks(ctx, book)
	if err != nil {
		return []*domain.Book{}, err
	}
	return books, nil
}

func (b *BookUsecase) CategoryBooks(ctx context.Context, book *domain.Book) ([]*domain.Book, error) {
	books, err := b.bookRepository.CategoryBooks(ctx, book)
	if err != nil {
		return []*domain.Book{}, err
	}
	return books, nil
}

func (b *BookUsecase) AvailableBooks(ctx context.Context) ([]*domain.Book, error) {
	books, err := b.bookRepository.AvailableBooks(ctx)
	if err != nil {
		return []*domain.Book{}, err
	}
	return books, nil
}
