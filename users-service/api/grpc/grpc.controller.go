package grpc

import (
	"context"
	"library-management-api/pkg/proto/user"
	"library-management-api/users-service/core/usecase"
)

type UserController struct {
	user.UsersServiceServer
	userUseCase *usecase.UserUseCase
}

func NewUserController() *UserController {
	return &UserController{
		userUseCase: usecase.NewUserUseCase(),
	}
}

func (c *UserController) GetUserByUsername(ctx context.Context, req *user.GetUserReq) (*user.UserRes, error) {
	res, err := c.userUseCase.GetUserByUsername(ctx, MapProtoGetUserReqToDomainAuth(req))
	if err != nil {
		return &user.UserRes{}, err
	}
	return MapDomainAuthToProtoUserRes(res), nil
}
