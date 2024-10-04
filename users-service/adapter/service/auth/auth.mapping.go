package service

import (
	"library-management-api/auth-service/api/http"
	"library-management-api/auth-service/core/domain"
)

func MapDomainAuthToLoginReq(req *domain.Auth) *http.AuthLoginReq {
	return &http.AuthLoginReq{
		Username: req.User.Username,
		Password: req.User.Password,
	}
}

func MapLoginResToDomainAuth(res *http.AuthLoginRes) *domain.Auth {
	return &domain.Auth{
		AccessToken:           res.AccessToken,
		AccessTokenExpiresAt:  res.AccessTokenExpiresAt,
		RefreshToken:          res.RefreshToken,
		RefreshTokenExpiresAt: res.RefreshTokenExpiresAt,
		RefreshTokenSessionID: res.SessionID,
		User: domain.User{
			ID:        res.User.ID,
			Username:  res.User.Username,
			Email:     res.User.Email,
			IsAdmin:   res.User.IsAdmin,
			CreatedAt: res.User.CreatedAt,
		},
	}
}

func MapDomainAuthToLogoutReq(req *domain.Auth) *http.AuthLogoutReq {
	return &http.AuthLogoutReq{
		SessionID: req.RefreshTokenSessionID,
	}
}

func MapDomainAuthToRefreshTokenReq(req *domain.Auth) *http.AuthRefreshTokenReq {
	return &http.AuthRefreshTokenReq{
		RefreshToken: req.RefreshToken,
		UserID:       req.User.ID,
	}
}

func MapRefreshTokenResToDomainAuth(res *http.AuthRefreshTokenRes) *domain.Auth {
	return &domain.Auth{
		AccessToken:          res.AccessToken,
		AccessTokenExpiresAt: res.AccessTokenExpiresAt,
	}
}

func MapDomainAuthToRevokeTokenReq(req *domain.Auth) *http.AuthRevokeTokenReq {
	return &http.AuthRevokeTokenReq{
		SessionID: req.RefreshTokenSessionID,
	}
}
