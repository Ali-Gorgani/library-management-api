package http

import "library-management-api/books-service/core/domain"

func MapBookToBookRes(book *domain.Book) *BookRes {
	return &BookRes{
		ID:            book.ID,
		Title:         book.Title,
		Author: 	  book.Author,
		Category:      book.Category,
		Subject:       book.Subject,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
		Available:     book.Available,
		BorrowerID:    book.BorrowerID,
		CreatedAt:     book.CreatedAt,
	}
}

func MapBooksToBooksRes(books []*domain.Book) []*BookRes {
	var booksRes []*BookRes
	for _, book := range books {
		booksRes = append(booksRes, MapBookToBookRes(book))
	}
	return booksRes
}

func MapAddBookReqToBook(req *AddBookReq) *domain.Book {
	return &domain.Book{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
	}
}

func MapGetBookReqToBook(req *GetBookReq) *domain.Book {
	return &domain.Book{
		ID: req.BookID,
	}
}

func MapUpdateBookReqToBook(req *UpdateBookReq) *domain.Book {
	return &domain.Book{
		ID:            req.BookID,
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
		Available:     req.Available,
		BorrowerID:    req.BorrowerID,
	}
}

func MapDeleteBookReqToBook(req *DeleteBookReq) *domain.Book {
	return &domain.Book{
		ID: req.BookID,
	}
}

func MapBorrowBookReqToBook(req *BorrowBookReq) *domain.Book {
	return &domain.Book{
		ID:         req.BookID,
		BorrowerID: req.BorrowerID,
	}
}

func MapReturnBookReqToBook(req *ReturnBookReq) *domain.Book {
	return &domain.Book{
		ID:         req.BookID,
		BorrowerID: req.BorrowerID,
	}
}

func MapSearchBooksReqToBook(req *SearchBooksReq) *domain.Book {
	return &domain.Book{
		Title:    req.Title,
		Author:   req.Author,
		Category: req.Category,
	}
}

func MapCategoryBooksReqToBook(req *CategoryBooksReq) *domain.Book {
	if req.CategoryType == "subject" {
		return &domain.Book{
			Subject: req.CategoryValue,
		}
	}
	return &domain.Book{
		Genre: req.CategoryValue,
	}
}
