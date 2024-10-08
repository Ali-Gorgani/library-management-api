package database

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "root",
		Password: "secret",
		Database: "library_auth_db",
		SSLMode:  "disable",
	}
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}

type postgres struct {
	DB *sql.DB
}

var p *postgres

func P() *postgres {
	return p
}

// Open opens a database connection using the provided PostgresConfig.
// It attempts to establish a connection and verify it with a ping.
// Caller must ensure that the connection is closed via db.Close() method.
func Open(cfg PostgresConfig) {
	var err error
	p = &postgres{}
	p.DB, err = sql.Open("pgx", cfg.String())
	if err != nil {
		log.Fatal().Err(err).Msg("Open: failed to open database connection")
	}
	err = p.DB.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Open: ping failed")
	}
	log.Info().Msg("Database connected!")
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	return nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	if dir == "" {
		dir = "."
	}
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}
