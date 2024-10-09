package usecase

import (
	"context"
	"github.com/rs/zerolog/log"
	"library-management-api/auth-service/adapter/repository"
	userService "library-management-api/auth-service/adapter/service/user"
	"library-management-api/auth-service/configs"
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/core/ports"
	"library-management-api/auth-service/pkg/token"
	"library-management-api/auth-service/pkg/util"
	"library-management-api/util/errorhandler"
	"time"
)

type AuthUseCase struct {
	authRepository ports.AuthRepository
	userService    *userService.UsersService
	config         configs.Config
}

func NewAuthUseCase() *AuthUseCase {
	config, err := configs.LoadConfig("auth-service")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}
	return &AuthUseCase{
		authRepository: repository.NewAuthRepository(),
		userService:    userService.NewUserService(),
		config:         config,
	}
}

// Login handles logic for user login
func (a *AuthUseCase) Login(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	user, err := a.userService.GetUserByUsername(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}

	ok, err := util.ComparePassword(user.Password, auth.Password)
	if err != nil {
		return domain.Auth{}, err
	}
	if !ok {
		return domain.Auth{}, errorhandler.ErrInvalidCredentials
	}

	auth.Claims = domain.Claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Duration: 15 * time.Minute,
	}
	accessToken, err := a.CreateToken(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}

	auth.Claims = domain.Claims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Duration: 24 * time.Hour,
	}
	refreshToken, err := a.CreateToken(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}

	auth = domain.Auth{
		RefreshTokenUserID:    user.ID,
		RefreshToken:          refreshToken.AccessToken,
		RefreshTokenIsRevoked: false,
		RefreshTokenCreatedAt: refreshToken.Claims.IssuedAt,
		RefreshTokenExpiresAt: refreshToken.AccessTokenExpiresAt,
		AccessToken:           accessToken.AccessToken,
		AccessTokenExpiresAt:  accessToken.AccessTokenExpiresAt,
	}
	auth, err = a.authRepository.CreateToken(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}
	return auth, nil
}

// Logout handles logic for user logout
func (a *AuthUseCase) Logout(ctx context.Context) error {
	contextToken, ok := ctx.Value("token").(string)
	if !ok {
		return errorhandler.ErrInvalidSession
	}

	auth := domain.Auth{
		AccessToken: contextToken,
	}
	claims, err := a.VerifyToken(ctx, auth)
	if err != nil {
		return err
	}

	auth = domain.Auth{
		RefreshTokenUserID: claims.Claims.ID,
	}
	err = a.authRepository.DeleteToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

// RefreshToken handles logic for refreshing a token
func (a *AuthUseCase) RefreshToken(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	contextToken, ok := ctx.Value("token").(string)
	if !ok {
		return domain.Auth{}, errorhandler.ErrInvalidSession
	}

	verifyTokenReq := domain.Auth{
		AccessToken: contextToken,
	}
	verifyTokenRes, err := a.VerifyToken(ctx, verifyTokenReq)
	if err != nil {
		return domain.Auth{}, err
	}
	claims := verifyTokenRes.Claims

	auth, err = a.authRepository.GetToken(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}

	if auth.RefreshTokenUserID != claims.ID {
		return domain.Auth{}, errorhandler.ErrForbidden
	}

	if auth.RefreshTokenIsRevoked {
		return domain.Auth{}, errorhandler.ErrSessionRevoked
	}

	auth.Claims = domain.Claims{
		ID:       claims.ID,
		Username: claims.Username,
		Email:    claims.Email,
		IsAdmin:  claims.IsAdmin,
		Duration: 15 * time.Minute,
	}
	newAuth, err := a.CreateToken(ctx, auth)
	if err != nil {
		return domain.Auth{}, err
	}

	return newAuth, nil
}

// RevokeToken handles logic for revoking a token
func (a *AuthUseCase) RevokeToken(ctx context.Context, auth domain.Auth) error {
	contextToken, ok := ctx.Value("token").(string)
	if !ok {
		return errorhandler.ErrInvalidSession
	}

	verifyTokenReq := domain.Auth{
		AccessToken: contextToken,
	}
	verifyTokenRes, err := a.VerifyToken(ctx, verifyTokenReq)
	if err != nil {
		return err
	}
	claims := verifyTokenRes.Claims

	auth, err = a.authRepository.GetToken(ctx, auth)
	if err != nil {
		return err
	}

	if auth.RefreshTokenUserID != claims.ID {
		return errorhandler.ErrForbidden
	}

	if auth.RefreshTokenIsRevoked {
		return errorhandler.ErrSessionRevoked
	}

	err = a.authRepository.RevokeToken(ctx, auth)
	if err != nil {
		return err
	}
	return nil
}

// HashPassword handles logic for hashing a password
func (a *AuthUseCase) HashPassword(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	hashedPassword, err := util.HashedPassword(auth.Password)
	if err != nil {
		return domain.Auth{}, err
	}
	auth.Password = hashedPassword
	return auth, nil
}

// CreateToken handles logic for creating a token
func (a *AuthUseCase) CreateToken(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	secretKey := a.config.JWT.SecretKey
	userClaims := token.UserClaims{
		ID:       auth.Claims.ID,
		Username: auth.Claims.Username,
		Email:    auth.Claims.Email,
		IsAdmin:  auth.Claims.IsAdmin,
		Duration: auth.Claims.Duration,
	}
	claims, err := token.NewUserClaims(userClaims)
	if err != nil {
		return domain.Auth{}, errorhandler.ErrInvalidSession
	}
	accessToken, err := token.CreateToken(secretKey, claims)
	if err != nil {
		return domain.Auth{}, errorhandler.ErrInvalidSession
	}
	auth = domain.Auth{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: claims.RegisteredClaims.ExpiresAt.Time,
		Claims: domain.Claims{
			ID:        claims.ID,
			Username:  claims.Username,
			Email:     claims.Email,
			IsAdmin:   claims.IsAdmin,
			Duration:  claims.Duration,
			IssuedAt:  claims.RegisteredClaims.IssuedAt.Time,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Time,
		},
	}
	return auth, nil
}

// VerifyToken handles logic for verifying a token
func (a *AuthUseCase) VerifyToken(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	secretKey := a.config.JWT.SecretKey
	claims, err := token.VerifyToken(auth.AccessToken, secretKey)
	if err != nil {
		return domain.Auth{}, errorhandler.ErrInvalidSession
	}
	auth.Claims = domain.Claims{
		ID:        claims.ID,
		Username:  claims.Username,
		Email:     claims.Email,
		IsAdmin:   claims.IsAdmin,
		Duration:  claims.Duration,
		IssuedAt:  claims.RegisteredClaims.IssuedAt.Time,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Time,
	}
	return auth, nil
}
