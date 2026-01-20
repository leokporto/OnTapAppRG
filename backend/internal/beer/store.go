package beer

import (
	"context"
	"database/sql"
)

type BeerStore interface {
	Create(ctx context.Context, beer *Beer) error
	ListAll(ctx context.Context) ([]BeerResponse, error)
	GetByID(ctx context.Context, beerId int64) (BeerResponse, error)
	GetStyles(ctx context.Context) ([]BeerStyle, error)
	GetBreweries(ctx context.Context) ([]Brewery, error)
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

func (pgSqlBeerStore *pgSqlBeerStore) GetByID(ctx context.Context, beerId int64) (BeerResponse, error) {
	query := `SELECT beers.id, beers.name, styles.name, breweries.name, fullname, abv, minibu, maxibu FROM beers, styles, breweries
			  WHERE (beers.style_id = styles.id AND beers.brewery_id = breweries.id) AND beers.id = $1`

	row := pgSqlBeerStore.db.QueryRowContext(ctx, query, beerId)

	var filteredBeer BeerResponse

	err := row.Scan(&filteredBeer.ID, &filteredBeer.Name, &filteredBeer.Style, &filteredBeer.Brewery,
		&filteredBeer.FullName, &filteredBeer.ABV, &filteredBeer.MinIBU, &filteredBeer.MaxIBU)

	if err != nil {
		return BeerResponse{}, err
	}

	return filteredBeer, nil
}

func (pgSqlBeerStore *pgSqlBeerStore) GetStyles(ctx context.Context) ([]BeerStyle, error) {
	query := `SELECT id, name FROM styles`

	rows, err := pgSqlBeerStore.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var styles []BeerStyle
	for rows.Next() {
		var style BeerStyle
		if err := rows.Scan(&style.ID, &style.Name); err != nil {
			return nil, err
		}
		styles = append(styles, style)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return styles, nil
}

func (pgSqlBeerStore *pgSqlBeerStore) GetBreweries(ctx context.Context) ([]Brewery, error) {
	query := `SELECT id, name FROM breweries`

	rows, err := pgSqlBeerStore.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var breweries []Brewery
	for rows.Next() {
		var brewery Brewery
		if err := rows.Scan(&brewery.ID, &brewery.Name); err != nil {
			return nil, err
		}
		breweries = append(breweries, brewery)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return breweries, nil
}
