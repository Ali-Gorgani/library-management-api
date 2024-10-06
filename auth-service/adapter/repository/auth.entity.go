package repository

import (
	"database/sql"
	"library-management-api/auth-service/core/domain"
)

type Auth struct {
	ID           uint
	UserID       uint
	RefreshToken sql.NullString
	IsRevoked    sql.NullBool
	CreatedAt    sql.NullTime
	ExpiresAt    sql.NullTime
}

func MapAuthEntityToAuthDomain(auth Auth) domain.Auth {
	return domain.Auth{
		RefreshTokenID:        auth.ID,
		RefreshTokenUserID:    auth.UserID,
		RefreshToken:          auth.RefreshToken.String,
		RefreshTokenIsRevoked: auth.IsRevoked.Bool,
		RefreshTokenCreatedAt: auth.CreatedAt.Time,
		RefreshTokenExpiresAt: auth.ExpiresAt.Time,
	}
}

func MapAuthDomainToAuthEntity(auth domain.Auth) Auth {
	return Auth{
		ID:           auth.RefreshTokenID,
		UserID:       auth.RefreshTokenUserID,
		RefreshToken: sql.NullString{String: auth.RefreshToken, Valid: auth.RefreshToken != ""},
		IsRevoked:    sql.NullBool{Bool: auth.RefreshTokenIsRevoked, Valid: true},
		CreatedAt:    sql.NullTime{Time: auth.RefreshTokenCreatedAt, Valid: !auth.RefreshTokenCreatedAt.IsZero()},
		ExpiresAt:    sql.NullTime{Time: auth.RefreshTokenExpiresAt, Valid: !auth.RefreshTokenExpiresAt.IsZero()},
	}
}
