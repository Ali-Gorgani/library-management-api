package grpc

import (
	"library-management-api/auth-service/core/domain"
	"library-management-api/pkg/proto/auth"
)

func MapProtoHashedPasswordReqToDomainAuth(in *auth.HashedPasswordReq) domain.Auth {
	return domain.Auth{
		Password: in.Password,
	}
}

func MapDomainAuthToProtoHashedPasswordRes(res domain.Auth) *auth.HashedPasswordRes {
	return &auth.HashedPasswordRes{
		HashedPassword: res.Password,
	}
}

func MapProtoVerifyTokenReqToDomainAuth(in *auth.VerifyTokenReq) domain.Auth {
	return domain.Auth{
		AccessToken: in.Token,
	}
}

func MapDomainAuthToProtoVerifyTokenRes(res domain.Auth) *auth.VerifyTokenRes {
	return &auth.VerifyTokenRes{
		Id:       int32(res.Claims.ID),
		Username: res.Claims.Username,
		Email:    res.Claims.Email,
		IsAdmin:  res.Claims.IsAdmin,
		Duration: int64(res.Claims.Duration),
	}
}
