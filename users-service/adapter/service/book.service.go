package service

import (
	"context"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/third-party/book"
	"log"
)

type BookService struct {
	c book.IClient
}

func NewBookService() *BookService {
	c, err := book.NewClient()
	if err != nil {
		log.Println()
		return nil
	}
	return &BookService{
		c: c,
	}
}

func (bs *BookService) AddBook(ctx context.Context, b domain.Books) (domain.Books, error) {
	var err error
	res, err = bs.c.AddBook(ctx, book.AddBookReq{
		// map b to addbookreq
	})
	return domain.Books{}, err // map to domain.books
}
