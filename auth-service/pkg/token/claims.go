package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	ID       uint          `json:"id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	IsAdmin  bool          `json:"is_admin"`
	Duration time.Duration `json:"duration"`
	jwt.RegisteredClaims
}

func NewUserClaims(claim UserClaims) (*UserClaims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &UserClaims{
		ID:       id,
		Username: username,
		Email:    email,
		IsAdmin:  isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil
}
