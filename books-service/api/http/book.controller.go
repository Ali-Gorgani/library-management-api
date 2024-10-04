package http

import (
	"library-management-api/auth-service/pkg/token"
	"library-management-api/books-service/core/usecase"
	"library-management-api/util/errorhandler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookUseCase *usecase.BookUsecase
}

func NewBookController() *BookController {
	return &BookController{
		bookUseCase: usecase.NewBookUseCase(),
	}
}

func (b *BookController) AddBook(c *gin.Context) {
	var book AddBookReq

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	addedBook, err := b.bookUseCase.AddBook(c.Request.Context(), MapAddBookReqToBook(&book))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}
	res := MapBookToBookRes(addedBook)
	c.JSON(http.StatusCreated, res)
}

func (b *BookController) GetBooks(c *gin.Context) {
	books, err := b.bookUseCase.GetBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapBooksToBooksRes(books)
	c.JSON(http.StatusOK, res)
}

func (b *BookController) GetBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	getBookReq := &GetBookReq{
		ID: bookID,
	}

	foundBook, err := b.bookUseCase.GetBook(c.Request.Context(), MapGetBookReqToBook(getBookReq))
	if err != nil {
		if err.Error() == errorhandler.ErrBookNotFound.Error() {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrBookNotFound))
			return
		}
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapBookToBookRes(foundBook)
	c.JSON(http.StatusOK, res)
}

func (b *BookController) UpdateBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	var book UpdateBookReqToBind
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	updateBookReq := &UpdateBookReq{
		ID:            bookID,
		Title:         book.Title,
		Author:        book.Author,
		Category:      book.Category,
		Subject:       book.Subject,
		Genre:         book.Genre,
		PublishedYear: book.PublishedYear,
		Available:     book.Available,
		BorrowerID:    book.BorrowerID,
	}

	updatedBook, err := b.bookUseCase.UpdateBook(c.Request.Context(), MapUpdateBookReqToBook(updateBookReq))
	if err != nil {
		if err.Error() == errorhandler.ErrBookNotFound.Error() {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrBookNotFound))
			return
		}
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}
	res := MapBookToBookRes(updatedBook)
	c.JSON(http.StatusOK, res)
}

func (b *BookController) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	deleteBookReq := &DeleteBookReq{
		ID: bookID,
	}

	err = b.bookUseCase.DeleteBook(c.Request.Context(), MapDeleteBookReqToBook(deleteBookReq))
	if err != nil {
		if err.Error() == errorhandler.ErrBookNotFound.Error() {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrBookNotFound))
			return
		}
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (b *BookController) BorrowBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	claims := c.Value("authKey").(*token.UserClaims)

	borrowBook := &BorrowBookReq{
		ID:         bookID,
		BorrowerID: &claims.ID,
	}

	borrowedBook, err := b.bookUseCase.BorrowBook(c.Request.Context(), MapBorrowBookReqToBook(borrowBook))
	if err != nil {
		switch err.Error() {
		case errorhandler.ErrBookNotFound.Error():
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookNotFound))
		case errorhandler.ErrBookAlreadyBorrowed.Error():
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookAlreadyBorrowed))
		default:
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapBookToBookRes(borrowedBook)
	c.JSON(http.StatusOK, res)
}

func (b *BookController) ReturnBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	claims := c.Value("authKey").(*token.UserClaims)

	returnBook := &ReturnBookReq{
		ID:         bookID,
		BorrowerID: &claims.ID,
	}

	returnedBook, err := b.bookUseCase.ReturnBook(c.Request.Context(), MapReturnBookReqToBook(returnBook))
	if err != nil {
		switch err.Error() {
		case errorhandler.ErrBookNotFound.Error():
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookNotFound))
		case errorhandler.ErrBookAlreadyAvailable.Error():
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookAlreadyAvailable))
		case errorhandler.ErrBorrowerIDMismatch.Error():
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBorrowerIDMismatch))
		default:
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapBookToBookRes(returnedBook)
	c.JSON(http.StatusOK, res)
}

// SearchBooks handles GET requests for searching books by title, author, or category
func (b *BookController) SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	category := c.Query("category")

	// Check if at least one of the fields is provided
	if title == "" && author == "" && category == "" {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, errorhandler.ErrInvalidSearchQuery))
		return
	}

	searchBooksReq := &SearchBooksReq{
		Title:    title,
		Author:   author,
		Category: category,
	}

	// Call the service layer with all non-empty query parameters
	books, err := b.bookUseCase.SearchBooks(c.Request.Context(), MapSearchBooksReqToBook(searchBooksReq))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapBooksToBooksRes(books)
	c.JSON(http.StatusOK, res)
}

func (b *BookController) CategoryBooks(c *gin.Context) {
	categoryType := c.Query("type")
	categoryValue := c.Query("value")

	if categoryType == "" || (categoryType != "subject" && categoryType != "genre") {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, errorhandler.ErrInvalidCategoryType))
		return
	}

	if categoryValue == "" {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, errorhandler.ErrEmptyCategoryValue))
		return
	}

	categoryBooksReq := &CategoryBooksReq{
		CategoryType:  categoryType,
		CategoryValue: categoryValue,
	}

	books, err := b.bookUseCase.CategoryBooks(c.Request.Context(), MapCategoryBooksReqToBook(categoryBooksReq))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapBooksToBooksRes(books)
	c.JSON(http.StatusOK, res)
}

func (b *BookController) AvailableBooks(c *gin.Context) {
	books, err := b.bookUseCase.AvailableBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapBooksToBooksRes(books)
	c.JSON(http.StatusOK, res)
}
