package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/third_party/user"
)

type UserService struct {
	c user.IClient
}

func NewUserService() *UserService {
	c, err := user.NewClient()
	if err != nil {
		log.Error().Err(err).Msg("failed to create user grpc client")
		return nil
	}
	return &UserService{
		c: c,
	}
}

func (s *UserService) GetUserByUsername(ctx context.Context, req domain.Auth) (domain.User, error) {
	dtoReq := MapDomainUserToDtoGetUserReq(req)
	dtoRes, err := s.c.GetUserByUsername(ctx, dtoReq)
	if err != nil {
		log.Error().Err(err).Msg("failed to call GetUser")
		return domain.User{}, err
	}
	return MapDtoUserResToDomainUser(dtoRes), nil
}
