package grpc

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "library-management-api/books-service/core/usecase"
    pb "library-management-api/books-service/pkg/proto"
)

type BookController struct {
    pb.BookServiceServer
    bookUsecase *usecase.BookUsecase
}

func NewBookController() *BookController {
    return &BookController{
       bookUsecase: usecase.NewBookUseCase(),
    }
}

func (bc *BookController) GetBooks(*pb.GetBooksReq, grpc.ServerStreamingServer[pb.BookRes]) error {
    return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}

func (bc *BookController) AddBook(ctx context.Context, req *pb.AddBookReq) (*pb.BookRes, error) {
    book, err := bc.bookUsecase.AddBook(ctx, toBookDomainAddBook(req))
    if err != nil {
       return nil, err
    }
    return toBookProtoAddBook(book), nil
}

//func (s *BookController) GetBooks(ctx context.Context, req *pb.GetBooksReq) (*pb.GetBooksRes, error) {
//  books, err := s.bookUsecase.GetBooks(ctx)
//  if err != nil {
//     return nil, err
//  }
//  return &pb.GetBooksRes{Books: toBooksProtoGetBooks(books)}, nil
//}

func (s *BookController) GetBook(ctx context.Context, req *pb.GetBookReq) (*pb.BookRes, error) {
    book, err := s.bookUsecase.GetBook(ctx, int(req.Id))
    if err != nil {
       return nil, err
    }
    return toBookProtoGetBook(book), nil
}

func (s *BookController) UpdateBook(ctx context.Context, req *pb.UpdateBookReq) (*pb.BookRes, error) {
    book, err := s.bookUsecase.UpdateBook(ctx, int(req.Id), toBookRepositoryUpdateBook(req))
    if err != nil {
       return nil, err
    }
    return toBookProtoUpdateBook(book), nil
}

func (s *BookController) DeleteBook(ctx context.Context, req *pb.DeleteBookReq) (*pb.DeleteBookRes, error) {
    err := s.bookUsecase.DeleteBook(ctx, int(req.Id))
    if err != nil {
       return nil, err
    }
    return &pb.DeleteBookRes{}, nil
}

func (s *BookController) BorrowBook(ctx context.Context, req *pb.BorrowBookReq) (*pb.BookRes, error) {
    book, err := s.bookUsecase.BorrowBook(ctx, toBookRepositoryBorrowBook(req))
    if err != nil {
       return nil, err
    }
    return toBookProtoBorrowBook(book), nil
}

func (s *BookController) ReturnBook(ctx context.Context, req *pb.ReturnBookReq) (*pb.BookRes, error) {
    book, err := s.bookUsecase.ReturnBook(ctx, toBookRepositoryReturnBook(req))
    if err != nil {
       return nil, err
    }
    return toBookProtoBorrowBook(book), nil
}

func (s *BookController) SearchBooks(*pb.SearchBooksReq, grpc.ServerStreamingServer[pb.BookRes]) error {
    return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}

func (s *BookController) CategoryBooks(*pb.CategoryBooksReq, grpc.ServerStreamingServer[pb.BookRes]) error {
    return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")

}

func (s *BookController) AvailableBooks(*pb.AvailableBooksReq, grpc.ServerStreamingServer[pb.BookRes]) error {
    return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")

}