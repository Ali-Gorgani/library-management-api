package main

import (
	"library-management-api/books-service/init/database"
	"library-management-api/books-service/init/migrations"
	"os"

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

	return nil
}
