package health

import "github.com/go-chi/chi/v5"

func MapRoutes(r chi.Router, h *Handler) {
	r.Get("/api/health", h.Handler)
}
