package repository

import (
	"context"
	"database/sql"
	"errors"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"library-management-api/users-service/init/database"
)

var (
	ErrSessionNotFound  = errors.New("session not found")
	ErrSessionRevoked   = errors.New("session is revoked")
	ErrInvalidSession   = errors.New("session is invalid")
	ErrMissingSessionID = errors.New("missing session ID")
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository() ports.SessionRepository {
	return &SessionRepository{
		db: database.P().DB,
	}
}

// CreateSession implements ports.SessionRepository.
func (s SessionRepository) CreateSession(ctx context.Context, session *domain.Session) (*domain.Session, error) {
	query := "INSERT INTO sessions (id, username, user_email, refresh_token, is_revoked, created_at, expires_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := s.db.Exec(query, session.ID, session.Username, session.UserEmail, session.RefreshToken, session.IsRevoked, session.CreatedAt, session.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetSession implements ports.SessionRepository.
func (s SessionRepository) GetSession(ctx context.Context, id string) (*domain.Session, error) {
	var session domain.Session
	query := "SELECT * FROM sessions WHERE id = $1"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&session.ID, &session.Username, &session.UserEmail, &session.RefreshToken, &session.IsRevoked, &session.CreatedAt, &session.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

// RevokeSession implements ports.SessionRepository.
func (s SessionRepository) RevokeSession(ctx context.Context, id string) error {
	query := "UPDATE sessions SET is_revoked = true WHERE id = $1"
	_, err := s.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrSessionNotFound
		}
		return err
	}
	return nil
}

// DeleteSession implements ports.SessionRepository.
func (s SessionRepository) DeleteSession(ctx context.Context, id string) error {
	query := "DELETE FROM sessions WHERE id = $1"
	_, err := s.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrSessionNotFound
		}
		return err
	}
	return nil
}
