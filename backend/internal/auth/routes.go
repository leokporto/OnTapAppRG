package auth

import "github.com/go-chi/chi/v5"

func MapRoutes(r chi.Router, h *Handler) {
	r.Route("/auth", func(r chi.Router) {
		r.Get("/google/login", h.GoogleLogin)
		r.Get("/google/callback", h.GoogleCallback)
	})
}
