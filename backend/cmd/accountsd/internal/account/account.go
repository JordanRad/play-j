package user

import (
	"context"
	"errors"
	"fmt"

	"git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/db/dbmodels"
	auth "git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/auth"
	"git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/gen/account"
)

type Store interface {
	CreateUser(context.Context, *User) (bool, error)
	GetUserByEmail(context.Context, string) (*dbmodels.Account, error)
}

type Service struct {
	store Store
}

type User struct {
	ID              string
	Email           string
	Password        string
	ConfirmPassword string
}
type UserEntry struct {
	ID        string
	Email     string
	Password  string
	Username  string
	FirstName string
	LastName  string
}
type LoginCredentials struct {
	Email    string
	Password string
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

// Compile-time assertion that *Service implements the user.Service interface.
var _ account.Service = (*Service)(nil)

func (s *Service) Register(ctx context.Context, p *account.RegisterPayload) (*account.RegisterResponse, error) {
	// Check passwords
	if p.Password != p.ConfirmedPassword {
		return nil, errors.New("passwords do not match")
	}

	// Encrypt password
	encryptedPassword, encryptErr := auth.EncryptPassword(p.Password)
	if encryptErr != nil {
		fmt.Errorf("error encrypting the password: %w", encryptErr)
	}

	newUser := &User{
		Email:    p.Email,
		Password: encryptedPassword,
	}

	// Save in the database
	_, err := s.store.CreateUser(ctx, newUser)

	if err != nil {
		return nil, fmt.Errorf("error creating new user: %w", err)
	}

	response := &account.RegisterResponse{
		Message: "Successfully created.",
	}

	return response, nil
}

func (s *Service) Login(ctx context.Context, p *account.LoginPayload) (*account.LoginResponse, error) {
	// Get the user by email
	foundAccount, err := s.store.GetUserByEmail(ctx, *p.Email)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	fmt.Printf("Token is: ")

	// Verify the password hash
	isPasswordCorrect := auth.CheckPassword(foundAccount.Password, *p.Password)

	if !isPasswordCorrect {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := auth.GenerateJWT(foundAccount.ID)
	if err != nil {
		return nil, fmt.Errorf("error extracting jwt: %w", err)
	}
	r := &account.LoginResponse{
		Email:        *p.Email,
		Token:        token,
		RefreshToken: "dfsdvsdfjv sjf",
		Role:         "ROLE_REG_USER",
	}
	return r, nil

}
