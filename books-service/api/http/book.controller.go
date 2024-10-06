package http

import (
	"errors"
	"library-management-api/books-service/core/usecase"
	"library-management-api/util/errorhandler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookUseCase *usecase.BookUseCase
}

func NewBookController() *BookController {
	return &BookController{
		bookUseCase: usecase.NewBookUseCase(),
	}
}

func (bc *BookController) AddBook(c *gin.Context) {
	var book AddBookReq

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	addedBook, err := bc.bookUseCase.AddBook(c.Request.Context(), MapDtoAddBookReqToDomainBook(book))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		}
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}
	res := MapDomainBookToDtoBookRes(addedBook)
	c.JSON(http.StatusCreated, res)
}

func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.bookUseCase.GetBooks(c.Request.Context())
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		}
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapDomainBooksToDtoBooksRes(books)
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) GetBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	getBookReq := GetBookReq{
		ID: uint(bookID),
	}

	foundBook, err := bc.bookUseCase.GetBook(c.Request.Context(), MapDtoGetBookReqToDomainBook(getBookReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
			return
		} else if errors.Is(err, errorhandler.ErrBookNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrBookNotFound))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
			return
		}
	}
	res := MapDomainBookToDtoBookRes(foundBook)
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	var updateBookReq UpdateBookReq
	if err := c.ShouldBindJSON(&updateBookReq); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}
	updateBookReq.ID = uint(bookID)

	updatedBook, err := bc.bookUseCase.UpdateBook(c.Request.Context(), MapDtoUpdateBookReqToDomainBook(updateBookReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
			return
		} else if errors.Is(err, errorhandler.ErrBookNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrBookNotFound))
			return
		} else {
			c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
			return
		}
	}
	res := MapDomainBookToDtoBookRes(updatedBook)
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	deleteBookReq := DeleteBookReq{
		ID: uint(bookID),
	}

	err = bc.bookUseCase.DeleteBook(c.Request.Context(), MapDtoDeleteBookReqToDomainBook(deleteBookReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
			return
		} else if errors.Is(err, errorhandler.ErrBookNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrBookNotFound))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
			return
		}
	}

	c.JSON(http.StatusOK, nil)
}

func (bc *BookController) BorrowBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	borrowBookReq := BorrowBookReq{
		ID: uint(bookID),
	}

	borrowedBook, err := bc.bookUseCase.BorrowBook(c.Request.Context(), MapDtoBorrowBookReqToDomainBook(borrowBookReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
			return
		} else if errors.Is(err, errorhandler.ErrBookNotFound) {
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookNotFound))
			return
		} else if errors.Is(err, errorhandler.ErrBookAlreadyBorrowed) {
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookAlreadyBorrowed))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
			return
		}
	}
	res := MapDomainBookToDtoBookRes(borrowedBook)
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) ReturnBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	returnBookReq := ReturnBookReq{
		ID: uint(bookID),
	}

	returnedBook, err := bc.bookUseCase.ReturnBook(c.Request.Context(), MapDtoReturnBookReqToDomainBook(returnBookReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
			return
		} else if errors.Is(err, errorhandler.ErrBookNotFound) {
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookNotFound))
			return
		} else if errors.Is(err, errorhandler.ErrBookAlreadyAvailable) {
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBookAlreadyAvailable))
			return
		} else if errors.Is(err, errorhandler.ErrBorrowerIDMismatch) {
			c.JSON(http.StatusConflict, errorhandler.ErrorResponse(http.StatusConflict, errorhandler.ErrBorrowerIDMismatch))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
			return
		}
	}
	res := MapDomainBookToDtoBookRes(returnedBook)
	c.JSON(http.StatusOK, res)
}

// SearchBooks handles GET requests for searching books by title, author, or category
func (bc *BookController) SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	category := c.Query("category")

	searchBooksReq := SearchBooksReq{
		Title:    title,
		Author:   author,
		Category: category,
	}

	// Call the service layer with all non-empty query parameters
	books, err := bc.bookUseCase.SearchBooks(c.Request.Context(), MapDtoSearchBooksReqToDomainBook(searchBooksReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		} else if errors.Is(err, errorhandler.ErrInvalidSearchQuery) {
			c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, errorhandler.ErrInvalidSearchQuery))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
	}
	res := MapDomainBooksToDtoBooksRes(books)
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) CategoryBooks(c *gin.Context) {
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

	categoryBooksReq := CategoryBooksReq{
		CategoryType:  categoryType,
		CategoryValue: categoryValue,
	}

	books, err := bc.bookUseCase.CategoryBooks(c.Request.Context(), MapDtoCategoryBooksReqToDomainBook(categoryBooksReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		}
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapDomainBooksToDtoBooksRes(books)
	c.JSON(http.StatusOK, res)
}

func (bc *BookController) AvailableBooks(c *gin.Context) {
	books, err := bc.bookUseCase.AvailableBooks(c.Request.Context())
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		}
		c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		return
	}
	res := MapDomainBooksToDtoBooksRes(books)
	c.JSON(http.StatusOK, res)
}
