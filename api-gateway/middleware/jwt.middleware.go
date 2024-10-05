package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware that verifies the token from the Authorization header.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := VerifyClaimsFromAuthHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
		}

		// Pass the token to the context
		c.Set("token", token)

		// Proceed to the next handler
		c.Next()
	}
}

// VerifyClaimsFromAuthHeader verifies the token from the Authorization header.
func VerifyClaimsFromAuthHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is missing")
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
		return "", fmt.Errorf("invalid authorization header: %s", authHeader)
	}
	tokenString := fields[1]

	return tokenString, nil
}
