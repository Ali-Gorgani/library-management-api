package auth

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"library-management-api/users-service/pkg/proto"
)

type IClient interface {
	HashedPassword(ctx context.Context, req HashedPasswordReq) (HashedPasswordRes, error)
	VerifyToken(ctx context.Context, req VerifyTokenReq) (VerifyTokenRes, error)
}

type Client struct {
	c *grpc.ClientConn
}

func NewClient() (IClient, error) {
	client, err := grpc.NewClient("localhost:8081")
	if err != nil {
		log.Error().Err(err).Msg("failed to create grpc client")
		return nil, err
	}
	return &Client{
		c: client,
	}, nil
}

func (c *Client) HashedPassword(ctx context.Context, req HashedPasswordReq) (HashedPasswordRes, error) {
	client := proto.NewAuthServiceClient(c.c)
	res, err := client.HashedPassword(ctx, MapDtoHashedPasswordReqToPbHashedPasswordReq(req))
	if err != nil {
		log.Error().Err(err).Msg("failed to call HashedPassword")
		return HashedPasswordRes{}, err
	}
	return MapPbHashedPasswordResToDtoHashedPasswordRes(res), nil
}

func (c *Client) VerifyToken(ctx context.Context, req VerifyTokenReq) (VerifyTokenRes, error) {
	client := proto.NewAuthServiceClient(c.c)
	res, err := client.VerifyToken(ctx, MapDtoVerifyTokenReqToPbVerifyTokenReq(req))
	if err != nil {
		log.Error().Err(err).Msg("failed to call VerifyToken")
		return VerifyTokenRes{}, err
	}
	return MapPbVerifyTokenResToDtoVerifyTokenRes(res), nil
}
