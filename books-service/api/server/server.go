package server

import (
	"context"
	"library-management-api/books-service/adapter/repository"
	"library-management-api/books-service/api/pb"
	"library-management-api/books-service/core/ports"
)

type Server struct {
	pb.UnimplementedBookServiceServer
	bookRepository ports.BookRepository
}

func NewServer() *Server {
	return &Server{
		bookRepository: repository.NewBookRepository(),
	}
}

func (s *Server) AddBook(ctx context.Context, req *pb.AddBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.AddBook(ctx, toBookRepositoryAddBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoAddBook(book), nil
}

func (s *Server) GetBooks(ctx context.Context, req *pb.GetBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoGetBooks(books)}, nil
}

func (s *Server) UpdateBook(ctx context.Context, req *pb.UpdateBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.UpdateBook(ctx, int(req.Id), toBookRepositoryUpdateBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoUpdateBook(book), nil
}

func (s *Server) DeleteBook(ctx context.Context, req *pb.DeleteBookReq) (*pb.DeleteBookRes, error) {
	err := s.bookRepository.DeleteBook(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBookRes{}, nil
}

func (s *Server) BorrowBook(ctx context.Context, req *pb.BorrowBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.BorrowBook(ctx, toBookRepositoryBorrowBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoBorrowBook(book), nil
}

func (s *Server) ReturnBook(ctx context.Context, req *pb.ReturnBookReq) (*pb.BookRes, error) {
	book, err := s.bookRepository.ReturnBook(ctx, toBookRepositoryReturnBook(req))
	if err != nil {
		return nil, err
	}
	return toBookProtoBorrowBook(book), nil
}

func (s *Server) SearchBooks(ctx context.Context, req *pb.SearchBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.SearchBooks(ctx, req.Title, req.Author, req.Category)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoSearchBooks(books)}, nil
}

func (s *Server) CategoryBooks(ctx context.Context, req *pb.CategoryBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.CategoryBooks(ctx, req.CategoryType, req.CategoryValue)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoCategoryBooks(books)}, nil
}

func (s *Server) AvailableBooks(ctx context.Context, req *pb.AvailableBooksReq) (*pb.GetBooksRes, error) {
	books, err := s.bookRepository.AvailableBooks(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetBooksRes{Books: toBooksProtoAvailableBooks(books)}, nil
}
