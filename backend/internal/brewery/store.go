package brewery

import (
	"context"
	"database/sql"
)

type Store interface {
	Create(ctx context.Context, brewery *Brewery) error
	List(ctx context.Context) ([]Brewery, error)
	GetById(ctx context.Context, breweryId int64) (Brewery, error)
	Find(ctx context.Context, brewery *Brewery, filter string) ([]Brewery, error)
}

type pgSqlStore struct {
	db *sql.DB
}

func NewPgSqlStore(db *sql.DB) Store {
	return &pgSqlStore{db: db}
}

func (pgSqlStore *pgSqlStore) Create(ctx context.Context, brewery *Brewery) error {
	query := `INSERT INTO breweries (name) 
				VALUES ($1) 
				RETURNING id`

	return pgSqlStore.db.QueryRowContext(
		ctx, query, brewery.Name,
	).Scan(&brewery.ID)
}

func (pgSqlStore *pgSqlStore) List(ctx context.Context) ([]Brewery, error) {
	query := `SELECT id, name FROM breweries`

	rows, err := pgSqlStore.db.QueryContext(ctx, query)
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

func (pgSqlStore *pgSqlStore) GetById(ctx context.Context, breweryId int64) (Brewery, error) {
	query := `SELECT id, name FROM breweries WHERE id = $1`

	row := pgSqlStore.db.QueryRowContext(ctx, query, breweryId)

	var brewery Brewery

	err := row.Scan(&brewery.ID, &brewery.Name)
	if err != nil {
		return Brewery{}, err
	}

	return brewery, nil
}

func (pgSqlStore *pgSqlStore) Find(ctx context.Context, brewery *Brewery, filter string) ([]Brewery, error) {
	query := `SELECT id, name FROM breweries
			  WHERE name ILIKE $1`

	rows, err := pgSqlStore.db.QueryContext(ctx, query, "%"+filter+"%")
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
