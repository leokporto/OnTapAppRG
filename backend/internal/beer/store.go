package beer

import (
	"context"
	"database/sql"
)

type BeerStore interface {
	Create(ctx context.Context, beer *Beer) error
	ListAll(ctx context.Context) ([]BeerResponse, error)
}

type pgSqlBeerStore struct {
	db *sql.DB
}

func NewPgSqlBeerStore(db *sql.DB) BeerStore {
	return &pgSqlBeerStore{db: db}
}

func (pgSqlBeerStore *pgSqlBeerStore) Create(ctx context.Context, beer *Beer) error {
	query := `INSERT INTO beers (name, styleid, breweryid, fullname, abv, minibu, maxibu) 
				VALUES ($1, $2, $3, $4, $5, $6, $7) 
				RETURNING id`

	return pgSqlBeerStore.db.QueryRowContext(
		ctx, query, beer.Name, beer.StyleID, beer.BreweryID, beer.FullName, beer.ABV, beer.MinIBU, beer.MaxIBU,
	).Scan(&beer.ID)
}

func (pgSqlBeerStore *pgSqlBeerStore) ListAll(ctx context.Context) ([]BeerResponse, error) {
	query := `SELECT beers.id, beers.name, styles.name, breweries.name, fullname, abv, minibu, maxibu FROM beers, styles, breweries
			  WHERE beers.style_id = styles.id AND beers.brewery_id = breweries.id`

	rows, err := pgSqlBeerStore.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beers []BeerResponse
	for rows.Next() {
		var beer BeerResponse
		if err := rows.Scan(&beer.ID, &beer.Name, &beer.Style, &beer.Brewery,
			&beer.FullName, &beer.ABV, &beer.MinIBU, &beer.MaxIBU); err != nil {
			return nil, err
		}
		beers = append(beers, beer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return beers, nil
}
