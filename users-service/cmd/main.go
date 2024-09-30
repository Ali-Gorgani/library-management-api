package main

import (
	"library-management-api/users-service/core/usecase"
	"library-management-api/users-service/init/database"
	"library-management-api/users-service/init/migrations"
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

	uuc := usecase.NewUserUseCase()
	r := setupRouter(uuc)

	log.Info().Msg("Starting Users Service on :8081")
	http.ListenAndServe(":8081", r)

	return nil
}

func setupRouter(uuc *usecase.UserUsecase) *gin.Engine {
	r := gin.Default()

	r.POST("/users", uuc.AddUser)
	r.GET("/users", uuc.GetUsers)
	r.GET("/users/:id", uuc.GetUser)
	r.PUT("/users/:id", uuc.UpdateUser)
	r.DELETE("/users/:id", uuc.DeleteUser)
	r.POST("/users/login", uuc.Login)

	r.NoRoute(func(c *gin.Context) {
		log.Warn().Str("path", c.Request.URL.Path).Int("status", http.StatusNotFound).Str("status_text", http.StatusText(http.StatusNotFound)).Msg("page not found")
		c.String(http.StatusNotFound, "404 page not found")
	})

	return r
}
