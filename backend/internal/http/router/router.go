package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	appmw "github.com/leokporto/OnTapAppRG/backend/internal/http/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(15*time.Second),
		appmw.CORS(),
	)

	return r
}
