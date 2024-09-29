package main

import (
	"library-management-api/users-service/core/usecase"
	"library-management-api/users-service/init/database"
	"library-management-api/users-service/init/migrations"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func init() {
	database.Open(database.DefaultPostgresConfig())
}

func run() error {
	db := database.P().DB
	err := database.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		return err
	}

	uuc := usecase.NewUserUseCase()

	r := chi.NewRouter()

	r.Post("/users", uuc.AddUser)
	r.Get("/users", uuc.GetUsers)
	r.Get("/users/{id}", uuc.GetUser)
	r.Put("/users/{id}", uuc.UpdateUser)
	r.Delete("/users/{id}", uuc.DeleteUser)
	r.Post("/users/login", uuc.Login)

	// start the server
	log.Println("Starting Users Service on :8081")
	http.ListenAndServe(":8081", r)

	return nil
}
