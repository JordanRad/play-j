package dbuser

import (
	"context"

	"git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/user"
)

type Store struct {
}

func (s *Store) CreateUser(ctx context.Context, user *user.User) (createdUser *user.User, err error) {
	return nil, nil
}

var _ user.Store = (*Store)(nil)
