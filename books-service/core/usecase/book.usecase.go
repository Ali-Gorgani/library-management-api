package usecase

import (
	"library-management-api/books-service/adapter/repository"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/core/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookUsecase struct {
	bookRepository ports.BookRepository
}

func NewBookUseCase() *BookUsecase {
	return &BookUsecase{
		bookRepository: repository.NewBookRepository(),
	}
}

// errorResponse returns error details in JSON format.
func errorResponse(statusCode int, err error) gin.H {
	return gin.H{"status": statusCode, "error": err.Error()}
}

// AddBook handles POST requests for adding a new book
func (u *BookUsecase) AddBook(c *gin.Context) {
	var book domain.AddBookParam
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	addedBook, err := u.bookRepository.AddBook(c.Request.Context(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusCreated, addedBook)
}

// GetBooks handles GET requests for retrieving all books
func (u *BookUsecase) GetBooks(c *gin.Context) {
	books, err := u.bookRepository.GetBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, books)
}

// UpdateBook handles PUT requests for updating a book
func (u *BookUsecase) UpdateBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	var book domain.UpdateBookParam
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	updatedBook, err := u.bookRepository.UpdateBook(c.Request.Context(), bookID, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

// DeleteBook handles DELETE requests for deleting a book
func (u *BookUsecase) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	err = u.bookRepository.DeleteBook(c.Request.Context(), bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// BorrowBook handles POST requests for borrowing a book
func (u *BookUsecase) BorrowBook(c *gin.Context) {
	var borrowBook domain.BorrowBookRequest

	if err := c.ShouldBindJSON(&borrowBook); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	borrowedBook, err := u.bookRepository.BorrowBook(c.Request.Context(), borrowBook)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, err))
			return
		} else if err.Error() == "book is already borrowed" {
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, borrowedBook)
}

// ReturnBook handles POST requests for returning a book
func (u *BookUsecase) ReturnBook(c *gin.Context) {
	var returnBook domain.BorrowBookRequest

	if err := c.ShouldBindJSON(&returnBook); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	returnedBook, err := u.bookRepository.ReturnBook(c.Request.Context(), returnBook)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, err))
			return
		} else if err.Error() == "book is already available" {
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, err))
			return
		} else if err.Error() == "borrower ID does not match" {
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, returnedBook)
}

// SearchBooks handles GET requests for searching books
func (u *BookUsecase) SearchBooks(c *gin.Context) {
	query := c.Query("title")
	if query == "" {
		query = c.Query("author")
	}
	if query == "" {
		query = c.Query("category")
	}
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search query"})
		return
	}

	books, err := u.bookRepository.SearchBooks(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, books)
}

// CategoryBooks handles GET requests for retrieving book categories
func (u *BookUsecase) CategoryBooks(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	books, err := u.bookRepository.CategoryBooks(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, books)
}

// AvailableBooks handles GET requests for retrieving available books
func (u *BookUsecase) AvailableBooks(c *gin.Context) {
	books, err := u.bookRepository.AvailableBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, books)
}
