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

// UpdateBook handles PUT requests for updating a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A book has been updated"))
}

// DeleteBook handles DELETE requests for deleting a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A book has been deleted"))
}

// BorrowBook handles POST requests for borrowing a book
func BorrowBook(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A book has been borrowed"))
}

// ReturnBook handles POST requests for returning a book
func ReturnBook(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("A book has been returned"))
}
