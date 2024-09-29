package usecase

import (
	"encoding/json"
	"library-management-api/books-service/adapter/repository"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BookUsecase struct {
	bookRepository ports.BookRepository
	userService    ports.UserService
}

func NewBookUseCase() *BookUsecase {
	return &BookUsecase{
		bookRepository: repository.NewBookRepository(),
	}
}

// GetBooks handles GET requests for retrieving all books
func (u *BookUsecase) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := u.bookRepository.GetBooks(r.Context())
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode books to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// AddBook handles POST requests for adding a new book
func (u *BookUsecase) AddBook(w http.ResponseWriter, r *http.Request) {
	var book domain.AddBookParam
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	addedBook, err := u.bookRepository.AddBook(r.Context(), book)
	if err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(addedBook); err != nil {
		http.Error(w, "Failed to encode book to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateBook handles PUT requests for updating a book
func (u *BookUsecase) UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book domain.UpdateBookParam
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedBook, err := u.bookRepository.UpdateBook(r.Context(), bookID, book)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedBook); err != nil {
		http.Error(w, "Failed to encode book to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteBook handles DELETE requests for deleting a book
func (u *BookUsecase) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = u.bookRepository.DeleteBook(r.Context(), bookID)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// BorrowBook handles POST requests for borrowing a book
func (u *BookUsecase) BorrowBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	borrowedBook, err := u.bookRepository.BorrowBook(r.Context(), bookID)
	if err != nil {
		http.Error(w, "Failed to borrow book", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(borrowedBook); err != nil {
		http.Error(w, "Failed to encode book to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// ReturnBook handles POST requests for returning a book
func (u *BookUsecase) ReturnBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := chi.URLParam(r, "id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	returnedBook, err := u.bookRepository.ReturnBook(r.Context(), bookID)
	if err != nil {
		http.Error(w, "Failed to return book", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(returnedBook); err != nil {
		http.Error(w, "Failed to encode book to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// SearchBooks handles GET requests for searching books
func (u *BookUsecase) SearchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Invalid search query", http.StatusBadRequest)
		return
	}

	books, err := u.bookRepository.SearchBooks(r.Context(), query)
	if err != nil {
		http.Error(w, "Failed to search books", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode books to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CategoryBooks handles GET requests for retrieving book categories
func (u *BookUsecase) CategoryBooks(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	if category == "" {
		http.Error(w, "Invalid category", http.StatusBadRequest)
		return
	}

	books, err := u.bookRepository.CategoryBooks(r.Context(), category)
	if err != nil {
		http.Error(w, "Failed to get books by category", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode books to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// AvailableBooks handles GET requests for retrieving available books
func (u *BookUsecase) AvailableBooks(w http.ResponseWriter, r *http.Request) {
	books, err := u.bookRepository.AvailableBooks(r.Context())
	if err != nil {
		http.Error(w, "Failed to get available books", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode books to JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
