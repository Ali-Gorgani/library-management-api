package main

import (
	"log"
	"net/http"
	"users-service/migrations"
	"users-service/pkg/database"
	"users-service/pkg/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	// Setup the database
	db, err := database.Open(database.DefaultPostgresConfig())
	if err != nil {
		return err
	}
	defer db.Close()

	err = database.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Get("/users/{id}", handlers.GetUser)       // Get user by ID
	r.Post("/users/signup", handlers.Signup)     // Create user
	r.Post("/users/login", handlers.Login)       // Login user
	r.Put("/users/{id}", handlers.UpdateUser)    // Update user by ID
	r.Delete("/users/{id}", handlers.DeleteUser) // Delete user by ID

	log.Println("Starting Users Service on :8081")
	http.ListenAndServe(":8081", r)

	return nil
}
