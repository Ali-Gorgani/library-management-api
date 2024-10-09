package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/core/ports"
	"library-management-api/auth-service/init/database"
	"library-management-api/util/errorhandler"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository() ports.AuthRepository {
	return &AuthRepository{
		db: database.P().DB,
	}
}

// CreateToken implements ports.AuthRepository.
func (a *AuthRepository) CreateToken(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	mappedAuth := MapAuthDomainToAuthEntity(auth)
	query := "INSERT INTO sessions (user_id, refresh_token, is_revoked, created_at, expires_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	row := a.db.QueryRow(query, mappedAuth.UserID, mappedAuth.RefreshToken, mappedAuth.IsRevoked, mappedAuth.CreatedAt, mappedAuth.ExpiresAt)
	err := row.Scan(&mappedAuth.ID)
	if err != nil {
		return domain.Auth{}, err
	}
	res := MapAuthEntityToAuthDomain(mappedAuth)
	res.AccessToken = auth.AccessToken
	res.AccessTokenExpiresAt = auth.AccessTokenExpiresAt
	return res, nil
}

// GetToken implements ports.AuthRepository.
func (a *AuthRepository) GetToken(ctx context.Context, auth domain.Auth) (domain.Auth, error) {
	mappedAuth := MapAuthDomainToAuthEntity(auth)
	query := "SELECT id, user_id, refresh_token, is_revoked, created_at, expires_at FROM sessions WHERE refresh_token = $1"
	row := a.db.QueryRow(query, mappedAuth.RefreshToken.String)
	err := row.Scan(&mappedAuth.ID, &mappedAuth.UserID, &mappedAuth.RefreshToken, &mappedAuth.IsRevoked, &mappedAuth.CreatedAt, &mappedAuth.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Auth{}, errorhandler.ErrSessionNotFound
		}
		return domain.Auth{}, err
	}
	res := MapAuthEntityToAuthDomain(mappedAuth)
	return res, nil
}

// RevokeToken implements ports.AuthRepository.
func (a *AuthRepository) RevokeToken(ctx context.Context, auth domain.Auth) error {
	mappedAuth := MapAuthDomainToAuthEntity(auth)
	query := "UPDATE sessions SET is_revoked = $1 WHERE refresh_token = $2"
	_, err := a.db.Exec(query, true, mappedAuth.RefreshToken.String)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorhandler.ErrSessionNotFound
		}
		return err
	}
	return nil
}

// DeleteToken implements ports.AuthRepository.
func (a *AuthRepository) DeleteToken(ctx context.Context, auth domain.Auth) error {
	mappedAuth := MapAuthDomainToAuthEntity(auth)
	query := "DELETE FROM sessions WHERE user_id = $1"
	_, err := a.db.Exec(query, mappedAuth.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorhandler.ErrSessionNotFound
		}
		return err
	}
	return nil
}
