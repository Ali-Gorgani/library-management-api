package routes

import (
	"library-management-api/api-gateway/middleware"
	"library-management-api/users-service/api/http"

	"github.com/gin-gonic/gin"
)

var userController *http.UserController

func UserRoutes(r *gin.Engine) {
	userController = http.NewUserController()
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/", userController.AddUser)
		usersGroup.Use(middleware.AuthMiddleware())
		usersGroup.GET("/", userController.GetUsers)
		usersGroup.GET("/:id", userController.GetUser)
		usersGroup.PUT("/:id", userController.UpdateUser)
		usersGroup.DELETE("/:id", userController.DeleteUser)
	}
}
