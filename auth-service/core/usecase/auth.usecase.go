package usecase

import (
	"context"
	"library-management-api/auth-service/adapter/repository"
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/core/ports"
)

type AuthUsecase struct {
	AuthRepository ports.AuthRepository
}

func NewAuthUseCase() *AuthUsecase {
	return &AuthUsecase{
		AuthRepository: repository.NewAuthRepository(),
	}
}

// Login handles logic for user login
func (a *AuthUsecase) Login(ctx context.Context, auth *domain.Auth) (*domain.Auth, error) {
	// TODO: get user with auth.User.Username from users-service database
	var user domain.User
	if auth.User.Password != user.Password {
		return &domain.Auth{}, nil
	}
	auth.User = user

	token, err := a.AuthRepository.CreateToken(ctx, auth)
	if err != nil {
		return &domain.Auth{}, err
	}

	return token, nil
}

// Logout handles logic for user logout
func (a *AuthUsecase) Logout(ctx context.Context, auth *domain.Auth) error {
	err := a.AuthRepository.DeleteToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

// RefreshToken handles logic for refreshing a token
func (u *AuthUsecase) RefreshToken(ctx context.Context, auth *domain.Auth) (*domain.Auth, error) {
	refreshToken, err := u.AuthRepository.GetToken(ctx, auth)
	if err != nil {
		return &domain.Auth{}, err
	}
	if refreshToken.RefreshTokenIsRevoked {
		return &domain.Auth{}, err
	}
	newRefreshToken, err := u.AuthRepository.CreateToken(ctx, refreshToken)
	if err != nil {
		return &domain.Auth{}, err
	}
	return newRefreshToken, nil
}

// RevokeToken handles logic for revoking a token
func (u *AuthUsecase) RevokeToken(ctx context.Context, auth *domain.Auth) error {
	err := u.AuthRepository.RevokeToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}
