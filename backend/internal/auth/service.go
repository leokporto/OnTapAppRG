package auth

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/leokporto/OnTapAppRG/backend/internal/user"
)

type AuthService struct {
	users user.UserStore
	clock func() time.Time
}

func NewService(users user.UserStore) *AuthService {
	return &AuthService{
		users: users,
		clock: time.Now,
	}
}

func (s *AuthService) LoginWithOpenID(
	ctx context.Context,
	provider string,
	subject string,
	email string,
	name string,
) (*user.User, error) {

	u, err := s.users.GetBySubject(ctx, provider, subject)

	if err == sql.ErrNoRows {
		now := s.clock()

		newUser := &user.User{
			ID:             uuid.NewString(),
			OpenIDProvider: provider,
			OpenIDSubject:  subject,
			Email:          email,
			DisplayName:    name,
			CreatedAt:      now,
			LastLoginAt:    now,
		}

		if err := s.users.Create(ctx, newUser); err != nil {
			return nil, err
		}

		return newUser, nil
	}

	if err != nil {
		return nil, err
	}

	// opcional: atualizar last_login_at
	return &u, nil
}
