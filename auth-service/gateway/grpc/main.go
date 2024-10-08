package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"library-management-api/auth-service/adapter/repository"
	grpcController "library-management-api/auth-service/api/grpc"
	"library-management-api/auth-service/core/ports"
	"library-management-api/auth-service/init/database"
	"library-management-api/auth-service/init/migrations"
	"library-management-api/pkg/proto/auth"
	"net"
	"os"
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

type Server struct {
	auth.UnimplementedAuthServiceServer
	authRepository ports.AuthRepository
}

func NewServer() *Server {
	return &Server{
		authRepository: repository.NewAuthRepository(),
	}
}
