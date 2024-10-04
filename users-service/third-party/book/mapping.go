package book

import (
	"library-management-api/books-service/api/http"
	"library-management-api/users-service/pkg/book/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapAddBookReqToPb(req *http.AddBookReq) *pb.AddBookReq {
	return &pb.AddBookReq{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: int32(req.PublishedYear),
	}
}

func MapGetBookReqToPb(req *http.GetBookReq) *pb.GetBookReq {
	return &pb.GetBookReq{
		Id: int32(req.ID),
	}
}

func MapUpdateBookReqToPb(req *http.UpdateBookReq) *pb.UpdateBookReq {
	return &pb.UpdateBookReq{
		Id:            int32(req.ID),
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: int32(req.PublishedYear),
		Available:     req.Available,
		BorrowerId:    mapIntToInt32(req.BorrowerID),
	}
}

func MapDeleteBookReqToPb(req *http.DeleteBookReq) *pb.DeleteBookReq {
	return &pb.DeleteBookReq{
		Id: int32(req.ID),
	}
}

func MapBorrowBookReqToPb(req *http.BorrowBookReq) *pb.BorrowBookReq {
	return &pb.BorrowBookReq{
		Id:         int32(req.ID),
		BorrowerId: mapIntToInt32(req.BorrowerID),
	}
}

func MapReturnBookReqToPb(req *http.ReturnBookReq) *pb.ReturnBookReq {
	return &pb.ReturnBookReq{
		Id:         int32(req.ID),
		BorrowerId: mapIntToInt32(req.BorrowerID),
	}
}

func MapSearchBooksReqToPb(req *http.SearchBooksReq) *pb.SearchBooksReq {
	return &pb.SearchBooksReq{
		Title:    req.Title,
		Author:   req.Author,
		Category: req.Category,
	}
}

func MapCategoryBooksReqToPb(req *http.CategoryBooksReq) *pb.CategoryBooksReq {
	return &pb.CategoryBooksReq{
		CategoryType:  req.CategoryType,
		CategoryValue: req.CategoryValue,
	}
}

func MapBookResToDto(res *pb.BookRes) *http.BookRes {
	return &http.BookRes{
		ID:            int(res.GetId()),
		Title:         res.GetTitle(),
		Author:        res.GetAuthor(),
		Category:      res.GetCategory(),
		Subject:       res.GetSubject(),
		Genre:         res.GetGenre(),
		PublishedYear: int(res.GetPublishedYear()),
		Available:     res.GetAvailable(),
		BorrowerID:    mapInt32ToInt(pointerToInt32(res.GetBorrowerId())),
		CreatedAt:     mapTimestampToTime(res.GetCreatedAt()),
	}
}

func MapBooksResToDto(req grpc.ServerStreamingClient[pb.BookRes]) []*http.BookRes {
	var res []*http.BookRes
	for {
		book, err := req.Recv()
		if err != nil {
			break
		}
		res = append(res, MapBookResToDto(book))
	}
	return res
}

func mapIntToInt32(i *int) *int32 {
	if i == nil {
		return nil
	}
	val := int32(*i)
	return &val
}

func mapInt32ToInt(i *int32) *int {
	if i == nil {
		return nil
	}
	val := int(*i)
	return &val
}

func pointerToInt32(i int32) *int32 {
	return &i
}

func mapTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	if err := ts.CheckValid(); err != nil {
		return time.Time{}
	}
	t := ts.AsTime()
	return t
}
