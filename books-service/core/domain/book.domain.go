package domain

import "time"

type Book struct {
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
