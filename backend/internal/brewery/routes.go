package brewery

import (
	"github.com/go-chi/chi/v5"
)

func MapRoutes(r chi.Router, h *Handler) {
	r.Route("/api/breweries", func(r chi.Router) {
		r.Get("/", h.List)
		r.Get("/{id}", h.GetById)
	})
}
