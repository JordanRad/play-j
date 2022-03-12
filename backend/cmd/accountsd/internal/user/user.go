package user

import (
	"context"
	"fmt"

	"git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/gen/user"
)

type Store interface {
	CreateUser(context.Context, *User) (*User, error)
}

type User struct {
	ID       string
	Email    string
	Password string
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

// Compile-time assertion that *Service implements the user.Service interface.
var _ user.Service = (*Service)(nil)

func (s *Service) Register(ctx context.Context, p *user.RegisterPayload) (*user.RegisterResponse, error) {
	// Check passwords
	// Encrypt password

	newUser := &User{
		Email:    p.Email,
		Password: p.Password,
	}

	fmt.Println(newUser)
	//Save in the DB
	_, err := s.store.CreateUser(ctx, newUser)

	if err != nil {
		return nil, fmt.Errorf("Error creating new user: %w", err)
	}

	response := &user.RegisterResponse{
		Message: "Successfully created.",
	}

	return response, nil
}
