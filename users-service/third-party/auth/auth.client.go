package auth

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"library-management-api/pkg/proto/auth"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// Client interface for AuthService
type IClient interface {
	HashedPassword(ctx context.Context, req HashedPasswordReq) (HashedPasswordRes, error)
	VerifyToken(ctx context.Context, req VerifyTokenReq) (VerifyTokenRes, error)
}

// Client struct for managing connection
type Client struct {
	c auth.AuthServiceClient // gRPC client
}

// NewClient creates a new gRPC client for AuthService
func NewClient() (IClient, error) {
	// Establish gRPC connection with the server
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("failed to create grpc client")
		return nil, err
	}
	client := auth.NewAuthServiceClient(conn)

	return &Client{
		c: client,
	}, nil
}

func (c *Client) HashedPassword(ctx context.Context, req HashedPasswordReq) (HashedPasswordRes, error) {
	res, err := c.c.HashedPassword(ctx, MapDtoHashedPasswordReqToPbHashedPasswordReq(req))
	if err != nil {
		log.Error().Err(err).Msg("failed to call HashedPassword")
		return HashedPasswordRes{}, err
	}
	return MapPbHashedPasswordResToDtoHashedPasswordRes(res), nil
}

func (c *Client) VerifyToken(ctx context.Context, req VerifyTokenReq) (VerifyTokenRes, error) {
	res, err := c.c.VerifyToken(ctx, MapDtoVerifyTokenReqToPbVerifyTokenReq(req))
	if err != nil {
		log.Error().Err(err).Msg("failed to call VerifyToken")
		return VerifyTokenRes{}, err
	}
	return MapPbVerifyTokenResToDtoVerifyTokenRes(res), nil
}
