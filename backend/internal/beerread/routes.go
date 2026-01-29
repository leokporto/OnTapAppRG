package beerread

import (
	"github.com/go-chi/chi/v5"
)

func MapRoutes(r chi.Router, h *Handler) {
	r.Route("/api/beers", func(r chi.Router) {
		r.Get("/", h.Find)
		r.Get("/{id}", h.GetById)
	})
}
