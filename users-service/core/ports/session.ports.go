package ports

import (
	"context"
	"library-management-api/users-service/core/domain"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session *domain.Session) (*domain.Session, error)
	GetSession(ctx context.Context, id string) (*domain.Session, error)
	RevokeSession(ctx context.Context, id string) error
	DeleteSession(ctx context.Context, id string) error
}