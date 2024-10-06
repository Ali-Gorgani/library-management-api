package http

import "library-management-api/auth-service/core/domain"

func MapDtoAuthLoginReqToDomainAuth(authLoginReq AuthLoginReq) domain.Auth {
	return domain.Auth{
		Username: authLoginReq.Username,
		Password: authLoginReq.Password,
	}
}

func MapDomainAuthToDtoAuthLoginRes(authRes domain.Auth) AuthLoginRes {
	return AuthLoginRes{
		ID:                    authRes.RefreshTokenID,
		AccessToken:           authRes.AccessToken,
		RefreshToken:          authRes.RefreshToken,
		AccessTokenExpiresAt:  authRes.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: authRes.RefreshTokenExpiresAt,
		UserID:                authRes.RefreshTokenUserID,
	}
}

func MapDtoAuthRefreshTokenReqToDomainAuth(refreshTokenReq AuthRefreshTokenReq) domain.Auth {
	return domain.Auth{
		RefreshToken: refreshTokenReq.RefreshToken,
	}
}

func MapDomainAuthToDtoAuthRefreshTokenRes(authRes domain.Auth) AuthRefreshTokenRes {
	return AuthRefreshTokenRes{
		AccessToken:          authRes.AccessToken,
		AccessTokenExpiresAt: authRes.AccessTokenExpiresAt,
	}
}

func MapDtoAuthRevokeTokenReqToDomainAuth(revokeTokenReq AuthRevokeTokenReq) domain.Auth {
	return domain.Auth{
		RefreshToken: revokeTokenReq.RefreshToken,
	}
}
