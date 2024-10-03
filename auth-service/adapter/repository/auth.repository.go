package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/core/ports"
	"library-management-api/auth-service/init/database"
	"library-management-api/auth-service/pkg/token"
	"library-management-api/util/errorhandler"
	"time"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository() ports.AuthRepository {
	return &AuthRepository{
		db: database.P().DB,
	}
}

// CreateAuth implements ports.AuthRepository.
func (s AuthRepository) CreateToken(ctx context.Context, auth *domain.Auth) (*domain.Auth, error) {
	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="
	tokenMaker := token.NewJWTMaker(secretKey)
	accessToken, accessClaims, err := tokenMaker.CreateToken(auth.User.ID, auth.User.Username, auth.User.Email, auth.User.IsAdmin, 15*time.Minute)
	if err != nil {
		return &domain.Auth{}, err
	}
	refreshToken, refreshClaims, err := tokenMaker.CreateToken(auth.User.ID, auth.User.Username, auth.User.Email, auth.User.IsAdmin, 24*time.Hour)
	if err != nil {
		return &domain.Auth{}, err
	}
	query := "INSERT INTO sessions (id, username, user_email, refresh_token, created_at, expires_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = s.db.Exec(
		query,
		refreshClaims.RegisteredClaims.ID,
		refreshClaims.Username,
		refreshClaims.Email,
		refreshToken,
		refreshClaims.RegisteredClaims.IssuedAt,
		refreshClaims.RegisteredClaims.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}

	res := &domain.Auth{
		RefreshTokenSessionID: refreshClaims.RegisteredClaims.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time,
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time,
		User:                  auth.User,
	}
	return res, nil
}

// GetAuth implements ports.AuthRepository.
func (s AuthRepository) GetToken(ctx context.Context, auth *domain.Auth) (*domain.Auth, error) {
	query := "SELECT * FROM sessions WHERE id = $1"
	row := s.db.QueryRow(query, auth.RefreshTokenSessionID)
	err := row.Scan(&auth.RefreshTokenSessionID, &auth.User.Username, &auth.User.Email, &auth.RefreshToken, &auth.RefreshTokenIsRevoked, &auth.RefreshTokenCreatedAt, &auth.RefreshTokenExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &domain.Auth{}, errorhandler.ErrSessionNotFound
		}
		return nil, err
	}

	// TODO: secret key must comes from env
	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="
	tokenMaker := token.NewJWTMaker(secretKey)
	claims, err := tokenMaker.VerifyToken(auth.RefreshToken)
	if err != nil {
		return &domain.Auth{}, err
	}

	if claims.ID != auth.User.ID {
		return &domain.Auth{}, errorhandler.ErrInvalidSession
	}

	res := &domain.Auth{
		RefreshTokenSessionID: auth.RefreshTokenSessionID,
		RefreshToken:          auth.RefreshToken,
		RefreshTokenExpiresAt: auth.RefreshTokenExpiresAt,
		RefreshTokenCreatedAt: auth.RefreshTokenCreatedAt,
		RefreshTokenIsRevoked: auth.RefreshTokenIsRevoked,
		User: domain.User{
			ID:       claims.ID,
			Username: claims.Username,
			Email:    claims.Email,
			IsAdmin:  claims.IsAdmin,
		},
	}
	return res, nil
}

// RevokeAuth implements ports.AuthRepository.
func (s AuthRepository) RevokeToken(ctx context.Context, auth *domain.Auth) error {
	query := "UPDATE sessions SET is_revoked = true WHERE id = $1"
	_, err := s.db.Exec(query, auth.RefreshTokenSessionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorhandler.ErrSessionNotFound
		}
		return err
	}
	return nil
}

// DeleteAuth implements ports.AuthRepository.
func (s AuthRepository) DeleteToken(ctx context.Context, auth *domain.Auth) error {
	query := "DELETE FROM sessions WHERE id = $1"
	_, err := s.db.Exec(query, auth.RefreshTokenSessionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorhandler.ErrSessionNotFound
		}
		return err
	}
	return nil
}
