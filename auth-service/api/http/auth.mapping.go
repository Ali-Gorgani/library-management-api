package http

import "library-management-api/auth-service/core/domain"

func MapAuthLoginReqToAuth(authLoginReq *AuthLoginReq) *domain.Auth {
	return &domain.Auth{
		User: domain.User{
			Username: authLoginReq.Username,
			Password: authLoginReq.Password,
		},
	}
}

func MapAuthToAuthLoginRes(authRes *domain.Auth) *AuthLoginRes {
	return &AuthLoginRes{
		SessionID:             authRes.RefreshTokenSessionID,
		AccessToken:           authRes.AccessToken,
		RefreshToken:          authRes.RefreshToken,
		AccessTokenExpiresAt:  authRes.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: authRes.RefreshTokenExpiresAt,
		User: UserRes{
			ID:        authRes.User.ID,
			Username:  authRes.User.Username,
			Email:     authRes.User.Email,
			IsAdmin:   authRes.User.IsAdmin,
			CreatedAt: authRes.User.CreatedAt,
		},
	}
}

func MapAuthLogoutReqToAuth(authLogoutReq *AuthLogoutReq) *domain.Auth {
	return &domain.Auth{
		RefreshTokenSessionID: authLogoutReq.SessionID,
	}
}

func MapAuthRefreshTokenReqToAuth(refreshTokenReq *AuthRefreshTokenReq) *domain.Auth {
	return &domain.Auth{
		RefreshToken: refreshTokenReq.RefreshToken,
		User: domain.User{
			ID: refreshTokenReq.UserID,
		},
	}
}

func MapAuthToAuthRefreshTokenRes(authRes *domain.Auth) *AuthRefreshTokenRes {
	return &AuthRefreshTokenRes{
		AccessToken:          authRes.AccessToken,
		AccessTokenExpiresAt: authRes.AccessTokenExpiresAt,
	}
}

func MapAuthRevokeTokenReqToAuth(authRevokeTokenReq *AuthRevokeTokenReq) *domain.Auth {
	return &domain.Auth{
		RefreshTokenSessionID: authRevokeTokenReq.SessionID,
	}
}
