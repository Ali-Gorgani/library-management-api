package grpc

import (
	"context"
	"library-management-api/users-service/core/usecase"
	"library-management-api/users-service/pkg/proto"
)

type UserController struct {
	proto.UsersServiceServer
	userUseCase *usecase.UserUseCase
}

func NewUserController() *UserController {
	return &UserController{
		userUseCase: usecase.NewUserUseCase(),
	}
}

func (c *UserController) GetUser(ctx context.Context, req *proto.GetUserReq) (*proto.UserRes, error) {
	res, err := c.userUseCase.GetUserByUsername(ctx, MapProtoGetUserReqToDomainAuth(req))
	if err != nil {
		return nil, err
	}
	return MapDomainAuthToProtoUserRes(res), nil
}
