package auth

import (
	"library-management-api/pkg/proto/auth"
	"time"
)

func MapDtoHashedPasswordReqToPbHashedPasswordReq(dto HashedPasswordReq) *auth.HashedPasswordReq {
	return &auth.HashedPasswordReq{
		Password: dto.Password,
	}
}

func MapPbHashedPasswordResToDtoHashedPasswordRes(pb *auth.HashedPasswordRes) HashedPasswordRes {
	return HashedPasswordRes{
		HashedPassword: pb.HashedPassword,
	}
}

func MapDtoVerifyTokenReqToPbVerifyTokenReq(dto VerifyTokenReq) *auth.VerifyTokenReq {
	return &auth.VerifyTokenReq{
		Token: dto.Token,
	}
}

func MapPbVerifyTokenResToDtoVerifyTokenRes(pb *auth.VerifyTokenRes) VerifyTokenRes {
	return VerifyTokenRes{
		ID:       uint(pb.ID),
		Username: pb.Username,
		Email:    pb.Email,
		IsAdmin:  pb.IsAdmin,
		Duration: time.Duration(pb.Duration),
	}
}
