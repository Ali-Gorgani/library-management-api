package http

import "time"

type UserRes struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLoginRes struct {
	SessionID             string    `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	User                  UserRes   `json:"user"`
}

type AuthLogoutReq struct {
	SessionID string
}

type AuthLogoutRes struct{}

type AuthRefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
	UserID       int
}

type AuthRefreshTokenRes struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type AuthRevokeTokenReq struct {
	SessionID string
}

type AuthRevokeTokenRes struct{}
