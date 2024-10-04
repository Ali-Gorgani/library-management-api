package service

import (
	"context"
	"library-management-api/auth-service/core/domain"
	"library-management-api/users-service/third-party/auth"
	"github.com/rs/zerolog/log"
)

type AuthService struct {
	c auth.IClient
}

func NewAuthService() *AuthService {
	c, err := auth.NewClient()
	if err != nil {
		log.Error().Err(err).Msg("failed to create auth grpc client")
		return nil
	}
	return &AuthService{
		c: c,
	}
}

func (as *AuthService) Login(ctx context.Context, req *domain.Auth) (*domain.Auth, error) {
	auth, err := as.c.Login(ctx, MapDomainAuthToLoginReq(req))
	if err != nil {
		return &domain.Auth{}, err
	}
	res := MapLoginResToDomainAuth(auth)
	return res, nil
}

func (as *AuthService) Logout(ctx context.Context, req *domain.Auth) error {
	err := as.c.Logout(ctx, MapDomainAuthToLogoutReq(req))
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) RefreshToken(ctx context.Context, req *domain.Auth) (*domain.Auth, error) {
	auth, err := as.c.RefreshToken(ctx, MapDomainAuthToRefreshTokenReq(req))
	if err != nil {
		return &domain.Auth{}, err
	}
	res := MapRefreshTokenResToDomainAuth(auth)
	return res, nil
}

func (as *AuthService) RevokeToken(ctx context.Context, req *domain.Auth) error {
	err := as.c.RevokeToken(ctx, MapDomainAuthToRevokeTokenReq(req))
	if err != nil {
		return err
	}
	return nil
}
