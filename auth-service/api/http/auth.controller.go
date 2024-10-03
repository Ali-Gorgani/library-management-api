package http

import (
	"errors"
	"library-management-api/auth-service/core/usecase"
	"library-management-api/auth-service/pkg/token"
	"library-management-api/util/errorhandler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUseCase *usecase.AuthUsecase
}

func NewAuthController() *AuthController {
	return &AuthController{
		AuthUseCase: usecase.NewAuthUseCase(),
	}
}

// Login handles POST requests for Auth login
func (ac *AuthController) Login(c *gin.Context) {
	var req AuthLoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	loggedInAuth, err := ac.AuthUseCase.Login(c.Request.Context(), MapAuthLoginReqToAuth(&req))
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
	res := MapAuthToAuthLoginRes(loggedInAuth)
	c.JSON(http.StatusOK, res)
}

// Logout handles DELETE requests for Auth logout
func (ac *AuthController) Logout(c *gin.Context) {
	claims := c.Value("authKey").(*token.UserClaims)
	logoutReq := &AuthLogoutReq{
		SessionID: claims.RegisteredClaims.ID,
	}

	err := ac.AuthUseCase.Logout(c.Request.Context(), MapAuthLogoutReqToAuth(logoutReq))
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

// RenewAccessToken handles POST requests for renewing access token
func (ac *AuthController) RefreshToken(c *gin.Context) {
	var req string
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorhandler.ErrorResponse(http.StatusBadRequest, err))
		return
	}

	claims := c.Value("authKey").(*token.UserClaims)
	refreshTokenReq := &AuthRefreshTokenReq{
		RefreshToken: req,
		UserID:       claims.ID,
	}
	refreshClaims, err := ac.AuthUseCase.RefreshToken(c.Request.Context(), MapAuthRefreshTokenReqToAuth(refreshTokenReq))
	if err != nil {
		if errors.Is(err, errorhandler.ErrSessionNotFound) {
			c.JSON(http.StatusNotFound, errorhandler.ErrorResponse(http.StatusNotFound, errorhandler.ErrSessionNotFound))
		} else if errors.Is(err, errorhandler.ErrInvalidSession) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrInvalidSession))
		} else if errors.Is(err, errorhandler.ErrSessionRevoked) {
			c.JSON(http.StatusUnauthorized, errorhandler.ErrorResponse(http.StatusUnauthorized, errorhandler.ErrSessionRevoked))
		} else {
			c.JSON(http.StatusInternalServerError, errorhandler.ErrorResponse(http.StatusInternalServerError, err))
		}
		return
	}

	res := MapAuthToAuthRefreshTokenRes(refreshClaims)
	c.JSON(http.StatusOK, res)
}

// RevokeSession handles DELETE requests for revoking a session
func (ac *AuthController) RevokeToken(c *gin.Context) {
	claims := c.Value("authKey").(*token.UserClaims)
	revokeReq := &AuthRevokeTokenReq{
		SessionID: claims.RegisteredClaims.ID,
	}

	err := ac.AuthUseCase.RevokeToken(c.Request.Context(), MapAuthRevokeTokenReqToAuth(revokeReq))
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
