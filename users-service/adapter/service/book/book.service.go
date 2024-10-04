package service

import (
	"context"
	"library-management-api/books-service/api/http"
	"library-management-api/books-service/core/domain"
	"library-management-api/users-service/third-party/book"

	"github.com/rs/zerolog/log"
)

type BookService struct {
	c book.IClient
}

func NewBookService() *BookService {
	c, err := book.NewClient()
	if err != nil {
		log.Error().Err(err).Msg("failed to create book grpc client")
		return nil
	}
	return &BookService{
		c: c,
	}
}

func (bs *BookService) AddBook(ctx context.Context, req domain.Book) (domain.Book, error) {
	addedBook, err := bs.c.AddBook(ctx, MapDomainBookToAddBookReq(req))
	if err != nil {
		return domain.Book{}, err
	}
	res := MapBookResToDomainBook(addedBook)
	return res, nil
}

func (bs *BookService) GetBooks(ctx context.Context) ([]domain.Book, error) {
	books, err := bs.c.GetBooks(ctx, http.GetBooksReq{})
	if err != nil {
		return []domain.Book{}, err
	}
	res := MapBooksResToDomainBooks(books)
	return res, nil
}

func (bs *BookService) GetBook(ctx context.Context, req domain.Book) (domain.Book, error) {
	book, err := bs.c.GetBook(ctx, MapDomainBookToGetBookReq(req))
	if err != nil {
		return domain.Book{}, err
	}
	res := MapBookResToDomainBook(book)
	return res, nil
}

func (bs *BookService) UpdateBook(ctx context.Context, req domain.Book) (domain.Book, error) {
	updatedBook, err := bs.c.UpdateBook(ctx, MapDomainBookToUpdateBookReq(req))
	if err != nil {
		return domain.Book{}, err
	}
	res := MapBookResToDomainBook(updatedBook)
	return res, nil
}

func (bs *BookService) DeleteBook(ctx context.Context, req domain.Book) error {
	err := bs.c.DeleteBook(ctx, MapDomainBookToDeleteBookReq(req))
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookService) BorrowBook(ctx context.Context, req domain.Book) (domain.Book, error) {
	borrowedBook, err := bs.c.BorrowBook(ctx, MapDomainBookToBorrowBookReq(req))
	if err != nil {
		return domain.Book{}, err
	}
	res := MapBookResToDomainBook(borrowedBook)
	return res, nil
}

func (bs *BookService) ReturnBook(ctx context.Context, req domain.Book) (domain.Book, error) {
	returnedBook, err := bs.c.ReturnBook(ctx, MapDomainBookToReturnBookReq(req))
	if err != nil {
		return domain.Book{}, err
	}
	res := MapBookResToDomainBook(returnedBook)
	return res, nil
}

func (bs *BookService) SearchBooks(ctx context.Context, req domain.Book) ([]domain.Book, error) {
	books, err := bs.c.SearchBooks(ctx, MapDomainBookToSearchBooksReq(req))
	if err != nil {
		return []domain.Book{}, err
	}
	res := MapBooksResToDomainBooks(books)
	return res, nil
}

func (bs *BookService) CategoryBooks(ctx context.Context, req domain.Book) ([]domain.Book, error) {
	books, err := bs.c.CategoryBooks(ctx, MapDomainBookToCategoryBooksReq(req))
	if err != nil {
		return []domain.Book{}, err
	}
	res := MapBooksResToDomainBooks(books)
	return res, nil
}

func (bs *BookService) AvailableBooks(ctx context.Context) ([]domain.Book, error) {
	books, err := bs.c.AvailableBooks(ctx, http.AvailableBooksReq{})
	if err != nil {
		return []domain.Book{}, err
	}
	res := MapBooksResToDomainBooks(books)
	return res, nil
}
