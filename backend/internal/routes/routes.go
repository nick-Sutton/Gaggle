package routes

import (
    "github.com/go-chi/chi/v5"
    "myapp/internal/handlers"
)

func SetupRouter() *chi.Mux {
    r := chi.NewRouter()

    r.Get("/", handlers.HelloWorld)

    return r
}
