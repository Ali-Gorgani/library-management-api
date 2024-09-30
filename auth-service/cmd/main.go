package main

import (
	"library-management-api/auth-service/core/usecase"
	"library-management-api/auth-service/init/database"
	"library-management-api/auth-service/init/migrations"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

	auc := usecase.NewAuthUseCase()
	r := setupRouter(auc)

	log.Info().Msg("Starting auth Service on :8081")
	http.ListenAndServe(":8081", r)

	return nil
}

func setupRouter(auc *usecase.AuthUsecase) *gin.Engine {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		log.Warn().Str("path", c.Request.URL.Path).Int("status", http.StatusNotFound).Str("status_text", http.StatusText(http.StatusNotFound)).Msg("page not found")
		c.String(http.StatusNotFound, "404 page not found")
	})

	return r
}
