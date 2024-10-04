package http

import "time"

type BookRes struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Category      string    `json:"category"`
	Subject       string    `json:"subject"`
	Genre         string    `json:"genre"`
	PublishedYear int       `json:"published_year"`
	Available     bool      `json:"available"`
	BorrowerID    *int      `json:"borrower_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type AddBookReq struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"published_year"`
}

type GetBooksReq struct{}

type GetBookReq struct {
	ID int
}

type UpdateBookReqToBind struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"published_year"`
	Available     bool   `json:"available"`
	BorrowerID    *int   `json:"borrower_id"`
}

type UpdateBookReq struct {
	ID        int
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"published_year"`
	Available     bool   `json:"available"`
	BorrowerID    *int   `json:"borrower_id"`
}

type DeleteBookReq struct {
	ID int
}

type BorrowBookReq struct {
	ID     int  `json:"book_id"`
	BorrowerID *int `json:"borrower_id"`
}

type ReturnBookReq struct {
	ID     int  `json:"book_id"`
	BorrowerID *int `json:"borrower_id"`
}

type SearchBooksReq struct {
	Title    string
	Author   string
	Category string
}

type CategoryBooksReq struct {
	CategoryType  string
	CategoryValue string
}

type AvailableBooksReq struct{}
