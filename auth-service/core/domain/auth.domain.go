package domain

import (
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateAuthParam struct {
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	ExpiresAt    time.Time
}

type Payload struct {
	ID        uuid.UUID     `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
