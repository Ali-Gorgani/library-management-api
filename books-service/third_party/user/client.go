package user

import (
	"context"
	"library-management-api/books-service/pkg/user/pb"
	"library-management-api/users-service/api/http"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type IClient interface {
	AddUser(ctx context.Context, req *http.AddUserReq) (*http.UserRes, error)
	GetUsers(ctx context.Context, req *http.GetUsersReq) ([]*http.UserRes, error)
	GetUser(ctx context.Context, req *http.GetUserReq) (*http.UserRes, error)
	UpdateUser(ctx context.Context, req *http.UpdateUserReq) (*http.UserRes, error)
	DeleteUser(ctx context.Context, req *http.DeleteUserReq) error
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

func (c *Client) AddUser(ctx context.Context, req *http.AddUserReq) (*http.UserRes, error) {
	client := pb.NewUsersServiceClient(c.c)
	addedUser, err := client.AddUser(context.Background(), MapAddUserReqToPb(req))
	if err != nil {
		return &http.UserRes{}, err
	}
	res := MapUserResToDto(addedUser)
	return res, nil
}

func (c *Client) GetUsers(ctx context.Context, req *http.GetUsersReq) ([]*http.UserRes, error) {
	client := pb.NewUsersServiceClient(c.c)
	users, err := client.GetUsers(context.Background(), &pb.GetUsersReq{})
	if err != nil {
		return []*http.UserRes{}, err
	}
	res := MapUsersResToDto(users)
	return res, nil
}

func (c *Client) GetUser(ctx context.Context, req *http.GetUserReq) (*http.UserRes, error) {
	client := pb.NewUsersServiceClient(c.c)
	user, err := client.GetUser(context.Background(), MapGetUserReqToPb(req))
	if err != nil {
		return &http.UserRes{}, err
	}
	res := MapUserResToDto(user)
	return res, nil
}

func (c *Client) UpdateUser(ctx context.Context, req *http.UpdateUserReq) (*http.UserRes, error) {
	client := pb.NewUsersServiceClient(c.c)
	updatedUser, err := client.UpdateUser(context.Background(), MapUpdateUserReqToPb(req))
	if err != nil {
		return &http.UserRes{}, err
	}
	res := MapUserResToDto(updatedUser)
	return res, nil
}

func (c *Client) DeleteUser(ctx context.Context, req *http.DeleteUserReq) error {
	client := pb.NewUsersServiceClient(c.c)
	_, err := client.DeleteUser(context.Background(), MapDeleteUserReqToPb(req))
	if err != nil {
		return err
	}
	return nil
}
