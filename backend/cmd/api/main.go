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

	//beerStore := beer.NewPgSqlStore(db)
	//beerHandler := beer.NewHandler(beerStore)

	beerReadStore := beerread.NewPgSqlStore(db)
	beerReadHandler := beerread.NewHandler(beerReadStore)

	breweryStore := brewery.NewPgSqlStore(db)
	breweryHandler := brewery.NewHandler(breweryStore)

	beerStyleStore := beerstyle.NewPgSqlStore(db)
	beerStyleHandler := beerstyle.NewHandler(beerStyleStore)

	r.Get("/api/health", health.Handler())

	r.Get("/api/beers", beerReadHandler.List)
	r.Get("/api/beers/{id}", beerReadHandler.GetById)
	r.Get("/api/beers/styles", beerStyleHandler.List)
	r.Get("/api/breweries", breweryHandler.List)
	r.Get("/api/breweries/{id}", breweryHandler.GetById)
	http.ListenAndServe(":8080", r)
}
