package routes

import (
	"library-management-api/api-gateway/middleware"
	"library-management-api/auth-service/api/http"

	"github.com/gin-gonic/gin"
)

var authController *http.AuthController

func AuthRoutes(r *gin.Engine) {
	authController = http.NewAuthController()

	r.POST("/login", authController.Login)
	r.POST("/logout", middleware.UserAuthMiddleware(), authController.Logout)

	tokensGroup := r.Group("/tokens", middleware.UserAuthMiddleware())
	{
		tokensGroup.POST("/refresh-token", authController.RefreshToken)
		tokensGroup.POST("/revoke-token", authController.RevokeToken)
	}
}
