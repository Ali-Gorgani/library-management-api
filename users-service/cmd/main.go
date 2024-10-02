package main

import (
	"library-management-api/users-service/adapter/middleware"
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

	secretKey := "mrlIpbCvRvrNubGCvf2CPy3OMZCXwXDHRz4SyPfFVcU="
	uuc := usecase.NewUserUseCase(secretKey)
	r := setupRouter(uuc)

	log.Info().Msg("Starting Users Service on :8081")
	http.ListenAndServe(":8081", r)

	return nil
}

func setupRouter(uuc *usecase.UserUsecase) *gin.Engine {
	r := gin.Default()
	tokenMaker := uuc.TokenMaker

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/", uuc.AddUser)
		usersGroup.POST("/login", uuc.Login)

		adminAuth := usersGroup.Group("/", middleware.AdminAuthMiddleware(tokenMaker))
		{
			adminAuth.GET("/", uuc.GetUsers)
			adminAuth.GET("/:id", uuc.GetUser)
			adminAuth.DELETE("/:id", uuc.DeleteUser)
		}

		userAuth := usersGroup.Group("/", middleware.UserAuthMiddleware(tokenMaker))
		{
			userAuth.Use(middleware.UserAuthMiddleware(tokenMaker))
			userAuth.PUT("/:id", uuc.UpdateUser)
			userAuth.POST("/logout", uuc.Logout)
		}

	}

	tokensGroup := r.Group("/tokens", middleware.UserAuthMiddleware(tokenMaker))
	{
		tokensGroup.POST("/renew", uuc.RenewAccessToken)
		tokensGroup.POST("/revoke", uuc.RevokeSession)
	}

	r.NoRoute(func(c *gin.Context) {
		log.Warn().Str("path", c.Request.URL.Path).Int("status", http.StatusNotFound).Str("status_text", http.StatusText(http.StatusNotFound)).Msg("page not found")
		c.String(http.StatusNotFound, "404 page not found")
	})

	return r
}

// RolePermissions defines which actions each role can perform
var RolePermissions = map[string][]string{
	"User":  {"GetBooks", "AddBook", "BorrowBook", "ReturnBook"},
	"Admin": {"UpdateBook", "DeleteBook", "GetBooks", "AddBook", "BorrowBook", "ReturnBook"},
}
