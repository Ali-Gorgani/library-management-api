package http

import "library-management-api/books-service/core/domain"

func MapDomainBookToDtoBookRes(book domain.Book) BookRes {
	return BookRes{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		Category:      book.Category,
		Subject:       book.Subject,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
		Available:     book.Available,
		BorrowerID:    book.BorrowerID,
		CreatedAt:     book.CreatedAt,
	}
}

func MapDomainBooksToDtoBooksRes(books []domain.Book) []BookRes {
	var booksRes []BookRes
	for _, book := range books {
		booksRes = append(booksRes, MapDomainBookToDtoBookRes(book))
	}
	return booksRes
}

func MapDtoAddBookReqToDomainBook(req AddBookReq) domain.Book {
	return domain.Book{
		Title:         req.Title,
		Author:        req.Author,
		Category:      req.Category,
		Subject:       req.Subject,
		Genre:         req.Genre,
		PublishedYear: req.PublishedYear,
	}
}

func MapDtoGetBookReqToDomainBook(req GetBookReq) domain.Book {
	return domain.Book{
		ID: req.ID,
	}
}

func MapDtoUpdateBookReqToDomainBook(req UpdateBookReq) domain.Book {
	return domain.Book{
		ID:            req.ID,
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

func MapDtoDeleteBookReqToDomainBook(req DeleteBookReq) domain.Book {
	return domain.Book{
		ID: req.ID,
	}
}

func MapDtoBorrowBookReqToDomainBook(req BorrowBookReq) domain.Book {
	return domain.Book{
		ID: req.ID,
	}
}

func MapDtoReturnBookReqToDomainBook(req ReturnBookReq) domain.Book {
	return domain.Book{
		ID: req.ID,
	}
}

func MapDtoSearchBooksReqToDomainBook(req SearchBooksReq) domain.Book {
	return domain.Book{
		Title:    req.Title,
		Author:   req.Author,
		Category: req.Category,
	}
}

func MapDtoCategoryBooksReqToDomainBook(req CategoryBooksReq) domain.Book {
	if req.CategoryType == "subject" {
		return domain.Book{
			Subject: req.CategoryValue,
		}
	}
	return domain.Book{
		Genre: req.CategoryValue,
	}
}
