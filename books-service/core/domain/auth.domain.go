package domain

import (
	"time"
)

type Auth struct {
	RefreshTokenID        uint
	RefreshTokenUserID    uint
	RefreshToken          string
	RefreshTokenIsRevoked bool
	RefreshTokenCreatedAt time.Time
	RefreshTokenExpiresAt time.Time
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	Username              string
	Password              string
	Claims                Claims
}

type Claims struct {
	ID       uint
	Username string
	Email    string
	IsAdmin  bool
	Duration time.Duration
}
