package auth

import (
	"library-management-api/books-service/pkg/proto"
	"time"
)

func MapDtoHashedPasswordReqToPbHashedPasswordReq(dto HashedPasswordReq) *proto.HashedPasswordReq {
	return &proto.HashedPasswordReq{
		Password: dto.Password,
	}
}

func MapPbHashedPasswordResToDtoHashedPasswordRes(pb *proto.HashedPasswordRes) HashedPasswordRes {
	return HashedPasswordRes{
		HashedPassword: pb.HashedPassword,
	}
}

func MapDtoVerifyTokenReqToPbVerifyTokenReq(dto VerifyTokenReq) *proto.VerifyTokenReq {
	return &proto.VerifyTokenReq{
		Token: dto.Token,
	}
}

func MapPbVerifyTokenResToDtoVerifyTokenRes(pb *proto.VerifyTokenRes) VerifyTokenRes {
	return VerifyTokenRes{
		ID:       uint(pb.ID),
		Username: pb.Username,
		Email:    pb.Email,
		IsAdmin:  pb.IsAdmin,
		Duration: time.Duration(pb.Duration),
	}
}
