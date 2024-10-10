package database

import (
	"database/sql"
	"fmt"
	"io/fs"
	"library-management-api/users-service/configs"
	"library-management-api/users-service/init/migrations"

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
	// Load the configuration from config file or environment variables
	dbConfig := configs.C().PSQL

	return PostgresConfig{
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		User:     dbConfig.User,
		Password: dbConfig.Password,
		Database: dbConfig.Database,
		SSLMode:  dbConfig.SSLMode,
	}
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)
}

type Postgres struct {
	DB *sql.DB
}

var p *Postgres

func P() *Postgres {
	return p
}

// Open opens a database connection using the provided PostgresConfig.
// It attempts to establish a connection and verify it with a ping.
// Caller must ensure that the connection is closed via db.Close() method.
func Open(cfg PostgresConfig) {
	var err error
	p = &Postgres{}
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

func Migrate(db *sql.DB, dir string) {
	err := goose.SetDialect("postgres")
	if err != nil {
		log.Fatal().Err(err).Msg("migrate: failed to set dialect")
	}

	err = goose.Up(db, dir)
	if err != nil {
		log.Fatal().Err(err).Msg("migrate: failed to migrate")
	}
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) {
	if dir == "" {
		dir = "."
	}
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	Migrate(db, dir)
}

func RunDB() {
	Open(DefaultPostgresConfig())
	db := P().DB
	MigrateFS(db, migrations.FS, ".")
}
