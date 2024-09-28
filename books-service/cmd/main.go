package main

import (
	"books-service/migrations"
	"books-service/pkg/database"
	"books-service/pkg/handlers"
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
	r.Post("/books", handlers.AddBook) // Create book
	r.Get("/books", handlers.GetBooks) // Get book by ID

	log.Println("Starting Books Service on :8082")
	http.ListenAndServe(":8082", r)

	return nil
}
