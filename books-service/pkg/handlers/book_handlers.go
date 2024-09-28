package handlers

import (
    "net/http"
)

// GetBooks handles GET requests for retrieving all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Here are all the books"))
}

// AddBook handles POST requests for adding a new book
func AddBook(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A new book has been added"))
}
