package book

import (
	"context"
	"library-management-api/users-service/core/domain"
	"library-management-api/users-service/pkg/pb"
	"log"

	"google.golang.org/grpc"
)

type IClient interface {
	AddBook(ctx context.Context, req AddBookReq) (AddBookRes, error)
}

type Client struct {
	c *grpc.ClientConn
}

func NewClient() (IClient, error) {
	client, err := grpc.NewClient("localhost:8081")
	if err != nil {
		log.Fatal()
		return nil , err
	}
	return &Client{
		c: client,
	}, nil
}

func (c *Client) AddBook(ctx context.Context, req AddBookReq) (AddBookRes, error) {
	client := pb.NewBookServiceClient(c.c)
	res, err := client.AddBook(context.Background(), &pb.AddBookReq{ // map to pb.addbookreq

	})
	if err != nil {
		return AddBookRes{}, err
	}
	return AddBookRes{}, nil  // map res to domain.books
}
