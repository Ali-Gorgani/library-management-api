package errorhandler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrForbidden          = errors.New("you are not allowed to access this resource")
)

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionRevoked  = errors.New("session is revoked")
	ErrInvalidSession  = errors.New("session is invalid")
)

var (
	ErrBookNotFound         = errors.New("book not found")
	ErrBookAlreadyBorrowed  = errors.New("book is already borrowed")
	ErrBookAlreadyAvailable = errors.New("book is already available")
	ErrBorrowerIDMismatch   = errors.New("borrower ID does not match")
	ErrInvalidCategoryType  = errors.New("invalid category type: must be one of 'subject' or 'genre'")
	ErrEmptyCategoryValue   = errors.New("category value cannot be empty")
	ErrInvalidSearchQuery   = errors.New("at least one of the fields must be provided")
)

func ErrorResponse(status int, err error) gin.H {
	return gin.H{
		"status": status,
		"error":  err.Error(),
	}
}
