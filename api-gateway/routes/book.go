package routes

import (
	"library-management-api/api-gateway/middleware"
	"library-management-api/books-service/api/http"

	"github.com/gin-gonic/gin"
)

var bookController *http.BookController

func BookRoutes(r *gin.Engine) {
	bookController = http.NewBookController()

	booksGroup := r.Group("/books")
	{
		userAccessGroup := booksGroup.Group("/", middleware.UserAuthMiddleware())
		{
			userAccessGroup.POST("/", bookController.AddBook)
			userAccessGroup.GET("/", bookController.GetBooks)
			userAccessGroup.POST("/borrow/:id", bookController.BorrowBook)
			userAccessGroup.POST("/return/:id", bookController.ReturnBook)
			userAccessGroup.GET("/search", bookController.SearchBooks)
			userAccessGroup.GET("/category", bookController.CategoryBooks)
			userAccessGroup.GET("/available", bookController.AvailableBooks)
		}

		adminAccessGroup := booksGroup.Group("/", middleware.AdminAuthMiddleware())
		{
			adminAccessGroup.PUT("/:id", bookController.UpdateBook)
			adminAccessGroup.DELETE("/:id", bookController.DeleteBook)
		}
	}
}
