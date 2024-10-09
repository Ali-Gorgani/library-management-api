package user

import "time"

type GetUserReq struct {
	Username string
}

type UserRes struct {
	ID        uint
	Username  string
	Password  string
	Email     string
	IsAdmin   bool
	CreatedAt time.Time
}
