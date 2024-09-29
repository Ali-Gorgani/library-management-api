package ports

import (
	"context"
	"library-management-api/users-service/core/domain"
)

type UserService interface {
	FindByID(ctx context.Context, user domain.User) (domain.User, error)
}
