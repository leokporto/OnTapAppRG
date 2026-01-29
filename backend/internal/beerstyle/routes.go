package beerstyle

import (
	"github.com/go-chi/chi/v5"
)

func MapRoutes(r chi.Router, h *Handler) {
	r.Route("/api/beers/styles", func(r chi.Router) {
		r.Get("/", h.List)
	})
}
