package domain

import (
	"time"
)

type Auth struct {
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
	RefreshTokenSessionID string
	RefreshTokenIsRevoked bool
	RefreshTokenCreatedAt time.Time
	User                  User
}

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	IsAdmin   bool
	CreatedAt time.Time
}
