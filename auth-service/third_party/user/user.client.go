package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"library-management-api/auth-service/pkg/proto"
)

type IClient interface {
	GetUserByUsername(ctx context.Context, req GetUserReq) (UserRes, error)
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

func (c *Client) GetUserByUsername(ctx context.Context, req GetUserReq) (UserRes, error) {
	client := proto.NewUsersServiceClient(c.c)
	res, err := client.GetUserByUsername(ctx, MapDtoGetUserReqToPbGetUserReq(req))
	if err != nil {
		log.Error().Err(err).Msg("failed to call GetUser")
		return UserRes{}, err
	}
	return MapPbGetUserResToDtoGetUserRes(res), nil
}
