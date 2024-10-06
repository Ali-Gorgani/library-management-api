package http

import (
	"errors"
	"library-management-api/auth-service/core/usecase"
	"library-management-api/util/errorhandler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthController() *AuthController {
	return &AuthController{
		authUseCase: usecase.NewAuthUseCase(),
	}
}

// Login handles POST requests for Auth login
func (ac *AuthController) Login(c *gin.Context) {
	var req AuthLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	auth, err := ac.authUseCase.Login(c.Request.Context(), MapDtoAuthLoginReqToDomainAuth(req))
	if err != nil {
		if errors.Is(err, errorhandler.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrUserNotFound))
		} else if errors.Is(err, errorhandler.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidCredentials))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapDomainAuthToDtoAuthLoginRes(auth)
	c.JSON(http.StatusOK, res)
}

// Logout handles DELETE requests for Auth logout
func (ac *AuthController) Logout(c *gin.Context) {
	err := ac.authUseCase.Logout(c.Request.Context())
	if err != nil {
		if errors.Is(err, errorhandler.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrSessionNotFound))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// RefreshToken handles POST requests for refreshing a token
func (ac *AuthController) RefreshToken(c *gin.Context) {
	var req AuthRefreshTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	auth, err := ac.authUseCase.RefreshToken(c.Request.Context(), MapDtoAuthRefreshTokenReqToDomainAuth(req))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrSessionNotFound))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		} else if errors.Is(err, errorhandler.ErrSessionRevoked) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrSessionRevoked))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	res := MapDomainAuthToDtoAuthRefreshTokenRes(auth)
	c.JSON(http.StatusOK, res)
}

// RevokeToken handles POST requests for revoking a token
func (ac *AuthController) RevokeToken(c *gin.Context) {
	var req AuthRevokeTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	err := ac.authUseCase.RevokeToken(c.Request.Context(), MapDtoAuthRevokeTokenReqToDomainAuth(req))
	if err != nil {
		if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrSessionNotFound))
		} else if errors.Is(err, errorhandler.ErrForbidden) {
			c.JSON(http.StatusForbidden, errorhandler.ErrorResponse(http.StatusForbidden, errorhandler.ErrForbidden))
		} else if errors.Is(err, errorhandler.ErrSessionRevoked) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrSessionRevoked))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
