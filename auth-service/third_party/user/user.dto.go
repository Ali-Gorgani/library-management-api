package user

import "time"

type GetUserReq struct {
	UserName string
}

type UserRes struct {
	ID        uint
	Username  string
	Email     string
	IsAdmin   bool
	CreatedAt time.Time
}
