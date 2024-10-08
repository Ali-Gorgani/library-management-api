package main

import (
	"library-management-api/users-service/adapter/repository"
	grpcController "library-management-api/users-service/api/grpc"
	"library-management-api/users-service/core/ports"
	"library-management-api/users-service/init/database"
	"library-management-api/users-service/init/migrations"
	"library-management-api/users-service/pkg/proto"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
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
	userController := grpcController.NewUserController()

	srv := grpc.NewServer()

	proto.RegisterUsersServiceServer(srv, userController)

	if err = srv.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}

}

type Server struct {
	proto.UnimplementedAuthServiceServer
	userRepository ports.UserRepository
}

func NewServer() *Server {
	return &Server{
		userRepository: repository.NewUserRepository(),
	}
}
