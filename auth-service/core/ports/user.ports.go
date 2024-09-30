package ports

import (
	"context"
	"library-management-api/users-service/core/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, user domain.AddUserParam) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	UpdateUser(ctx context.Context, id int, user domain.UpdateUserParam) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	Login(ctx context.Context, user domain.UserLoginParam) (domain.User, error)
}
