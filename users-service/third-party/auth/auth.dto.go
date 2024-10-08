package auth

import "time"

type HashedPasswordReq struct {
	Password string
}

type HashedPasswordRes struct {
	HashedPassword string
}

type VerifyTokenReq struct {
	Token string
}

type VerifyTokenRes struct {
	ID       uint
	Username string
	Email    string
	IsAdmin  bool
	Duration time.Duration
}
