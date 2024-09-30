package usecase

import (
	"library-management-api/auth-service/adapter/repository"
	"library-management-api/auth-service/core/ports"

	"github.com/gin-gonic/gin"
)

type AuthUsecase struct {
	AuthRepository ports.AuthRepository
}

func NewAuthUseCase() *AuthUsecase {
	return &AuthUsecase{
		AuthRepository: repository.NewAuthRepository(),
	}
}

// errorResponse returns error details in JSON format.
func errorResponse(statusCode int, err error) gin.H {
	return gin.H{"status": statusCode, "error": err.Error()}
}
