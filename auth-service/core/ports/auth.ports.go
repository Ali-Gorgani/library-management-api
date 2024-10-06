package ports

import (
	"context"
	"library-management-api/auth-service/core/domain"
)

type AuthRepository interface {
	CreateToken(ctx context.Context, Auth domain.Auth) (domain.Auth, error)
	GetToken(ctx context.Context, Auth domain.Auth) (domain.Auth, error)
	RevokeToken(ctx context.Context, Auth domain.Auth) error
	DeleteToken(ctx context.Context, Auth domain.Auth) error
}
