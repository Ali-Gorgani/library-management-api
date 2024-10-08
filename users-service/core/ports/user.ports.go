package ports

import (
	"context"
	"library-management-api/users-service/core/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUserByUsername(ctx context.Context, user domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, user domain.User) error
}
