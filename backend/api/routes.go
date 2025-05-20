package api

import (
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", HelloWorld)

	return r
}
