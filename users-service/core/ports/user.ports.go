package ports

import (
	"context"
	"library-management-api/users-service/core/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, user domain.AddUserReq) (domain.UserRes, error)
	GetUsers(ctx context.Context) ([]domain.UserRes, error)
	GetUser(ctx context.Context, id int) (domain.UserRes, error)
	UpdateUser(ctx context.Context, id int, user domain.UpdateUserReq) (domain.UserRes, error)
	DeleteUser(ctx context.Context, id int) error
	Login(ctx context.Context, user domain.UserLoginReq) (domain.UserRes, error)
}
