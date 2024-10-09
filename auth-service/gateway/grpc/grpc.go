package grpc

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	grpcController "library-management-api/auth-service/api/grpc"
	"library-management-api/pkg/proto/auth"
	"net"
)

func RunGRPC() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	srv := grpc.NewServer()
	authController := grpcController.NewAuthController()
	auth.RegisterAuthServiceServer(srv, authController)

	log.Info().Msgf("server started at %s", lis.Addr().String())
	if err = srv.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
