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
}
