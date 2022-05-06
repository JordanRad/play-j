package dbaccount

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	account "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
)

type Store struct {
	DB *sql.DB
}

var _ account.Store = (*Store)(nil)

func (s *Store) CreateUser(ctx context.Context, user *account.User) (bool, error) {
	log.Printf("Creating new user with email: %s.\n", user.Email)
	result, err := s.DB.Exec("INSERT INTO accounts (email,username,password, firstName,lastName) VALUES ($1,$2,$3,$4,$5);", user.Email, user.Username, user.Password, user.FirstName, user.LastName)

	if err != nil {
		return false, fmt.Errorf("error creating a user: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 1 {
		log.Printf("New user %s has been created successfully.\n", user.Email)
		return true, nil
	}

	return false, fmt.Errorf("error inserting the user")
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*dbmodels.Account, error) {
	log.Printf("Get user with email: %s\n", email)
	var result dbmodels.Account

	row := s.DB.QueryRow(`SELECT * FROM accounts WHERE email = $1;`, email)
	err := row.Scan(&result.ID, &result.FirstName, &result.LastName,
		&result.Username, &result.Email, &result.Password, &result.IsActive, &result.HasActiveSubscription)

	if err == sql.ErrNoRows {
		log.Printf("User with email: %s has not been found.\n", email)
		return nil, fmt.Errorf("cannot find account by email: %w", err)
	}

	return &result, nil
}
