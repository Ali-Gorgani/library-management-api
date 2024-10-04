package book

import (
	"context"
	"library-management-api/books-service/api/http"
	"library-management-api/users-service/pkg/book/pb"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type IClient interface {
	AddBook(ctx context.Context, req *http.AddBookReq) (*http.BookRes, error)
	GetBooks(ctx context.Context, req *http.GetBooksReq) ([]*http.BookRes, error)
	GetBook(ctx context.Context, req *http.GetBookReq) (*http.BookRes, error)
	UpdateBook(ctx context.Context, req *http.UpdateBookReq) (*http.BookRes, error)
	DeleteBook(ctx context.Context, req *http.DeleteBookReq) error
	BorrowBook(ctx context.Context, req *http.BorrowBookReq) (*http.BookRes, error)
	ReturnBook(ctx context.Context, req *http.ReturnBookReq) (*http.BookRes, error)
	SearchBooks(ctx context.Context, req *http.SearchBooksReq) ([]*http.BookRes, error)
	CategoryBooks(ctx context.Context, req *http.CategoryBooksReq) ([]*http.BookRes, error)
	AvailableBooks(ctx context.Context, req *http.AvailableBooksReq) ([]*http.BookRes, error)
}

type Client struct {
	c *grpc.ClientConn
}

func NewClient() (IClient, error) {
	client, err := grpc.NewClient("localhost:8081")
	if err != nil {
		log.Error().Err(err).Msg("failed to create grpc client")
		return nil, err
	}
	return &Client{
		c: client,
	}, nil
}

func (c *Client) AddBook(ctx context.Context, req *http.AddBookReq) (*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	addedBook, err := client.AddBook(context.Background(), MapAddBookReqToPb(req))
	if err != nil {
		return &http.BookRes{}, err
	}
	res := MapBookResToDto(addedBook)
	return res, nil
}

func (c *Client) GetBooks(ctx context.Context, req *http.GetBooksReq) ([]*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	books, err := client.GetBooks(context.Background(), &pb.GetBooksReq{})
	if err != nil {
		return []*http.BookRes{}, err
	}
	res := MapBooksResToDto(books)
	return res, nil
}

func (c *Client) GetBook(ctx context.Context, req *http.GetBookReq) (*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	book, err := client.GetBook(context.Background(), MapGetBookReqToPb(req))
	if err != nil {
		return &http.BookRes{}, err
	}
	res := MapBookResToDto(book)
	return res, nil
}

func (c *Client) UpdateBook(ctx context.Context, req *http.UpdateBookReq) (*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	updatedBook, err := client.UpdateBook(context.Background(), MapUpdateBookReqToPb(req))
	if err != nil {
		return &http.BookRes{}, err
	}
	res := MapBookResToDto(updatedBook)
	return res, nil
}

func (c *Client) DeleteBook(ctx context.Context, req *http.DeleteBookReq) error {
	client := pb.NewBookServiceClient(c.c)
	_, err := client.DeleteBook(context.Background(), MapDeleteBookReqToPb(req))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) BorrowBook(ctx context.Context, req *http.BorrowBookReq) (*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	borrowBook, err := client.BorrowBook(context.Background(), MapBorrowBookReqToPb(req))
	if err != nil {
		return &http.BookRes{}, err
	}
	res := MapBookResToDto(borrowBook)
	return res, nil
}

func (c *Client) ReturnBook(ctx context.Context, req *http.ReturnBookReq) (*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	returnBook, err := client.ReturnBook(context.Background(), MapReturnBookReqToPb(req))
	if err != nil {
		return &http.BookRes{}, err
	}
	res := MapBookResToDto(returnBook)
	return res, nil
}

func (c *Client) SearchBooks(ctx context.Context, req *http.SearchBooksReq) ([]*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	books, err := client.SearchBooks(context.Background(), MapSearchBooksReqToPb(req))
	if err != nil {
		return []*http.BookRes{}, err
	}
	res := MapBooksResToDto(books)
	return res, nil
}

func (c *Client) CategoryBooks(ctx context.Context, req *http.CategoryBooksReq) ([]*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	books, err := client.CategoryBooks(context.Background(), MapCategoryBooksReqToPb(req))
	if err != nil {
		return []*http.BookRes{}, err
	}
	res := MapBooksResToDto(books)
	return res, nil
}

func (c *Client) AvailableBooks(ctx context.Context, req *http.AvailableBooksReq) ([]*http.BookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	books, err := client.AvailableBooks(context.Background(), &pb.AvailableBooksReq{})
	if err != nil {
		return []*http.BookRes{}, err
	}
	res := MapBooksResToDto(books)
	return res, nil
}
