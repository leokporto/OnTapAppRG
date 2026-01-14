package beer

import (
	"context"
	"database/sql"
)

type BeerStore interface {
	Create(ctx context.Context, beer *Beer) error
	ListAll(ctx context.Context) ([]Beer, error)
}

type pgSqlBeerStore struct {
	db *sql.DB
}

func NewPgSqlBeerStore(db *sql.DB) BeerStore {
	return &pgSqlBeerStore{db: db}
}

func (pgSqlBeerStore *pgSqlBeerStore) Create(ctx context.Context, beer *Beer) error {
	query := `INSERT INTO beers (name, style, brewery, fullname, abv, minibu, maxibu) 
				VALUES ($1, $2, $3, $4, $5, $6, $7) 
				RETURNING id`

	return pgSqlBeerStore.db.QueryRowContext(
		ctx, query, beer.Name, beer.Style, beer.Brewery, beer.FullName, beer.ABV, beer.MinIBU, beer.MaxIBU,
	).Scan(&beer.ID)
}

func (pgSqlBeerStore *pgSqlBeerStore) ListAll(ctx context.Context) ([]Beer, error) {
	query := `SELECT id, name, style, brewery, fullname, abv, minibu, maxibu FROM beers`

	rows, err := pgSqlBeerStore.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var beers []Beer
	for rows.Next() {
		var beer Beer
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

// func OpenConnection() (*sql.DB, error) {
// 	db, err := sql.Open("postgres", "postgresql://postgres:P@55w0rd@localhost:5432/beers_db?sslmode=disable")
// 	if err != nil {
// 		return nil, err
// 	}

// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(5)
// 	db.SetConnMaxLifetime(time.Hour)

// 	return db, db.Ping()
// }
