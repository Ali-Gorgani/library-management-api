package user

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"library-management-api/pkg/proto/user"
)

// Client interface for UserService
type IClient interface {
	GetUserByUsername(ctx context.Context, req GetUserReq) (UserRes, error)
}

// Client struct for managing connection
type Client struct {
	c user.UsersServiceClient // gRPC client
}

// NewClient creates a new gRPC client for AuthService
func NewClient() (IClient, error) {
	// Establish gRPC connection with the server
	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("failed to create grpc client")
		return nil, err
	}
	client := user.NewUsersServiceClient(conn)

	return &Client{
		c: client,
	}, nil
}

func (c *Client) GetUserByUsername(ctx context.Context, req GetUserReq) (UserRes, error) {
	res, err := c.c.GetUserByUsername(ctx, MapDtoGetUserReqToPbGetUserReq(req))
	if err != nil {
		log.Error().Err(err).Msg("failed to call GetUserByUsername")
		return UserRes{}, err
	}
	return MapPbGetUserResToDtoGetUserRes(res), nil
}
