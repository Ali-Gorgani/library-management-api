package main

import (
	"library-management-api/books-service/core/usecase"
	"library-management-api/books-service/init/database"
	"library-management-api/books-service/init/migrations"
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

	buc := usecase.NewBookUseCase()

	r := chi.NewRouter()
	r.Post("/books", buc.AddBook)
	r.Get("/books", buc.GetBooks)
	r.Put("/books/{id}", buc.UpdateBook)
	r.Delete("/books/{id}", buc.DeleteBook)
	r.Post("/books/{id}/borrow", buc.BorrowBook)
	r.Post("/books/{id}/return", buc.ReturnBook)
	r.Get("/books/search", buc.SearchBooks)
	r.Get("/books/category/{category}", buc.CategoryBooks)
	r.Get("/books/available", buc.AvailableBooks)

	log.Println("Starting Books Service on :8082")
	http.ListenAndServe(":8082", r)

	return nil
}
