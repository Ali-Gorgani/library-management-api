package domain

type UserRes struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
}

type AddUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRes struct {
	SessionID             string  `json:"session_id"`
	AccessToken           string  `json:"access_token"`
	RefreshToken          string  `json:"refresh_token"`
	AccessTokenExpiresAt  string  `json:"access_token_expires_at"`
	RefreshTokenExpiresAt string  `json:"refresh_token_expires_at"`
	User                  UserRes `json:"user"`
}


