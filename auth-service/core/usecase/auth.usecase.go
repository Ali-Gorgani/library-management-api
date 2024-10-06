package usecase

import (
	"context"
	"library-management-api/auth-service/adapter/repository"
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/core/ports"
	"library-management-api/auth-service/pkg/token"
	"library-management-api/auth-service/pkg/util"
	"library-management-api/util/errorhandler"
	"time"
)

type AuthUseCase struct {
	AuthRepository ports.AuthRepository
}

func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{
		AuthRepository: repository.NewAuthRepository(),
	}
}

// Login handles logic for user login
func (a *AuthUseCase) Login(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	// TODO: get user with auth.Username from users-service database with gRPC
	if err != nil {
		return domain.Auth{}, err
	}

	var user domain.User
	ok := util.ComparePassword(user.Password, auth.Password)
	if !ok {
		return domain.Auth{}, nil
	}

	// TODO: get secretKey from env
	userClaims := token.UserClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Duration: 15 * time.Minute,
	}
	accessTokenClaims, err := token.NewUserClaims(userClaims)
	if err != nil {
		return domain.Auth{}, err
	}
	accessToken, err := token.CreateToken(secretKey, accessTokenClaims)
	if err != nil {
		return domain.Auth{}, err
	}

	userClaims = token.UserClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Duration: 24 * time.Hour,
	}
	refreshTokenClaims, err := token.NewUserClaims(userClaims)
	if err != nil {
		return domain.Auth{}, err
	}
	refreshToken, err := token.CreateToken(secretKey, refreshTokenClaims)
	if err != nil {
		return domain.Auth{}, err
	}

	auth = domain.Auth{
		RefreshTokenUserID:    user.ID,
		RefreshToken:          refreshToken,
		RefreshTokenIsRevoked: false,
		RefreshTokenCreatedAt: refreshTokenClaims.RegisteredClaims.IssuedAt.Time,
		RefreshTokenExpiresAt: refreshTokenClaims.RegisteredClaims.ExpiresAt.Time,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenClaims.RegisteredClaims.ExpiresAt.Time,
	}
	auth, err = a.AuthRepository.CreateToken(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}
	return auth, nil
}

// Logout handles logic for user logout
func (a *AuthUseCase) Logout(ctx context.Context) error {
	contextToken := ctx.Value("token").(string)

	// TODO: get secretKey from env
	_, err := token.VerifyToken(secretKey, contextToken)
	if err != nil {
		return err
	}

	auth := domain.Auth{
		RefreshToken: contextToken,
	}

	err = a.AuthRepository.DeleteToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

// RefreshToken handles logic for refreshing a token
func (a *AuthUseCase) RefreshToken(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	contextToken := ctx.Value("token").(string)

	// TODO: get secretKey from env
	claims, err := token.VerifyToken(secretKey, contextToken)
	if err != nil {
		return domain.Auth{}, errorhandler.ErrInvalidSession
	}

	auth, err = a.AuthRepository.GetToken(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}

	if auth.RefreshTokenUserID != claims.ID {
		return domain.Auth{}, errorhandler.ErrForbidden
	}

	if auth.RefreshTokenIsRevoked {
		return domain.Auth{}, errorhandler.ErrSessionRevoked
	}

	userClaims := token.UserClaims{
		ID:       claims.ID,
		Username: claims.Username,
		Email:    claims.Email,
		IsAdmin:  claims.IsAdmin,
		Duration: 15 * time.Minute,
	}
	newAccessTokenClaims, err := token.NewUserClaims(userClaims)
	if err != nil {
		return domain.Auth{}, err
	}
	newAccessToken, err := token.CreateToken(secretKey, newAccessTokenClaims)
	if err != nil {
		return domain.Auth{}, err
	}
	auth.AccessToken = newAccessToken
	auth.AccessTokenExpiresAt = newAccessTokenClaims.RegisteredClaims.ExpiresAt.Time

	return auth, nil
}

// RevokeToken handles logic for revoking a token
func (a *AuthUseCase) RevokeToken(ctx context.Context, auth domain.Auth) error {
	contextToken := ctx.Value("token").(string)

	// TODO: get secretKey from env
	claims, err := token.VerifyToken(secretKey, contextToken)
	if err != nil {
		return errorhandler.ErrInvalidSession
	}

	auth, err = a.AuthRepository.GetToken(ctx, auth)
	if err != nil {
		return err
	}

	if auth.RefreshTokenUserID != claims.ID {
		return errorhandler.ErrForbidden
	}

	if auth.RefreshTokenIsRevoked {
		return errorhandler.ErrSessionRevoked
	}

	err = a.AuthRepository.RevokeToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}
