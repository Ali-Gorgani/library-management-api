package grpc

import (
	"context"
	"library-management-api/auth-service/core/usecase"
	"library-management-api/pkg/proto/auth"
)

type AuthController struct {
	auth.AuthServiceServer
	authUseCase *usecase.AuthUseCase
}

func NewAuthController() *AuthController {
	return &AuthController{
		authUseCase: usecase.NewAuthUseCase(),
	}
}

func (c *AuthController) HashedPassword(ctx context.Context, in *auth.HashedPasswordReq) (*auth.HashedPasswordRes, error) {
	hashedPassword, err := c.authUseCase.HashPassword(ctx, MapProtoHashedPasswordReqToDomainAuth(in))
	if err != nil {
		return nil, err
	}
	return MapDomainAuthToProtoHashedPasswordRes(hashedPassword), nil
}

func (c *AuthController) VerifyToken(ctx context.Context, in *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	claims, err := c.authUseCase.VerifyToken(ctx, MapProtoVerifyTokenReqToDomainAuth(in))
	if err != nil {
		return nil, err
	}
	return MapDomainAuthToProtoVerifyTokenRes(claims), nil
}
