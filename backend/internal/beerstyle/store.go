package beerstyle

import (
	"context"
	"database/sql"
)

type Store interface {
	Create(ctx context.Context, beer *BeerStyle) error
	List(ctx context.Context) ([]BeerStyle, error)
}

type pgSqlStore struct {
	db *sql.DB
}

func NewPgSqlStore(db *sql.DB) Store {
	return &pgSqlStore{db: db}
}

func (pgSqlStore *pgSqlStore) Create(ctx context.Context, style *BeerStyle) error {
	query := `INSERT INTO styles (name) 
				VALUES ($1) 
				RETURNING id`

	return pgSqlStore.db.QueryRowContext(
		ctx, query, style.Name,
	).Scan(&style.ID)
}

func (pgSqlStore *pgSqlStore) List(ctx context.Context) ([]BeerStyle, error) {
	query := `SELECT id, name FROM styles`

	rows, err := pgSqlStore.db.QueryContext(ctx, query)
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
