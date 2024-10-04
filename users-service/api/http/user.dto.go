package http

import "time"

type AddUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserRes struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUsersReq struct{}

type GetUserReq struct {
	ID int
}

type UpdateUserReqToBind struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserReq struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type DeleteUserReq struct {
	ID int
}

type DeleteUserRes struct{}
