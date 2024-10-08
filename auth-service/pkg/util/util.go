package util

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// HashedPassword hashes the given password and returns the hashed password or logs a fatal error using zerolog.
func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to hash password")
		return "", err
	}
	return string(hashedPassword), err
}

// ComparePassword compares the given hashed password with the plain password and returns true if they match or logs a fatal error using zerolog.
func ComparePassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Error().Err(err).Msg("failed to compare password")
		return false, err
	}
	return true, nil
}
