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

		adminAccessGroup := usersGroup.Group("/", middleware.AdminAuthMiddleware())
		{
			adminAccessGroup.GET("/", userController.GetUsers)
			adminAccessGroup.GET("/:id", userController.GetUser)
		}
		
		userAccessGroup := usersGroup.Group("/", middleware.UserAuthMiddleware())
		{
			userAccessGroup.PUT("/:id", userController.UpdateUser)
			adminAccessGroup.DELETE("/:id", userController.DeleteUser)
		}
	}
}
