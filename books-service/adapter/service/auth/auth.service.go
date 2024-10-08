package auth

import (
	"context"
	"github.com/rs/zerolog/log"
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/third-party/auth"
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

func (s *AuthService) HashedPassword(ctx context.Context, req domain.Auth) (domain.Auth, error) {
	dtoReq := MapDomainHashedPasswordReqToDtoHashedPasswordReq(req)
	dtoRes, err := s.c.HashedPassword(ctx, dtoReq)
	if err != nil {
		log.Error().Err(err).Msg("failed to call HashedPassword")
		return domain.Auth{}, err
	}
	return MapDtoHashedPasswordResToDomainHashedPasswordRes(dtoRes), nil
}

func (s *AuthService) VerifyToken(ctx context.Context, req domain.Auth) (domain.Auth, error) {
	dtoReq := MapDomainVerifyTokenReqToDtoVerifyTokenReq(req)
	dtoRes, err := s.c.VerifyToken(ctx, dtoReq)
	if err != nil {
		log.Error().Err(err).Msg("failed to call VerifyToken")
		return domain.Auth{}, err
	}
	return MapDtoVerifyTokenResToDomainVerifyTokenRes(dtoRes), nil
}
