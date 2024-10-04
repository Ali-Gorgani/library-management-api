package grpc

import (
	"context"
	"library-management-api/books-service/core/usecase"
	pb "library-management-api/books-service/pkg/proto"
)

type BookController struct {
	pb.UnimplementedBookServiceServer
	bookUsecase *usecase.BookUsecase
}

func NewBookController() *BookController {
	return &BookController{
		bookUsecase: usecase.NewBookUseCase(),
	}
}

func (bc *BookController) AddBook(ctx context.Context, req *pb.AddBookReq) (*pb.BookRes, error) {
	book, err := bc.bookUsecase.AddBook(ctx, toBookDomainAddBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoAddBook(book), nil
}

func (s *BookController) GetBooks(ctx context.Context, req *pb.GetBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoGetBooks(books)}, nil
}

func (s *BookController) GetBook(ctx context.Context, req *pb.GetBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.GetBook(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return toBookProtoGetBook(book), nil
}

func (s *BookController) UpdateBook(ctx context.Context, req *pb.UpdateBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.UpdateBook(ctx, int(req.Id), toBookRepositoryUpdateBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoUpdateBook(book), nil
}

func (s *BookController) DeleteBook(ctx context.Context, req *pb.DeleteBookReq) (*pb.DeleteBookRes, error) {
	err := s.bookRepository.DeleteBook(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBookRes{}, nil
}

func (s *BookController) BorrowBook(ctx context.Context, req *pb.BorrowBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.BorrowBook(ctx, toBookRepositoryBorrowBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoBorrowBook(book), nil
}

func (s *BookController) ReturnBook(ctx context.Context, req *pb.ReturnBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.ReturnBook(ctx, toBookRepositoryReturnBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoBorrowBook(book), nil
}

func (s *BookController) SearchBooks(ctx context.Context, req *pb.SearchBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.SearchBooks(ctx, req.Title, req.Author, req.Category)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoSearchBooks(books)}, nil
}

func (s *BookController) CategoryBooks(ctx context.Context, req *pb.CategoryBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.CategoryBooks(ctx, req.CategoryType, req.CategoryValue)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoCategoryBooks(books)}, nil
}

func (s *BookController) AvailableBooks(ctx context.Context, req *pb.AvailableBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.AvailableBooks(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoAvailableBooks(books)}, nil
}
