package http

import "time"

type AuthLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLoginRes struct {
	ID                    uint      `json:"id"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	UserID                uint      `json:"user_id"`
}

type AuthLogoutReq struct{}

type AuthLogoutRes struct{}

type AuthRefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthRefreshTokenRes struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type AuthRevokeTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthRevokeTokenRes struct{}
