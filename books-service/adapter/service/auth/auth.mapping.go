package auth

import (
	"library-management-api/books-service/core/domain"
	"library-management-api/books-service/third-party/auth"
)

func MapDomainHashedPasswordReqToDtoHashedPasswordReq(domain domain.Auth) auth.HashedPasswordReq {
	return auth.HashedPasswordReq{
		Password: domain.Password,
	}
}

func MapDtoHashedPasswordResToDomainHashedPasswordRes(dto auth.HashedPasswordRes) domain.Auth {
	return domain.Auth{
		Password: dto.HashedPassword,
	}
}

func MapDomainVerifyTokenReqToDtoVerifyTokenReq(domain domain.Auth) auth.VerifyTokenReq {
	return auth.VerifyTokenReq{
		Token: domain.AccessToken,
	}
}

func MapDtoVerifyTokenResToDomainVerifyTokenRes(dto auth.VerifyTokenRes) domain.Auth {
	return domain.Auth{
		Claims: domain.Claims{
			ID:       dto.ID,
			Username: dto.Username,
			Email:    dto.Email,
			IsAdmin:  dto.IsAdmin,
			Duration: dto.Duration,
		},
	}
}
