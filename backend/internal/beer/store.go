package beer

import (
	"context"
	"database/sql"
)

type BeerStore interface {
	Create(ctx context.Context, beer *Beer) error
	// List(ctx context.Context) ([]Beer, error)
	GetById(ctx context.Context, beerId int64) (Beer, error)
	// Find(ctx context.Context, beer *Beer, filter string) ([]Beer, error)
	// ListByBrewery(ctx context.Context, breweryId int64) ([]Beer, error)
}

type pgSqlStore struct {
	db *sql.DB
}

func NewPgSqlStore(db *sql.DB) BeerStore {
	return &pgSqlStore{db: db}
}

func (pgSqlStore *pgSqlStore) GetById(ctx context.Context, beerId int64) (Beer, error) {
	query := `SELECT id, name, style_id, brewery_id, fullname, abv, minibu, maxibu FROM beers
			  WHERE id = $1`

	row := pgSqlStore.db.QueryRowContext(ctx, query, beerId)

	var filteredBeer Beer

	err := row.Scan(&filteredBeer.ID, &filteredBeer.Name, &filteredBeer.StyleID, &filteredBeer.BreweryID,
		&filteredBeer.FullName, &filteredBeer.ABV, &filteredBeer.MinIBU, &filteredBeer.MaxIBU)

	if err != nil {
		return Beer{}, err
	}

	return filteredBeer, nil
}

func (pgSqlStore *pgSqlStore) Create(ctx context.Context, beer *Beer) error {
	query := `INSERT INTO beers (name, styleid, breweryid, fullname, abv, minibu, maxibu) 
				VALUES ($1, $2, $3, $4, $5, $6, $7) 
				RETURNING id`

	return pgSqlStore.db.QueryRowContext(
		ctx, query, beer.Name, beer.StyleID, beer.BreweryID, beer.FullName, beer.ABV, beer.MinIBU, beer.MaxIBU,
	).Scan(&beer.ID)
}

// func (pgSqlStore *pgSqlStore) List(ctx context.Context) ([]Beer, error) {
// 	query := `SELECT id, name, style_id, brewery_id, fullname, abv, minibu, maxibu FROM beers`

// 	rows, err := pgSqlStore.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var beers []Beer
// 	for rows.Next() {
// 		var beer Beer
// 		if err := rows.Scan(&beer.ID, &beer.Name, &beer.StyleID, &beer.BreweryID,
// 			&beer.FullName, &beer.ABV, &beer.MinIBU, &beer.MaxIBU); err != nil {
// 			return nil, err
// 		}
// 		beers = append(beers, beer)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return beers, nil
// }

// func (pgSqlStore *pgSqlStore) Find(ctx context.Context, beer *Beer, filter string) ([]Beer, error) {
// 	query := `SELECT id, name, style_id, brewery_id, fullname, abv, minibu, maxibu FROM beers
// 			  WHERE fullname ILIKE $1`

// 	rows, err := pgSqlStore.db.QueryContext(ctx, query, "%"+filter+"%")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var beers []Beer
// 	for rows.Next() {
// 		var beer Beer
// 		if err := rows.Scan(&beer.ID, &beer.Name, &beer.StyleID, &beer.BreweryID,
// 			&beer.FullName, &beer.ABV, &beer.MinIBU, &beer.MaxIBU); err != nil {
// 			return nil, err
// 		}
// 		beers = append(beers, beer)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return beers, nil

// }

// func (pgSqlStore *pgSqlStore) ListByBrewery(ctx context.Context, breweryId int64) ([]Beer, error) {
// 	query := `SELECT id, name, style_id, brewery_id, fullname, abv, minibu, maxibu FROM beers
// 			  WHERE brewery_id = $1`

// 	rows, err := pgSqlStore.db.QueryContext(ctx, query, breweryId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var beers []Beer
// 	for rows.Next() {
// 		var beer Beer
// 		if err := rows.Scan(&beer.ID, &beer.Name, &beer.StyleID, &beer.BreweryID,
// 			&beer.FullName, &beer.ABV, &beer.MinIBU, &beer.MaxIBU); err != nil {
// 			return nil, err
// 		}
// 		beers = append(beers, beer)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return beers, nil
// }
