package util

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword hashes the given password and returns the hashed password or logs a fatal error using zerolog.
func HashedPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
	}
	return string(hashedPassword)
}

// ComparePassword compares the given hashed password with the plain password and returns true if they match or logs a fatal error using zerolog.
func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return false
	}
	return true
}
