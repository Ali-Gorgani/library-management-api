package auth

import (
	"context"
	"library-management-api/auth-service/api/http"
	"library-management-api/users-service/pkg/auth/pb"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type IClient interface {
	Login(ctx context.Context, req *http.AuthLoginReq) (*http.AuthLoginRes, error)
	Logout(ctx context.Context, req *http.AuthLogoutReq) error
	RefreshToken(ctx context.Context, req *http.AuthRefreshTokenReq) (*http.AuthRefreshTokenRes, error)
	RevokeToken(ctx context.Context, req *http.AuthRevokeTokenReq) error
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

func (c *Client) Login(ctx context.Context, req *http.AuthLoginReq) (*http.AuthLoginRes, error) {
	client := pb.NewAuthServiceClient(c.c)
	loginRes, err := client.Login(context.Background(), MapDtoAuthLoginReqToPb(req))
	if err != nil {
		return &http.AuthLoginRes{}, err
	}
	res := MapPbAuthLoginResToDto(loginRes)
	return res, nil
}

func (c *Client) Logout(ctx context.Context, req *http.AuthLogoutReq) error {
	client := pb.NewAuthServiceClient(c.c)
	_, err := client.Logout(context.Background(), MapDtoAuthLogoutReqToPb(req))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RefreshToken(ctx context.Context, req *http.AuthRefreshTokenReq) (*http.AuthRefreshTokenRes, error) {
	client := pb.NewAuthServiceClient(c.c)
	refreshTokenRes, err := client.RefreshToken(context.Background(), MapDtoAuthRefreshTokenReqToPb(req))
	if err != nil {
		return &http.AuthRefreshTokenRes{}, err
	}
	res := MapPbAuthRefreshTokenResToDto(refreshTokenRes)
	return res, nil
}

func (c *Client) RevokeToken(ctx context.Context, req *http.AuthRevokeTokenReq) error {
	client := pb.NewAuthServiceClient(c.c)
	_, err := client.RevokeToken(context.Background(), MapDtoAuthRevokeTokenReqToPb(req))
	if err != nil {
		return err
	}
	return nil
}
