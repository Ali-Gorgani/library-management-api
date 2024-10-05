package domain

import "time"

type Book struct {
	ID            uint
	Title         string
	Author        string
	Category      string
	Subject       string
	Genre         string
	PublishedYear uint
	Available     bool
	BorrowerID    uint
	CreatedAt     time.Time
}
