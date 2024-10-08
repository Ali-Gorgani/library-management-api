package grpc

import (
	"library-management-api/auth-service/core/domain"
	"library-management-api/auth-service/pkg/proto"
)

func MapProtoHashedPasswordReqToDomainAuth(in *proto.HashedPasswordReq) domain.Auth {
	return domain.Auth{
		Password: in.Password,
	}
}

func MapDomainAuthToProtoHashedPasswordRes(res domain.Auth) *proto.HashedPasswordRes {
	return &proto.HashedPasswordRes{
		HashedPassword: res.Password,
	}
}

func MapProtoVerifyTokenReqToDomainAuth(in *proto.VerifyTokenReq) domain.Auth {
	return domain.Auth{
		AccessToken: in.Token,
	}
}

func MapDomainAuthToProtoVerifyTokenRes(res domain.Auth) *proto.VerifyTokenRes {
	return &proto.VerifyTokenRes{
		ID:       int32(res.Claims.ID),
		Username: res.Claims.Username,
		Email:    res.Claims.Email,
		IsAdmin:  res.Claims.IsAdmin,
		Duration: int64(res.Claims.Duration),
	}
}
