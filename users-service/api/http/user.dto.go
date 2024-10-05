package http

import "time"

type AddUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserRes struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUsersReq struct{}

type GetUserReq struct {
	ID uint
}

type UpdateUserReq struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type DeleteUserReq struct {
	ID uint
}

type DeleteUserRes struct{}
