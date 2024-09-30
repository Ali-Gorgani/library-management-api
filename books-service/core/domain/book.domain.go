package domain

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear string `json:"published_year"`
	Available     bool   `json:"available"`
	BorrowerID    *int   `json:"borrower_id"`
	CreatedAt     string `json:"created_at"`
}

type AddBookParam struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear string `json:"published_year"`
}

type UpdateBookParam struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Category      string `json:"category"`
	Subject       string `json:"subject"`
	Genre         string `json:"genre"`
	PublishedYear string `json:"published_year"`
	Available     bool   `json:"available"`
	BorrowerID    *int   `json:"borrower_id"`
}

type BorrowBookRequest struct {
	BookID     int `json:"book_id"`
	BorrowerID *int `json:"borrower_id"`
}
