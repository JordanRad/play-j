package user

import (
	"context"
	"errors"
	"fmt"

	"git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/db/dbmodels"
	auth "git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/auth"
	"git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/gen/account"
)

type AccountStore interface {
	CreateUser(context.Context, *User) (bool, error)
	GetUserByEmail(context.Context, string) (*dbmodels.Account, error)
}

type PlaylistStore interface {
	CreateAccountPlaylist(context.Context, uint, string) (bool, error)
	GetUserPlaylist(context.Context, uint, uint) (string, error)
	GetAllUserPlaylists(context.Context, uint) ([]string, error)
}

type Service struct {
	accountStore  AccountStore
	playlistStore PlaylistStore
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

func NewService(accountStore AccountStore, playlistStore PlaylistStore) *Service {
	return &Service{
		accountStore:  accountStore,
		playlistStore: playlistStore,
	}
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
	_, err := s.accountStore.CreateUser(ctx, newUser)

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
	foundAccount, err := s.accountStore.GetUserByEmail(ctx, *p.Email)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Verify the password hash
	isPasswordCorrect := auth.CheckPassword(foundAccount.Password, *p.Password)

	if !isPasswordCorrect {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := auth.GenerateJWT(foundAccount.ID, foundAccount.Email)

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
