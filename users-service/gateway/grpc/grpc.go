package grpc

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"library-management-api/pkg/proto/user"
	grpcController "library-management-api/users-service/api/grpc"
	"net"
)

func RunGRPC() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	srv := grpc.NewServer()
	userController := grpcController.NewUserController()
	user.RegisterUsersServiceServer(srv, userController)

	log.Info().Msgf("server started at %s", lis.Addr().String())
	if err = srv.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
