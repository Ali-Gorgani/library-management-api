package http

import "time"

type BookRes struct {
	ID            uint      `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Category      string    `json:"category"`
	Subject       string    `json:"subject"`
	Genre         string    `json:"genre"`
	PublishedYear uint      `json:"published_year"`
	Available     bool      `json:"available"`
	BorrowerID    uint      `json:"borrower_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type AddBookReq struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear uint   `json:"published_year"`
}

type GetBooksReq struct{}

type GetBookReq struct {
	ID uint
}

type UpdateBookReq struct {
	ID            uint
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear uint   `json:"published_year"`
	Available     bool   `json:"available"`
	BorrowerID    uint   `json:"borrower_id"`
}

type DeleteBookReq struct {
	ID uint
}

type BorrowBookReq struct {
	ID uint
}

type ReturnBookReq struct {
	ID uint
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
