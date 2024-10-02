package usecase

import (
	"library-management-api/books-service/adapter/repository"
	"library-management-api/books-service/api/pb"
	"library-management-api/books-service/api/server"
	"library-management-api/books-service/core/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookUsecase struct {
	client *server.Server
}

func NewBookUseCase(client pb.BookServiceClient) *BookUsecase {
	return &BookUsecase{
		client: server.NewServer(),
	}
}

func errorResponse(statusCode int, err error) gin.H {
	return gin.H{"status": statusCode, "error": err.Error()}
}

func (u *BookUsecase) AddBook(c *gin.Context) {
	var book domain.AddBookReq
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	bookReqInProto := toProtoAddBookReq(&book)
	addedBook, err := u.client.AddBook(c.Request.Context(), bookReqInProto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}
	bookResInDomain := toDomainBook(addedBook)

	c.JSON(http.StatusCreated, bookResInDomain)
}

func (u *BookUsecase) GetBooks(c *gin.Context) {
	books, err := u.client.GetBooks(c.Request.Context(), &pb.GetBooksReq{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}
	bookResInDomain := toDomainBooks(books.Books)

	c.JSON(http.StatusOK, bookResInDomain)
}

func (u *BookUsecase) UpdateBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	var book domain.UpdateBookReq
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	bookReqInProto := toProtoUpdateBookReq(&book)
	bookReqInProto.Id = int32(bookID)
	updatedBook, err := u.client.UpdateBook(c.Request.Context(), bookReqInProto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}
	updatedBookResInDomain := toDomainBook(updatedBook)

	c.JSON(http.StatusOK, updatedBookResInDomain)
}

func (u *BookUsecase) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	_, err = u.client.DeleteBook(c.Request.Context(), &pb.DeleteBookReq{Id: int32(bookID)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (u *BookUsecase) BorrowBook(c *gin.Context) {
	var borrowBook domain.BorrowBookReq

	if err := c.ShouldBindJSON(&borrowBook); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	borrowBookReqInProto := toProtoBorrowBookReq(&borrowBook)
	borrowedBook, err := u.client.BorrowBook(c.Request.Context(), borrowBookReqInProto)
	if err != nil {
		switch err.Error() {
		case repository.ErrBookNotFound.Error():
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, repository.ErrBookNotFound))
		case repository.ErrBookAlreadyBorrowed.Error():
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, repository.ErrBookAlreadyBorrowed))
		default:
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	borrowedBookResInDomain := toDomainBook(borrowedBook)

	c.JSON(http.StatusOK, borrowedBookResInDomain)
}

func (u *BookUsecase) ReturnBook(c *gin.Context) {
	var returnBook domain.ReturnBookReq

	if err := c.ShouldBindJSON(&returnBook); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	returnBookReqInProto := toProtoReturnBookReq(&returnBook)
	returnedBook, err := u.client.ReturnBook(c.Request.Context(), returnBookReqInProto)
	if err != nil {
		switch err.Error() {
		case repository.ErrBookNotFound.Error():
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, repository.ErrBookNotFound))
		case repository.ErrBookAlreadyAvailable.Error():
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, repository.ErrBookAlreadyAvailable))
		case repository.ErrBorrowerIDMismatch.Error():
			c.JSON(http.StatusConflict, errorResponse(http.StatusConflict, repository.ErrBorrowerIDMismatch))
		default:
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	returnedBookResInDomain := toDomainBook(returnedBook)

	c.JSON(http.StatusOK, returnedBookResInDomain)
}

// SearchBooks handles GET requests for searching books by title, author, or category
func (u *BookUsecase) SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	category := c.Query("category")

	// // Check if at least one of the fields is provided
	// if title == "" && author == "" && category == "" {
	// 	c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, repository.ErrInvalidSearchQuery))
	// 	return
	// }

	// Call the service layer with all non-empty query parameters
	books, err := u.client.SearchBooks(c.Request.Context(), &pb.SearchBooksReq{
		Title:    title,
		Author:   author,
		Category: category,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	booksResInDomain := toDomainBooks(books.Books)
	c.JSON(http.StatusOK, booksResInDomain)
}


func (u *BookUsecase) CategoryBooks(c *gin.Context) {
	categoryType := c.Query("type")
	categoryValue := c.Query("value")

	if categoryType == "" || (categoryType != "subject" && categoryType != "genre") {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, repository.ErrInvalidCategoryType))
		return
	}

	if categoryValue == "" {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, repository.ErrEmptyCategoryValue))
		return
	}

	books, err := u.client.CategoryBooks(c.Request.Context(), &pb.CategoryBooksReq{CategoryType: categoryType, CategoryValue: categoryValue})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}
	booksResInDomain := toDomainBooks(books.Books)

	c.JSON(http.StatusOK, booksResInDomain)
}

func (u *BookUsecase) AvailableBooks(c *gin.Context) {
	books, err := u.client.AvailableBooks(c.Request.Context(), &pb.AvailableBooksReq{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}
	booksResInDomain := toDomainBooks(books.Books)

	c.JSON(http.StatusOK, booksResInDomain)
}
