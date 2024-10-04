package service

import (
	"library-management-api/books-service/api/http"
	"library-management-api/books-service/core/domain"
)

func MapDomainBookToAddBookReq(book domain.Book) http.AddBookReq {
	return http.AddBookReq{
		Title:         book.Title,
		Author:        book.Author,
		Category:      book.Category,
		Subject:       book.Subject,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
	}
}

func MapDomainBookToGetBookReq(b domain.Book) http.GetBookReq {
	return http.GetBookReq{
		ID: b.ID,
	}
}

func MapDomainBookToUpdateBookReq(b domain.Book) http.UpdateBookReq {
	return http.UpdateBookReq{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		Category:      b.Category,
		Subject:       b.Subject,
		Genre:         b.Genre,
		PublishedYear: b.PublishedYear,
		Available:     b.Available,
		BorrowerID:    b.BorrowerID,
	}
}

func MapDomainBookToDeleteBookReq(b domain.Book) http.DeleteBookReq {
	return http.DeleteBookReq{
		ID: b.ID,
	}
}

func MapDomainBookToBorrowBookReq(b domain.Book) http.BorrowBookReq {
	return http.BorrowBookReq{
		ID:         b.ID,
		BorrowerID: b.BorrowerID,
	}
}

func MapDomainBookToReturnBookReq(b domain.Book) http.ReturnBookReq {
	return http.ReturnBookReq{
		ID:         b.ID,
		BorrowerID: b.BorrowerID,
	}
}

func MapDomainBookToSearchBooksReq(b domain.Book) http.SearchBooksReq {
	return http.SearchBooksReq{
		Title:    b.Title,
		Author:   b.Author,
		Category: b.Category,
	}
}

func MapDomainBookToCategoryBooksReq(b domain.Book) http.CategoryBooksReq {
	if b.Subject != "" {
		return http.CategoryBooksReq{
			CategoryType:  "subject",
			CategoryValue: b.Subject,
		}
	}
	return http.CategoryBooksReq{
		CategoryType:  "genre",
		CategoryValue: b.Genre,
	}
}

func MapBookResToDomainBook(b http.BookRes) domain.Book {
	return domain.Book{
		ID:            b.ID,
		Title:         b.Title,
		Author:        b.Author,
		Category:      b.Category,
		Subject:       b.Subject,
		Genre:         b.Genre,
		PublishedYear: b.PublishedYear,
		Available:     b.Available,
		BorrowerID:    b.BorrowerID,
		CreatedAt:     b.CreatedAt,
	}
}

func MapBooksResToDomainBooks(books []http.BookRes) []domain.Book {
	var domainBooks []domain.Book
	for _, b := range books {
		domainBooks = append(domainBooks, MapBookResToDomainBook(b))
	}
	return domainBooks
}
