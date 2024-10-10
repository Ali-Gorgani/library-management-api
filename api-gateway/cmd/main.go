package main

import (
	"library-management-api/api-gateway/routes"
	authConfigs "library-management-api/auth-service/configs"
	authDB "library-management-api/auth-service/init/database"
	bookConfigs "library-management-api/books-service/configs"
	bookDB "library-management-api/books-service/init/database"
	userConfigs "library-management-api/users-service/configs"
	userDB "library-management-api/users-service/init/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	authConfigs.RunConfig("auth-service")
	userConfigs.RunConfig("users-service")
	bookConfigs.RunConfig("books-service")
	authDB.RunDB()
	userDB.RunDB()
	bookDB.RunDB()
}

func main() {
	r := gin.Default()
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.BookRoutes(r)

	r.NoRoute(func(c *gin.Context) {
		log.Warn().Str("path", c.Request.URL.Path).Int("status", http.StatusNotFound).Str("status_text", http.StatusText(http.StatusNotFound)).Msg("page not found")
		c.String(http.StatusNotFound, "404 page not found")
	})

	r.Static("/swagger", "./util/swagger")

	log.Info().Msg("Starting api-gateway on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start api gateway service")
	}
}
