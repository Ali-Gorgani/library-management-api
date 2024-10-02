package domain

type Session struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	UserEmail    string `json:"user_email"`
	RefreshToken string `json:"refresh_token"`
	IsRevoked    bool   `json:"is_revoked"`
	CreatedAt    string `json:"created_at"`
	ExpiresAt    string `json:"expires_at"`
}

type RenewAccessTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}

type RenewAccessTokenRes struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpiresAt string `json:"access_token_expires_at"`
}
