package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/leokporto/OnTapAppRG/backend/internal/beer"
	"github.com/leokporto/OnTapAppRG/backend/internal/config"
	"github.com/leokporto/OnTapAppRG/backend/internal/health"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(15*time.Second),
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

	beerStore := beer.NewPgSqlBeerStore(db)
	beerHandler := beer.NewHandler(beerStore)

	r.Get("/api/health", health.Handler())

	r.Get("/api/beers", beerHandler.GetAll)
	r.Get("/api/beers/{id}", beerHandler.GetById)
	r.Get("/api/beers/styles", beerHandler.GetStyles)
	r.Get("/api/breweries", beerHandler.GetBreweries)

	http.ListenAndServe(":8080", r)
}
