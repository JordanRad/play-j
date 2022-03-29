package dbaccount

import (
	"context"
	"database/sql"
	"fmt"

	account "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
)

type Store struct {
	DB *sql.DB
}

var _ account.AccountStore = (*Store)(nil)

func (s *Store) CreateUser(ctx context.Context, user *account.User) (bool, error) {

	result, err := s.DB.Exec("INSERT INTO accounts (email,username,password, firstName,lastName) VALUES ($1,$2,$3,$4,$5);", user.Email, "username1", user.Password, user.FirstName, user.LastName)

	if err != nil {
		return false, fmt.Errorf("error creating a user: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 1 {
		return true, nil
	}

	return false, fmt.Errorf("error inserting the user")
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*dbmodels.Account, error) {
	var result dbmodels.Account

	row := s.DB.QueryRow(`SELECT * FROM accounts WHERE email = $1;`, email)
	err := row.Scan(&result.ID, &result.FirstName, &result.LastName,
		&result.Username, &result.Email, &result.Password)

	if err == sql.ErrNoRows {
		fmt.Println(err)
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}

	return &result, nil
}
