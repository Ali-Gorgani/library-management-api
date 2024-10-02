package usecase

import (
	"library-management-api/books-service/api/pb"
	"library-management-api/books-service/core/domain"
)

func toDomainBook(book *pb.BookRes) *domain.BookRes {
	var borrowerID *int32
	if book.BorrowerId == nil {
		borrowerID = nil
	}

	return &domain.BookRes{
		ID:            book.Id,
		Title:         book.Title,
		Author:        book.Author,
		Category:      book.Category,
		Subject:       book.Subject,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
		Available:     book.Available,
		BorrowerID:    borrowerID,
		CreatedAt:     book.CreatedAt.AsTime().Local(),
	}
}

func toDomainBooks(pbBooks []*pb.BookRes) []*domain.BookRes {
	var books []*domain.BookRes
	for _, pbBook := range pbBooks {
		books = append(books, toDomainBook(pbBook))
	}
	return books
}

func toProtoAddBookReq(req *domain.AddBookReq) *pb.AddBookReq {
	return &pb.AddBookReq{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
	}
}

func toProtoUpdateBookReq(req *domain.UpdateBookReq) *pb.UpdateBookReq {
	return &pb.UpdateBookReq{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
		Available:     req.Available,
		BorrowerId:    req.BorrowerID,
	}
}

func toProtoBorrowBookReq(req *domain.BorrowBookReq) *pb.BorrowBookReq {
	return &pb.BorrowBookReq{
		BookId:     req.BookID,
		BorrowerId: req.BorrowerID,
	}
}

func toProtoReturnBookReq(req *domain.ReturnBookReq) *pb.ReturnBookReq {
	return &pb.ReturnBookReq{
		BookId:     req.BookID,
		BorrowerId: req.BorrowerID,
	}
}
