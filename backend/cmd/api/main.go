package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/leokporto/OnTapAppRG/backend/internal/beerread"
	"github.com/leokporto/OnTapAppRG/backend/internal/beerstyle"
	"github.com/leokporto/OnTapAppRG/backend/internal/brewery"
	"github.com/leokporto/OnTapAppRG/backend/internal/config"
	"github.com/leokporto/OnTapAppRG/backend/internal/health"

	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(15*time.Second),
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://localhost:5173"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	configVals, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", configVals.Conn_String)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//beerStore := beer.NewPgSqlStore(db)
	//beerHandler := beer.NewHandler(beerStore)

	beerReadStore := beerread.NewPgSqlStore(db)
	beerReadHandler := beerread.NewHandler(beerReadStore)

	breweryStore := brewery.NewPgSqlStore(db)
	breweryHandler := brewery.NewHandler(breweryStore)

	beerStyleStore := beerstyle.NewPgSqlStore(db)
	beerStyleHandler := beerstyle.NewHandler(beerStyleStore)

	r.Get("/api/health", health.Handler())

	r.Route("/api/beers", func(r chi.Router) {
		r.Get("/", beerReadHandler.Find)
		r.Get("/{id}", beerReadHandler.GetById)
		r.Get("/styles", beerStyleHandler.List)
	})

	r.Route("/api/breweries", func(r chi.Router) {
		r.Get("/", breweryHandler.List)
		r.Get("/{id}", breweryHandler.GetById)
	})

	http.ListenAndServe(":8080", r)
}
