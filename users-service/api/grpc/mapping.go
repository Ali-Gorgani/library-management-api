package grpc

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/pkg/proto"
)

func MapProtoGetUserReqToDomainAuth(req *proto.GetUserReq) domain.User {
	return domain.User{
		Username: req.Username,
	}
}

func MapDomainAuthToProtoUserRes(user domain.User) *proto.UserRes {
	return &proto.UserRes{
		Id:        int32(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
