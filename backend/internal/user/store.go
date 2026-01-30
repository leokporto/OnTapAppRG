package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type UserStore interface {
	Create(ctx context.Context, user *User) error
	GetBySubject(ctx context.Context, provider string, subject string) (User, error)
}

type pgSqlStore struct {
	db *sql.DB
}

func NewPgSqlStore(db *sql.DB) UserStore {
	return &pgSqlStore{db: db}
}

func (pgSqlStore *pgSqlStore) GetBySubject(ctx context.Context, provider string, subject string) (User, error) {
	query := `SELECT id, openid_provider, openid_subject, email, display_name, created_at, last_login_at FROM users
			  WHERE openid_provider = $1 AND openid_subject = $2`

	row := pgSqlStore.db.QueryRowContext(ctx, query, provider, subject)
	var filteredUser User

	err := row.Scan(&filteredUser.ID, &filteredUser.OpenIDProvider, &filteredUser.OpenIDSubject, &filteredUser.Email,
		&filteredUser.DisplayName, &filteredUser.CreatedAt, &filteredUser.LastLoginAt)

	if err != nil {
		return User{}, err
	}

	return filteredUser, nil
}

func (pgSqlStore *pgSqlStore) Create(ctx context.Context, user *User) error {

	if user.OpenIDProvider == "" || user.OpenIDSubject == "" {
		return errors.New("OpenId parameters missing")
	}

	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	query := `INSERT INTO users (id, openid_provider, openid_subject, email, display_name, created_at, last_login_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7) 
				RETURNING id`

	return pgSqlStore.db.QueryRowContext(
		ctx, query, user.ID, user.OpenIDProvider, user.OpenIDSubject, user.Email, user.DisplayName, user.CreatedAt, user.LastLoginAt,
	).Scan(&user.ID)
}
