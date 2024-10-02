package main

import (
	"library-management-api/books-service/api/pb"
	"library-management-api/books-service/core/usecase"
	"library-management-api/books-service/init/database"
	"library-management-api/books-service/init/migrations"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:8081", opts...)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to users-service")
		return err
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	buc := usecase.NewBookUseCase(client)
	r := setupRouter(buc)

	log.Info().Msg("Starting Books Service on :8082")
	err = http.ListenAndServe(":8082", r)
	if err != nil {
		return err
	}

	return nil
}

func setupRouter(buc *usecase.BookUsecase) *gin.Engine {
	r := gin.Default()

	r.POST("/books", buc.AddBook)
	r.GET("/books", buc.GetBooks)
	r.PUT("/books/:id", buc.UpdateBook)
	r.DELETE("/books/:id", buc.DeleteBook)
	r.POST("/books/borrow", buc.BorrowBook)
	r.POST("/books/return", buc.ReturnBook)
	r.GET("/books/search", buc.SearchBooks)
	r.GET("/books/category", buc.CategoryBooks)
	r.GET("/books/available", buc.AvailableBooks)

	r.NoRoute(func(c *gin.Context) {
		log.Warn().Str("path", c.Request.URL.Path).Int("status", http.StatusNotFound).Str("status_text", http.StatusText(http.StatusNotFound)).Msg("page not found")
		c.String(http.StatusNotFound, "404 page not found")
	})

	return r
}
