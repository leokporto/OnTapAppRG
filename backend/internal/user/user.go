package user

import (
	"time"
)

type User struct {
	ID             string    `db:"id"`
	OpenIDProvider string    `db:"openid_provider"`
	OpenIDSubject  string    `db:"openid_subject"`
	Email          string    `db:"email"`
	DisplayName    string    `db:"display_name"`
	CreatedAt      time.Time `db:"created_at"`
	LastLoginAt    time.Time `db:"last_login_at"`
}
