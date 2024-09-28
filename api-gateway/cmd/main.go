package main

import (
    "log"
    "net/http"
    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()

    // Define routes for user and book services
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome to the Library Management API!"))
    })
    
    r.Mount("/users", usersRouter())
    r.Mount("/books", booksRouter())

    log.Println("Starting API Gateway on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}

func usersRouter() http.Handler {
    r := chi.NewRouter()
    // Define user service routes here
    return r
}

func booksRouter() http.Handler {
    r := chi.NewRouter()
    // Define book service routes here
    return r
}
