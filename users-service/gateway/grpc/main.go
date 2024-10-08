package main

import (
	"github.com/rs/zerolog"
	"library-management-api/pkg/proto/user"
	"library-management-api/users-service/adapter/repository"
	grpcController "library-management-api/users-service/api/grpc"
	"library-management-api/users-service/core/ports"
	"library-management-api/users-service/init/database"
	"library-management-api/users-service/init/migrations"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	database.Open(database.DefaultPostgresConfig())
}

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

	srv := grpc.NewServer()
	userController := grpcController.NewUserController()
	user.RegisterUsersServiceServer(srv, userController)

	log.Info().Msgf("server started at %s", lis.Addr().String())
	if err = srv.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}

type Server struct {
	user.UnimplementedUsersServiceServer
	userRepository ports.UserRepository
}

func NewServer() *Server {
	return &Server{
		userRepository: repository.NewUserRepository(),
	}
}
