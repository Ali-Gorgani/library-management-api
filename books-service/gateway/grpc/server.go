package main

import (
	"library-management-api/books-service/adapter/repository"
	controller "library-management-api/books-service/api/grpc"
	"library-management-api/books-service/core/ports"
	"library-management-api/books-service/init/database"
	"library-management-api/books-service/init/migrations"
	pb "library-management-api/books-service/pkg/proto"
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
	bookController := controller.NewBookController()

	srv := grpc.NewServer()

	pb.RegisterBookServiceServer(srv, bookController)

	if err = srv.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}

}

type Server struct {
	pb.UnimplementedBookServiceServer
	bookRepository ports.BookRepository
}

func NewServer() *Server {
	return &Server{
		bookRepository: repository.NewBookRepository(),
	}
}
