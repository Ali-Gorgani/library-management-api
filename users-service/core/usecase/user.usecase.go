package usecase

import (
	"errors"
	"library-management-api/users-service/adapter/repository"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/core/ports"
	"library-management-api/users-service/pkg/token"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserUsecase struct {
	userRepository    ports.UserRepository
	sessionRepository ports.SessionRepository
	TokenMaker        *token.JWTMaker
}

func NewUserUseCase(secretKey string) *UserUsecase {
	return &UserUsecase{
		userRepository:    repository.NewUserRepository(),
		sessionRepository: repository.NewSessionRepository(),
		TokenMaker:        token.NewJWTMaker(secretKey),
	}
}

// errorResponse returns error details in JSON format.
func errorResponse(statusCode int, err error) gin.H {
	return gin.H{"status": statusCode, "error": err.Error()}
}

// AddUser handles POST requests for adding a new user
func (u *UserUsecase) AddUser(c *gin.Context) {
	var user domain.AddUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	addedUser, err := u.userRepository.AddUser(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusCreated, addedUser)
}

// GetUsers handles GET requests for retrieving all users
func (u *UserUsecase) GetUsers(c *gin.Context) {
	users, err := u.userRepository.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser handles GET requests for retrieving a single user by ID
func (u *UserUsecase) GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	user, err := u.userRepository.GetUser(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse(http.StatusNotFound, repository.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser handles PUT requests for updating a user
func (u *UserUsecase) UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	claims := c.Value("authKey").(*token.UserClaims)
	if claims.ID != userID && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, errorResponse(http.StatusForbidden, repository.ErrForbidden))
		return
	}

	var user domain.UpdateUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	updatedUser, err := u.userRepository.UpdateUser(c.Request.Context(), userID, user)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse(http.StatusNotFound, repository.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser handles DELETE requests for deleting a user
func (u *UserUsecase) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	err = u.userRepository.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse(http.StatusNotFound, repository.ErrUserNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Login handles POST requests for user login
func (u *UserUsecase) Login(c *gin.Context) {
	var user domain.UserLoginReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	loggedInUser, err := u.userRepository.Login(c.Request.Context(), user)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrUserNotFound):
			c.JSON(http.StatusUnauthorized, errorResponse(http.StatusUnauthorized, repository.ErrUserNotFound))
		case errors.Is(err, repository.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, errorResponse(http.StatusUnauthorized, repository.ErrInvalidCredentials))
		default:
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	accessToken, accessClaims, err := u.TokenMaker.CreateToken(loggedInUser.ID, loggedInUser.Username, loggedInUser.Email, loggedInUser.IsAdmin, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	refreshToken, refreshClaims, err := u.TokenMaker.CreateToken(loggedInUser.ID, loggedInUser.Username, loggedInUser.Email, loggedInUser.IsAdmin, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	session, err := u.sessionRepository.CreateSession(c.Request.Context(), &domain.Session{
		ID:           refreshClaims.RegisteredClaims.ID,
		Username:     loggedInUser.Username,
		UserEmail:    loggedInUser.Email,
		RefreshToken: refreshToken,
		IsRevoked:    false,
		CreatedAt:    refreshClaims.RegisteredClaims.IssuedAt.Time.Format(time.RFC3339),
		ExpiresAt:    refreshClaims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	res := domain.UserLoginRes{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessClaims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339),
		RefreshTokenExpiresAt: refreshClaims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339),
		User:                  loggedInUser,
	}

	c.JSON(http.StatusOK, res)
}

// Logout handles DELETE requests for user logout
func (u *UserUsecase) Logout(c *gin.Context) {
	claims := c.Value("authKey").(*token.UserClaims)

	err := u.sessionRepository.DeleteSession(c.Request.Context(), claims.RegisteredClaims.ID)
	if err != nil {
		if errors.Is(err, repository.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, errorResponse(http.StatusNotFound, repository.ErrSessionNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// renewAccessToken handles POST requests for renewing access token
func (u *UserUsecase) RenewAccessToken(c *gin.Context) {
	var req domain.RenewAccessTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(http.StatusBadRequest, err))
		return
	}

	refreshClaims, err := u.TokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(http.StatusUnauthorized, err))
		return
	}

	claims := c.Value("authKey").(*token.UserClaims)
	if claims.ID != refreshClaims.ID && !claims.IsAdmin {
		c.JSON(http.StatusForbidden, errorResponse(http.StatusForbidden, repository.ErrForbidden))
		return
	}

	session, err := u.sessionRepository.GetSession(c.Request.Context(), refreshClaims.RegisteredClaims.ID)
	if err != nil {
		if errors.Is(err, repository.ErrSessionNotFound) {
			c.JSON(http.StatusUnauthorized, errorResponse(http.StatusUnauthorized, repository.ErrSessionNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	if session.IsRevoked {
		c.JSON(http.StatusUnauthorized, errorResponse(http.StatusUnauthorized, repository.ErrSessionRevoked))
		return
	}

	if session.Username != refreshClaims.Username || session.UserEmail != refreshClaims.Email {
		c.JSON(http.StatusUnauthorized, errorResponse(http.StatusUnauthorized, repository.ErrInvalidSession))
		return
	}

	accessToken, accessClaims, err := u.TokenMaker.CreateToken(refreshClaims.ID, refreshClaims.Username, refreshClaims.Email, refreshClaims.IsAdmin, 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		return
	}

	res := domain.RenewAccessTokenRes{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, res)
}

// RevokeSession handles DELETE requests for revoking a session
func (u *UserUsecase) RevokeSession(c *gin.Context) {
	claims := c.Value("authKey").(*token.UserClaims)

	err := u.sessionRepository.RevokeSession(c.Request.Context(), claims.RegisteredClaims.ID)
	if err != nil {
		if errors.Is(err, repository.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, errorResponse(http.StatusNotFound, repository.ErrSessionNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
