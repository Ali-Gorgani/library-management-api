package main

import (
	"library-management-api/books-service/api/pb"
	"library-management-api/books-service/api/server"
	"library-management-api/books-service/init/database"
	"library-management-api/books-service/init/migrations"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run")
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	database.Open(database.DefaultPostgresConfig())
}

func run() error {
	db := database.P().DB
	err := database.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		log.Error().Err(err).Msg("migration failed")
		return err
	}

	srv := server.NewServer()
	grpcSrv := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcSrv, srv)

	listener, err := net.Listen("tcp", ":9082")
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return err
	}
	log.Info().Msg("Starting Books Service on :9082")
	err = grpcSrv.Serve(listener)
	if err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return err
	}

	return nil
}
