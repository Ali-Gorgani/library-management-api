package grpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"library-management-api/pkg/proto/user"
	"library-management-api/users-service/core/domain"
)

func MapProtoGetUserReqToDomainAuth(req *user.GetUserReq) domain.User {
	return domain.User{
		Username: req.Username,
	}
}

func MapDomainAuthToProtoUserRes(res domain.User) *user.UserRes {
	return &user.UserRes{
		Id:        int32(res.ID),
		Username:  res.Username,
		Email:     res.Email,
		IsAdmin:   res.IsAdmin,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}
}
