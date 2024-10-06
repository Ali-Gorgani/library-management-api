package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	ID       uint
	Username string
	Email    string
	IsAdmin  bool
	Duration time.Duration
	jwt.RegisteredClaims
}

func NewUserClaims(claim UserClaims) (UserClaims, error) {
	return UserClaims{
		ID:       claim.ID,
		Username: claim.Username,
		Email:    claim.Email,
		IsAdmin:  claim.IsAdmin,
		Duration: claim.Duration,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(claim.Duration)),
		},
	}, nil
}
