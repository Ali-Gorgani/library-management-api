package main

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"library-management-api/auth-service/adapter/repository"
	grpcController "library-management-api/auth-service/api/grpc"
	"library-management-api/auth-service/core/ports"
	"library-management-api/auth-service/init/database"
	"library-management-api/auth-service/init/migrations"
	"library-management-api/auth-service/pkg/proto"
	"net"
)

func main() {
	db := database.P().DB
	err := database.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		log.Fatal().Err(err).Msg("migration failed")
	}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	authController := grpcController.NewAuthController()

	srv := grpc.NewServer()

	proto.RegisterAuthServiceServer(srv, authController)

	if err = srv.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

type Server struct {
	proto.UnimplementedAuthServiceServer
	authRepository ports.AuthRepository
}

func NewServer() *Server {
	return &Server{
		authRepository: repository.NewAuthRepository(),
	}
}
