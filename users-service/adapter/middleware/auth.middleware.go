package middleware

import (
	"fmt"
	"library-management-api/users-service/pkg/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a Gin middleware that verifies the JWT token and adds claims to the context.
func UserAuthMiddleware(tokenMaker *token.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := VerifyClaimsFromAuthHeader(c, tokenMaker)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("invalid token: %s", err)})
			c.Abort()
			return
		}

		// Pass the claims to the context
		c.Set("authKey", claims)

		// Proceed to the next handler
		c.Next()
	}
}

// AdminMiddleware is a Gin middleware that checks if the user has admin privileges.
func AdminAuthMiddleware(tokenMaker *token.JWTMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := VerifyClaimsFromAuthHeader(c, tokenMaker)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("invalid token: %s", err)})
			c.Abort()
			return
		}

		if !claims.IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "user is not an admin"})
			c.Abort()
			return
		}

		// Pass the claims to the context
		c.Set("authKey", claims)

		// Proceed to the next handler
		c.Next()
	}
}

// VerifyClaimsFromAuthHeader extracts the token from the Authorization header and verifies it.
func VerifyClaimsFromAuthHeader(c *gin.Context, tokenMaker *token.JWTMaker) (*token.UserClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	fields := strings.Fields(authHeader)
	if len(fields) != 2 || strings.ToLower(fields[0]) != "bearer" {
		return nil, fmt.Errorf("invalid authorization header: %s", authHeader)
	}

	tokenString := fields[1]
	claims, err := tokenMaker.VerifyToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return claims, nil
}
