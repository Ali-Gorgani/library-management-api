package server

import (
	"library-management-api/books-service/api/pb"
	"library-management-api/books-service/core/domain"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func toBookRepositoryAddBook(req *pb.AddBookReq) *domain.AddBookReq {
	return &domain.AddBookReq{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
	}
}

func toBookProtoAddBook(book *domain.BookRes) *pb.BookRes {
	var borrowerID int32
	if book.BorrowerID != nil {
		borrowerID = *book.BorrowerID
	}

	return &pb.BookRes{
		Id:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		Category:      book.Category,
		Subject:       book.Subject,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
		Available:     book.Available,
		BorrowerId:    &borrowerID, // Direct assignment
		CreatedAt:     timestamppb.New(book.CreatedAt),
	}
}

func toBooksProtoGetBooks(books []*domain.BookRes) []*pb.BookRes {
	var pbBooks []*pb.BookRes
	for _, book := range books {
		pbBooks = append(pbBooks, toBookProtoAddBook(book))
	}
	return pbBooks
}

func toBookRepositoryUpdateBook(req *pb.UpdateBookReq) *domain.UpdateBookReq {
	return &domain.UpdateBookReq{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
		Available:     req.Available,
		BorrowerID:    req.BorrowerId, // Direct pointer assignment
	}
}

func toBookProtoUpdateBook(book *domain.BookRes) *pb.BookRes {
	return toBookProtoAddBook(book)
}

func toBookRepositoryBorrowBook(req *pb.BorrowBookReq) *domain.BorrowBookReq {
	return &domain.BorrowBookReq{
		BookID:     req.BookId,
		BorrowerID: req.BorrowerId,
	}
}

func toBookProtoBorrowBook(book *domain.BookRes) *pb.BookRes {
	return toBookProtoAddBook(book)
}

func toBookRepositoryReturnBook(req *pb.ReturnBookReq) *domain.ReturnBookReq {
	return &domain.ReturnBookReq{
		BookID:     req.BookId,
		BorrowerID: req.BorrowerId,
	}
}

func toBookProtoReturnBook(book *domain.BookRes) *pb.BookRes {
	return toBookProtoAddBook(book)
}

func toBooksProtoSearchBooks(books []*domain.BookRes) []*pb.BookRes {
	var pbBooks []*pb.BookRes
	for _, book := range books {
		pbBooks = append(pbBooks, toBookProtoAddBook(book))
	}
	return pbBooks
}

func toBooksProtoCategoryBooks(books []*domain.BookRes) []*pb.BookRes {
	var pbBooks []*pb.BookRes
	for _, book := range books {
		pbBooks = append(pbBooks, toBookProtoAddBook(book))
	}
	return pbBooks
}

func toBooksProtoAvailableBooks(books []*domain.BookRes) []*pb.BookRes {
	var pbBooks []*pb.BookRes
	for _, book := range books {
		pbBooks = append(pbBooks, toBookProtoAddBook(book))
	}
	return pbBooks
}
