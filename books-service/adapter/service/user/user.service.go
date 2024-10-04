package user

import (
	"context"
	"library-management-api/books-service/third_party/user"
	"library-management-api/users-service/api/http"
	"library-management-api/users-service/core/domain"

	"github.com/rs/zerolog/log"
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

func (us *UserService) AddUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	addedUser, err := us.c.AddUser(ctx, MapDomainUserToAddUserReq(req))
	if err != nil {
		return &domain.User{}, err
	}
	res := MapUserResToDomainUser(addedUser)
	return res, nil
}

func (us *UserService) GetUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := us.c.GetUsers(ctx, &http.GetUsersReq{})
	if err != nil {
		return []*domain.User{}, err
	}
	res := MapUsersResToDomainUsers(users)
	return res, nil
}

func (us *UserService) GetUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	user, err := us.c.GetUser(ctx, MapDomainUserToGetUserReq(req))
	if err != nil {
		return &domain.User{}, err
	}
	res := MapUserResToDomainUser(user)
	return res, nil
}

func (us *UserService) UpdateUser(ctx context.Context, req *domain.User) (*domain.User, error) {
	updatedUser, err := us.c.UpdateUser(ctx, MapDomainUserToUpdateUserReq(req))
	if err != nil {
		return &domain.User{}, err
	}
	res := MapUserResToDomainUser(updatedUser)
	return res, nil
}

func (us *UserService) DeleteUser(ctx context.Context, req *domain.User) error {
	err := us.c.DeleteUser(ctx, MapDomainUserToDeleteUserReq(req))
	if err != nil {
		return err
	}
	return nil
}
