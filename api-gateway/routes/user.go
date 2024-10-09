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
		usersGroupWithMW := usersGroup.Use(middleware.AuthMiddleware())
		usersGroupWithMW.GET("/", userController.GetUsers)
		usersGroupWithMW.GET("/:id", userController.GetUserByID)
		usersGroupWithMW.PUT("/:id", userController.UpdateUser)
		usersGroupWithMW.DELETE("/:id", userController.DeleteUser)
	}
}
