package main

import (
	"database/sql"
	"log"
	"net/http"

	"os"

	"github.com/leokporto/OnTapAppRG/backend/internal/auth"
	"github.com/leokporto/OnTapAppRG/backend/internal/beerread"
	"github.com/leokporto/OnTapAppRG/backend/internal/beerstyle"
	"github.com/leokporto/OnTapAppRG/backend/internal/brewery"
	"github.com/leokporto/OnTapAppRG/backend/internal/config"
	"github.com/leokporto/OnTapAppRG/backend/internal/health"
	"github.com/leokporto/OnTapAppRG/backend/internal/http/router"
	"github.com/leokporto/OnTapAppRG/backend/internal/user"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	r := router.New()

	//Config
	configVals, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//config google auth
	googleCfg := auth.NewGoogleConfig(
		os.Getenv("GOOGLE_CLIENT_ID"),
		os.Getenv("GOOGLE_CLIENT_SECRET"),
		os.Getenv("GOOGLE_REDIRECT_URL"),
	)

	//Db conn
	db, err := sql.Open("pgx", configVals.Conn_String)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Stores and Handlers
	userStore := user.NewPgSqlStore(db)
	authService := auth.NewService(userStore)
	jwtService := auth.NewJWTService(configVals.JwtSecret, configVals.JwtIssuer, configVals.JwtTTLMin)
	authHandler := auth.NewHandler(googleCfg, authService, jwtService)

	beerReadStore := beerread.NewPgSqlStore(db)
	beerReadHandler := beerread.NewHandler(beerReadStore)

	breweryStore := brewery.NewPgSqlStore(db)
	breweryHandler := brewery.NewHandler(breweryStore)

	beerStyleStore := beerstyle.NewPgSqlStore(db)
	beerStyleHandler := beerstyle.NewHandler(beerStyleStore)

	healthHandler := health.NewHandler()

	//Routes
	health.MapRoutes(r, healthHandler)
	beerread.MapRoutes(r, beerReadHandler)
	beerstyle.MapRoutes(r, beerStyleHandler)
	brewery.MapRoutes(r, breweryHandler)
	auth.MapRoutes(r, authHandler)

	http.ListenAndServe(":8080", r)
}
