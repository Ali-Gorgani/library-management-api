package domain

import "time"

type BookRes struct {
	ID            int32     `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Category      string    `json:"category"`
	Subject       string    `json:"subject"`
	Genre         string    `json:"genre"`
	PublishedYear int32     `json:"published_year"`
	Available     bool      `json:"available"`
	BorrowerID    *int32    `json:"borrower_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type AddBookReq struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear int32  `json:"published_year"`
}

type UpdateBookReq struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear int32  `json:"published_year"`
	Available     bool   `json:"available"`
	BorrowerID    *int32 `json:"borrower_id"`
}

type BorrowBookReq struct {
	BookID     int32  `json:"book_id"`
	BorrowerID *int32 `json:"borrower_id"`
}

type ReturnBookReq struct {
	BookID     int32  `json:"book_id"`
	BorrowerID *int32 `json:"borrower_id"`
}
