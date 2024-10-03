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
	ErrSessionNotFound  = errors.New("Session not found")
	ErrSessionRevoked   = errors.New("Session is revoked")
	ErrInvalidSession   = errors.New("Session is invalid")
	ErrMissingSessionID = errors.New("missing Session ID")
)

var (
	ErrBookNotFound         = errors.New("book not found")
	ErrBookAlreadyBorrowed  = errors.New("book is already borrowed")
	ErrBookAlreadyAvailable = errors.New("book is already available")
	ErrBorrowerIDMismatch   = errors.New("borrower ID does not match")
	ErrInvalidCategoryType  = errors.New("invalid category type")
	ErrEmptyCategoryValue   = errors.New("category value cannot be empty")
	ErrInvalidSearchQuery   = errors.New("invalid search query")
)

func ErrorResponse(status int, err error) gin.H {
	return gin.H{
		"status": status,
		"error":  err.Error(),
	}
}
