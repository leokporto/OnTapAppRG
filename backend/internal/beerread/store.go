package beerread

import (
	"context"
	"database/sql"
)

type BeerReadStore interface {
	List(ctx context.Context) ([]BeerDTO, error)
	GetById(ctx context.Context, beerId int64) (BeerDTO, error)
	Find(ctx context.Context, filter string) ([]BeerDTO, error)
	ListByBrewery(ctx context.Context, breweryId int64) ([]BeerDTO, error)
}

type pgSqlStore struct {
	db *sql.DB
}

func NewPgSqlStore(db *sql.DB) BeerReadStore {
	return &pgSqlStore{db: db}
}

func (pgSqlStore *pgSqlStore) List(ctx context.Context) ([]BeerDTO, error) {
	query := `SELECT beers.id, beers.name, styles.name, breweries.name, fullname, abv, minibu, maxibu FROM beers, styles, breweries
			  WHERE beers.style_id = styles.id AND beers.brewery_id = breweries.id`

	rows, err := pgSqlStore.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beers []BeerDTO
	for rows.Next() {
		var beer BeerDTO
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

func (pgSqlStore *pgSqlStore) Find(ctx context.Context, filter string) ([]BeerDTO, error) {
	query := `SELECT beers.id, beers.name, styles.name, breweries.name, fullname, abv, minibu, maxibu FROM beers, styles, breweries
			  WHERE fullname ILIKE $1`

	rows, err := pgSqlStore.db.QueryContext(ctx, query, "%"+filter+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beers []BeerDTO
	for rows.Next() {
		var beer BeerDTO
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

func (pgSqlStore *pgSqlStore) ListByBrewery(ctx context.Context, breweryId int64) ([]BeerDTO, error) {
	query := `SELECT beers.id, beers.name, styles.name, breweries.name, fullname, abv, minibu, maxibu FROM beers, styles, breweries
			  WHERE beers.style_id = styles.id AND beers.brewery_id = breweries.id AND beers.brewery_id = $1`

	rows, err := pgSqlStore.db.QueryContext(ctx, query, breweryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beers []BeerDTO
	for rows.Next() {
		var beer BeerDTO
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

func (pgSqlStore *pgSqlStore) GetById(ctx context.Context, beerId int64) (BeerDTO, error) {
	query := `SELECT beers.id, beers.name, styles.name, breweries.name, fullname, abv, minibu, maxibu FROM beers, styles, breweries
			  WHERE beers.style_id = styles.id AND beers.brewery_id = breweries.id AND beers.id = $1`

	row := pgSqlStore.db.QueryRowContext(ctx, query, beerId)

	var filteredBeer BeerDTO

	err := row.Scan(&filteredBeer.ID, &filteredBeer.Name, &filteredBeer.Style, &filteredBeer.Brewery,
		&filteredBeer.FullName, &filteredBeer.ABV, &filteredBeer.MinIBU, &filteredBeer.MaxIBU)

	if err != nil {
		return BeerDTO{}, err
	}

	return filteredBeer, nil
}
