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

	r.Get("/health", health.Handler())

	r.Get("/beers", beerHandler.GetAll)
	r.Get("/beers/{id}", beerHandler.GetById)
	r.Get("/beers/styles", beerHandler.GetStyles)

	http.ListenAndServe(":8080", r)
}
