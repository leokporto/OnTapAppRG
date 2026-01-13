package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/leokporto/OnTapAppRG/backend/internal/beer"
	"github.com/leokporto/OnTapAppRG/backend/internal/health"
)

func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(30),
	)

	r.Get("/health", health.Handler())

	r.Get("/beers", beer.GetAll)
	r.Get("/beers/{id}", beer.GetById)
	r.Get("/beers/styles", beer.GetStyles)

	http.ListenAndServe(":8080", r)
}
