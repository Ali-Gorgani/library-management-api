package routes

import (
	"library-management-api/api-gateway/middleware"
	"library-management-api/books-service/api/http"

	"github.com/gin-gonic/gin"
)

var bookController *http.BookController

func BookRoutes(r *gin.Engine) {
	bookController = http.NewBookController()

	booksGroup := r.Group("/books", middleware.AuthMiddleware())
	{
		booksGroup.POST("/", bookController.AddBook)
		booksGroup.GET("/", bookController.GetBooks)
		booksGroup.GET("/:id", bookController.GetBook)
		booksGroup.PUT("/:id", bookController.UpdateBook)
		booksGroup.DELETE("/:id", bookController.DeleteBook)
		booksGroup.POST("/borrow/:id", bookController.BorrowBook)
		booksGroup.POST("/return/:id", bookController.ReturnBook)
		booksGroup.GET("/search", bookController.SearchBooks)
		booksGroup.GET("/category", bookController.CategoryBooks)
		booksGroup.GET("/available", bookController.AvailableBooks)
	}
}
